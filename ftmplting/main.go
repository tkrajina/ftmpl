package ftmplting

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"sort"
	"strings"
)

const VERSION = "v0.2.2"

const (
	cmdPrefix    = "!#"
	cmdNocompile = "!#nocompile"
	cmdArg       = "!#arg"
	cmdEscape    = "!#escape"
	cmdImport    = "!#import"
	cmdErrorif   = "!#errif"
	cmdReturn    = "!#return"
	cmdSub       = "!#sub"
	cmdExtends   = "!#extends"
	cmdInclude   = "!#include"
	cmdInsert    = "!#insert"
	cmdGlobal    = "!#global"
)

type Params struct {
	SourceDir     string
	TargetDir     string
	TargetGoFile  string
	FuncPrefix    string
	FuncPrefixErr string
}

type compiledTemplate struct {
	imports      []string
	originalFile string
	functionCode string
}

var exclamationMarkFixRegex = regexp.MustCompile("{{.*?}}\\!")
var templateReplacementRegex = regexp.MustCompile("{{.*?}}")
var fmtFormatRegex = regexp.MustCompile("\\%\\s{0,1}[\\d\\.]*\\w\\s+")

func Do(ap Params) {
	var compiledTemplates []compiledTemplate

	fileInfos, err := ioutil.ReadDir(ap.SourceDir)
	HandleError(err, "Error listing directory")

	var files []string
	for _, fileInfo := range fileInfos {
		files = append(files, fileInfo.Name())
	}
	sort.Strings(files)

	for _, fileName := range files {
		if strings.HasSuffix(fileName, ".tmpl") {

			params := convertTemplateParams{
				NoerrFuncPrefix: ap.FuncPrefix,
				ErrFuncPrefix:   ap.FuncPrefixErr,
			}
			compiled := convertTemplate(ap.SourceDir, fileName, params)
			if len(ap.TargetDir) > 0 {
				_, fileName := getLastPathElements(fileName)
				fileName = strings.Replace(fileName, ".tmpl", ".go", -1)
				fileName = path.Join(ap.TargetDir, fileName)
				if len(compiled.functionCode) > 0 {
					saveTemplates(fileName, compiled)
				}
			} else {
				compiledTemplates = append(compiledTemplates, compiled)
			}
		}
	}

	if len(ap.TargetGoFile) > 0 {
		saveTemplates(ap.TargetGoFile, compiledTemplates...)
	}
}

func printfWarning(str string, args ...interface{}) {
	if !strings.HasSuffix(str, "\n") {
		str += "\n"
	}
	fmt.Fprintln(os.Stderr, "----------------------------------------------------------------------------------------------------")
	fmt.Fprintf(os.Stderr, str, args...)
	fmt.Fprintln(os.Stderr, "----------------------------------------------------------------------------------------------------")
}

