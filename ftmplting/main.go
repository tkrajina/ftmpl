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

func saveTemplates(destination string, compiled ...compiledTemplate) {
	fmt.Sprintln("destination=", destination)

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

	_, _ = fOut.WriteString("// package " + packageName + " is generated, do not edit!!!! */\n")
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

func processExtending(pckg string, lines []string) []string {
	extendingTemplate := ""
	for _, line := range lines {
		if strings.HasPrefix(line, cmdExtends) {
			extendingTemplate = strings.TrimSpace(line[len(cmdExtends):])
		}
	}

	if len(extendingTemplate) == 0 {
		return lines
	}

	// Load base template:
	mainTemplateLines := loadTemplateAndGetLines(path.Join(pckg, extendingTemplate)+".tmpl", false)

	// Parse sub template chunks:
	subTemplateCode, subTemplateChunks := loadTemplateSubChunks(lines)

	// Merge them together:
	var result []string

	var subTemplateArgsCode []string
	for _, line := range subTemplateCode {
		if strings.HasPrefix(line, cmdArg) {
			subTemplateArgsCode = append(subTemplateArgsCode, line)
		} else {
			result = append(result, line)
		}
	}

	for _, line := range mainTemplateLines {
		if strings.HasPrefix(line, cmdInclude) {
			subName := strings.TrimSpace(line[len(cmdInclude):])
			if subLines, found := subTemplateChunks[subName]; found {
				for _, subLine := range subLines {
					result = append(result, subLine)
				}
			}
		} else {
			result = append(result, line)
		}
	}

	// This is because we want the subtemplate args to be after the base template args in the function definition:
	for _, line := range subTemplateArgsCode {
		result = append(result, line)
	}

	return result
}

func loadTemplateSubChunks(lines []string) (general []string, subsections map[string][]string) {
	subReached := false
	currentSubName := ""

	general = make([]string, 0)
	subsections = make(map[string][]string)

	for _, line := range lines {
		if strings.Contains(line, cmdSub) {
			subReached = true
			currentSubName = strings.TrimSpace(line[len(cmdSub):])
		} else if subReached {
			if _, found := subsections[currentSubName]; !found {
				subsections[currentSubName] = make([]string, 0)
			}
			subsections[currentSubName] = append(subsections[currentSubName], line)
		} else {
			general = append(general, line)
		}
	}

	return general, subsections
}

func getLines(str string) []string {
	delimiter := "<<<" + getRandomString(15) + ">>>"
	res := ""
	lines := strings.Split(str, "\n")
	for n, line := range lines {
		if n < len(lines)-1 {
			line += "\n"
		}

		//fmt.Printf("line=%s\n", line)
		if len(line) == 0 {
		} else if strings.HasPrefix(line, "!!") {
			res += line
		} else if line[0] == '!' {
			res += delimiter + line + delimiter
		} else {
			line = exclamationMarkFixRegex.ReplaceAllStringFunc(line, func(s string) string {
				return s + "!"
			})
			res += templateReplacementRegex.ReplaceAllStringFunc(line, func(s string) string {
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
				if len(s) > 1 && s[1] == ' ' {
					valueExpr = s[1:]
					placeholder = "%" + string(s[0])
				} else {
					valueExpr = s
				}

				if !forceUnquoted && placeholder == "%s" {
					valueExpr = "_escape(" + valueExpr + ")"
				}

				if strings.HasPrefix(s, "/*") && strings.HasSuffix(s, "*/") {
					return ""
				}

				line := fmt.Sprintf("!_, _ = result.WriteString(fmt.Sprintf(`%s`, %s))", placeholder, valueExpr)
				return delimiter + line + delimiter
			})
		}
	}
	return strings.Split(res, delimiter)
}

// Loads the file and process it to get lines.
//
// Note that "line" is not always a line from the template it is part of the template which will be converted to a line of code.
func loadTemplateAndGetLines(fileName string, keepNoCompile bool) []string {
	str, err := loadFile(fileName)
	HandleError(err, "Error reading "+fileName)

	lines := getLines(str)

	// Process inserted templates:
	var expandedLines []string
	for _, line := range lines {
		if !keepNoCompile && strings.HasPrefix(line, cmdNocompile) {
			// Nothing
		} else if strings.HasPrefix(line, cmdInsert) {
			dir, _ := path.Split(fileName)
			fileToInsert := line[len(cmdInsert):]
			fileToInsert = strings.Trim(strings.TrimSpace(fileToInsert), "\"")
			fileToInsert = path.Join(dir, fileToInsert)
			expandedLines = append(expandedLines, loadTemplateAndGetLines(fileToInsert, false)...)
		} else {
			expandedLines = append(expandedLines, line)
		}
	}

	return expandedLines
}

type convertTemplateParams struct {
	ErrFuncPrefix, NoerrFuncPrefix string
}

func convertTemplate(packageDir, file string, params convertTemplateParams) compiledTemplate {
	lines := loadTemplateAndGetLines(path.Join(packageDir, file), true)

	// Some templates don't need to be "compiled" (for example base templates)
	for _, line := range lines {
		if strings.HasPrefix(line, cmdNocompile) {
			return compiledTemplate{}
		}
	}

	result := compiledTemplate{}
	result.originalFile = file

	lines = processExtending(packageDir, lines)

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

	for _, line := range lines {

		if strings.HasPrefix(line, cmdPrefix) {
			// Remove spaces after !#
			line = cmdPrefix + strings.TrimSpace(line[len(cmdPrefix):])
		}

		if strings.HasPrefix(line, cmdArg) {
			tmplParams.Args = append(tmplParams.Args, strings.TrimSpace(line[len(cmdArg):]))
		} else if strings.HasPrefix(line, cmdGlobal) {
			tmplParams.GlobalCode += strings.TrimSpace(line[len(cmdEscape):]) + "\n"
		} else if strings.HasPrefix(line, cmdEscape) {
			tmplParams.EscapeFunc = strings.TrimSpace(line[len(cmdEscape):])
		} else if strings.HasPrefix(line, cmdImport) {
			result.imports = append(result.imports, strings.TrimSpace(line[len(cmdImport):]))
		} else if strings.HasPrefix(line, cmdReturn) {
			tmplParams.Lines = append(tmplParams.Lines, "return result.String(), nil")
		} else if strings.HasPrefix(line, cmdErrorif) {
			parts := strings.Split(line[len(cmdErrorif):], "???")
			expression := parts[0]
			message := fmt.Sprintf("Error checking %s", expression)
			if len(parts) > 1 {
				message = strings.TrimSpace(parts[1])
			}
			var errBuffer bytes.Buffer
			err := errorTemplate.Execute(&errBuffer, errParams{Expression: expression, Message: message})
			handleLineError(err, line)
			tmplParams.Lines = append(tmplParams.Lines, errBuffer.String())
		} else if strings.HasPrefix(line, "!#") {
			// Ignore
		} else if strings.HasPrefix(line, "!!") {
			line = line[1:]
			if l, ok := handleTemplateLine(line); ok {
				tmplParams.Lines = append(tmplParams.Lines, l)
			}
		} else if strings.HasPrefix(line, "!") {
			tmplParams.Lines = append(tmplParams.Lines, prepareCommandLine(line[1:]))
		} else {
			if l, ok := handleTemplateLine(line); ok {
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

func prepareCommandLine(line string) string {
	line = strings.TrimSpace(line)
	if (strings.HasPrefix(line, "if ") || strings.HasPrefix(line, "for ")) && !strings.HasSuffix(line, "{") {
		line += "{"
	}
	if line == "else" {
		line = "} else {"
	}
	if line == "end" {
		line = "}"
	}
	if matches, _ := regexp.Match("else\\s+if.*", []byte(line)); matches && !strings.HasSuffix(line, "{") {
		line = fmt.Sprintf("} %s {", line)
	}
	return line
}

func handleTemplateLine(line string) (string, bool) {
	var buf bytes.Buffer
	err := stringTemplate.Execute(&buf, quote(line))
	handleLineError(err, line)
	return buf.String(), len(line) > 0
}