func saveTemplates(destination string, compiled ...compiledTemplate) {
	fmt.Sprintln("destination=", destination)

	if f, err := os.Open(destination); err == nil {
		var bytes = make([]byte, 200)
		_, _ = f.Read(bytes)
		previousVersion := regexp.MustCompile("{{{.*?}}}").Find(bytes)
		if len(previousVersion) == 0 {
			printfWarning("Previous version not found in %s!\n", destination)
		} else {
			prev := strings.Replace(string(previousVersion), "{{{", "", 1)
			prev = strings.Replace(prev, "}}}", "", 1)
			if prev != VERSION {
				printfWarning("%s was previously compiled with %s, new version is %s!\n", destination, prev, VERSION)
			}
		}
	}

	var imports []string
	for _, c := range compiled {
		imports = append(imports, c.imports...)
	}
	imports = append(imports, "\"fmt\"", "\"os\"")
	imports = rmDoubleImports(imports)

	pathElem1, pathElem2 := getLastPathElements(destination)
	packageName := pathElem2
	if strings.Contains(pathElem2, ".go") {
		packageName = pathElem1
	}

	fOut, err := os.Create(destination)
	HandleError(err, "Error creating file")
	defer fOut.Close()

	_, _ = fOut.WriteString("// package " + packageName + " is generated with ftmpl {{{" + VERSION + "}}}, do not edit!!!! */\n")
	_, _ = fOut.WriteString("package " + packageName + "\n\n")
	_, _ = fOut.WriteString("import (\n")
	for _, i := range imports {
		fOut.WriteString(i + "\n")
	}
	_, _ = fOut.WriteString(")\n\n")
	_, _ = fOut.WriteString(initFunctionTemplate)
	_, _ = fOut.WriteString("\n\n")

	for _, c := range compiled {
		_, _ = fOut.WriteString("\n\n")
		_, _ = fOut.WriteString(c.functionCode)

		fmt.Printf("%-15s -> %-15s\n", c.originalFile, destination)
	}

	cmd := exec.Command("go", "fmt", destination)
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, strings.Join(cmd.Args, " "))
		fmt.Fprintf(os.Stderr, "Cannot format %s: %s\n\n", destination, err.Error())
		parseGoFmtErrorOutput(bytes)
	}
}

func processExtending(pckg string, chunks []string) []string {
	extendingTemplate := ""
	for _, chunk := range chunks {
		if strings.HasPrefix(chunk, cmdExtends) {
			extendingTemplate = strings.TrimSpace(chunk[len(cmdExtends):])
		}
	}

	if len(extendingTemplate) == 0 {
		return chunks
	}

	// Load base template:
	mainTemplateLines := loadTemplateAndGetChunks(path.Join(pckg, extendingTemplate)+".tmpl", false)

	// Parse sub template chunks:
	subTemplateCode, subTemplateChunks := loadTemplateSubChunks(chunks)

	var usedSubtemplateChunks = map[string]bool{}
	for k, _ := range subTemplateChunks {
		usedSubtemplateChunks[k] = false
	}

	// Merge them together:
	var result []string

	var subTemplateArgsCode []string
	for _, chunk := range subTemplateCode {
		if strings.HasPrefix(chunk, cmdArg) {
			subTemplateArgsCode = append(subTemplateArgsCode, chunk)
		} else {
			result = append(result, chunk)
		}
	}

	for _, chunk := range mainTemplateLines {
		if strings.HasPrefix(chunk, cmdInclude) {
			subName := strings.TrimSpace(chunk[len(cmdInclude):])
			if subLines, found := subTemplateChunks[subName]; found {
				usedSubtemplateChunks[subName] = true
				for _, subChunk := range subLines {
					result = append(result, subChunk)
				}
			}
		} else {
			result = append(result, chunk)
		}
	}

	// This is because we want the subtemplate args to be after the base template args in the function definition:
	for _, chunk := range subTemplateArgsCode {
		result = append(result, chunk)
	}

	for name, used := range usedSubtemplateChunks {
		if !used {
			printfWarning("Sub %s is declared but never used", name)
		}
	}

	return result
}

func loadTemplateSubChunks(lines []string) (general []string, subsections map[string][]string) {
	subReached := false
	currentSubName := ""

	general = make([]string, 0)
	subsections = make(map[string][]string)

	for _, chunk := range lines {
		if strings.Contains(chunk, cmdSub) {
			subReached = true
			currentSubName = strings.TrimSpace(chunk[len(cmdSub):])
		} else if subReached {
			if _, found := subsections[currentSubName]; !found {
				subsections[currentSubName] = make([]string, 0)
			}
			subsections[currentSubName] = append(subsections[currentSubName], chunk)
		} else {
			general = append(general, chunk)
		}
	}

	return general, subsections
}

func getChunks(str string) []string {
	delimiter := "<<<" + getRandomString(15) + ">>>"
	res := ""
	chunks := strings.Split(str, "\n")
	for n, chunk := range chunks {
		if n < len(chunks)-1 {
			chunk += "\n"
		}

		//fmt.Printf("line=%s\n", line)
		if len(chunk) == 0 {
		} else if strings.HasPrefix(chunk, "!!") {
			res += chunk
		} else if chunk[0] == '!' {
			res += delimiter + chunk + delimiter
		} else {
			chunk = exclamationMarkFixRegex.ReplaceAllStringFunc(chunk, func(s string) string {
				return s + "!"
			})
			res += templateReplacementRegex.ReplaceAllStringFunc(chunk, func(s string) string {
				s = strings.TrimSpace(s[2 : len(s)-2])
				if len(s) == 0 {
					return ""
				}

				if s[0] == '!' {
					return delimiter + s + delimiter
				}

				forceUnquoted := false
				if s[0] == '=' {
					forceUnquoted = true
					s = s[1:]
				}

				placeholder := "%v"
				var valueExpr string
				if len(s) > 0 && s[0] == '%' {
					placeholders := fmtFormatRegex.FindAllString(s, -1)
					if len(placeholders) == 0 {
						panic("Cannot find valid fmt format in:" + s)
					}
					placeholder = placeholders[0]
					valueExpr = s[len(placeholder):]
				} else if len(s) > 1 && s[1] == ' ' {
					valueExpr = s[1:]
					placeholder = "%" + string(s[0])
				} else {
					valueExpr = s
				}

				placeholder = strings.TrimSpace(placeholder)
				valueExpr = strings.TrimSpace(valueExpr)

				if !forceUnquoted && strings.HasSuffix(placeholder, "s") {
					valueExpr = "_escape(" + valueExpr + ")"
				}

				if strings.HasPrefix(s, "/*") && strings.HasSuffix(s, "*/") {
					return ""
				}

				res := fmt.Sprintf("!_, _ = _ftmpl.WriteString(fmt.Sprintf(`%s`, %s))", placeholder, valueExpr)
				return delimiter + res + delimiter
			})
		}
	}
	return strings.Split(res, delimiter)
}

// Loads the file and process it to get chunks.
//
// Note that "line" is not always a chunk from the template it is part of the template which will be converted to a line of code.
func loadTemplateAndGetChunks(fileName string, keepNoCompile bool) []string {
	str, err := loadFile(fileName)
	HandleError(err, "Error reading "+fileName)

	chunks := getChunks(str)

	// Process inserted templates:
	var expandedChunks []string
	for _, chunk := range chunks {
		if !keepNoCompile && strings.HasPrefix(chunk, cmdNocompile) {
			// Nothing
		} else if strings.HasPrefix(chunk, cmdInsert) {
			dir, _ := path.Split(fileName)
			fileToInsert := chunk[len(cmdInsert):]
			fileToInsert = strings.Trim(strings.TrimSpace(fileToInsert), "\"")
			fileToInsert = path.Join(dir, fileToInsert)
			expandedChunks = append(expandedChunks, loadTemplateAndGetChunks(fileToInsert, false)...)
		} else {
			expandedChunks = append(expandedChunks, chunk)
		}
	}

	return expandedChunks
}

type convertTemplateParams struct {
	ErrFuncPrefix, NoerrFuncPrefix string
}

func convertTemplate(packageDir, file string, params convertTemplateParams) compiledTemplate {
	chunks := loadTemplateAndGetChunks(path.Join(packageDir, file), true)

	// Some templates don't need to be "compiled" (for example base templates)
	for _, chunk := range chunks {
		if strings.HasPrefix(chunk, cmdNocompile) {
			return compiledTemplate{}
		}
	}

	result := compiledTemplate{}
	result.originalFile = file

	chunks = processExtending(packageDir, chunks)

	if len(params.ErrFuncPrefix) == 0 {
		params.ErrFuncPrefix = "TEMPLATEERR"
	}
	if len(params.NoerrFuncPrefix) == 0 {
		params.NoerrFuncPrefix = "TEMPLATE"
	}

	funcName := strings.Replace(file, ".tmpl", "", -1)
	tmplParams := templateParams{
		ErrFuncPrefix:   params.ErrFuncPrefix,
		NoerrFuncPrefix: params.NoerrFuncPrefix,
		FuncName:        funcName,
		TemplateFile:    file,
		EscapeFunc:      "html.EscapeString",
	}

	for _, chunk := range chunks {

		if strings.HasPrefix(chunk, cmdPrefix) {
			// Remove spaces after !#
			chunk = cmdPrefix + strings.TrimSpace(chunk[len(cmdPrefix):])
		}

		if strings.HasPrefix(chunk, cmdArg) {
			tmplParams.Args = append(tmplParams.Args, strings.TrimSpace(chunk[len(cmdArg):]))
		} else if strings.HasPrefix(chunk, cmdGlobal) {
			tmplParams.GlobalCode += strings.TrimSpace(chunk[len(cmdEscape):]) + "\n"
		} else if strings.HasPrefix(chunk, cmdEscape) {
			tmplParams.EscapeFunc = strings.TrimSpace(chunk[len(cmdEscape):])
		} else if strings.HasPrefix(chunk, cmdImport) {
			result.imports = append(result.imports, strings.TrimSpace(chunk[len(cmdImport):]))
		} else if strings.HasPrefix(chunk, cmdReturn) {
			tmplParams.Lines = append(tmplParams.Lines, "return _ftmpl.String(), nil")
		} else if strings.HasPrefix(chunk, cmdErrorif) {
			parts := strings.Split(chunk[len(cmdErrorif):], "???")
			expression := parts[0]
			message := fmt.Sprintf("Error checking %s", expression)
			if len(parts) > 1 {
				message = strings.TrimSpace(parts[1])
			}
			var errBuffer bytes.Buffer
			err := errorTemplate.Execute(&errBuffer, errParams{Expression: expression, Message: message})
			handleLineError(err, chunk)
			tmplParams.Lines = append(tmplParams.Lines, errBuffer.String())
		} else if strings.HasPrefix(chunk, "!#") {
			// Ignore
		} else if strings.HasPrefix(chunk, "!!") {
			chunk = chunk[1:]
			if l, ok := handleTemplateLine(chunk); ok {
				tmplParams.Lines = append(tmplParams.Lines, l)
			}
		} else if strings.HasPrefix(chunk, "!") {
			tmplParams.Lines = append(tmplParams.Lines, prepareCommandLine(chunk[1:]))
		} else {
			if l, ok := handleTemplateLine(chunk); ok {
				tmplParams.Lines = append(tmplParams.Lines, l)
			}
		}
	}

	if len(tmplParams.EscapeFunc) == 0 {
		tmplParams.EscapeFunc = defaultQuoteTemplate
	}

	result.imports = append(result.imports, "\"bytes\"", "\"errors\"", "\"html\"", "\"fmt\"")

	tmplParams.Args = rmDoubleImports(tmplParams.Args)

	var buf bytes.Buffer
	generatorTemplate.Execute(&buf, tmplParams)
	result.functionCode = buf.String()

	return result
}

func prepareCommandLine(chunk string) string {
	chunk = strings.TrimSpace(chunk)
	if (strings.HasPrefix(chunk, "if ") || strings.HasPrefix(chunk, "for ")) && !strings.HasSuffix(chunk, "{") {
		chunk += "{"
	}
	if chunk == "else" {
		chunk = "} else {"
	}
	if chunk == "end" {
		chunk = "}"
	}
	if matches, _ := regexp.Match("else\\s+if.*", []byte(chunk)); matches && !strings.HasSuffix(chunk, "{") {
		chunk = fmt.Sprintf("} %s {", chunk)
	}
	return chunk
}

func handleTemplateLine(chunk string) (string, bool) {
	var buf bytes.Buffer
	err := stringTemplate.Execute(&buf, quote(chunk))
	handleLineError(err, chunk)
	return buf.String(), len(chunk) > 0
}
