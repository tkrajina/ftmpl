package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	texttmpl "text/template"
)

const CMDLINE_TARGETGO = "targetgo"
const CMDLINE_TARGETDIR = "targetdir"

type AppParams struct {
	targetDir    string
	targetGoFile string
}

const CMD_ARG = "!#arg"
const CMD_ESCAPE = "!#escape"
const CMD_IMPORT = "!#import"
const CMD_ERROIF = "!#errif"
const CMD_SUB = "!#sub"
const CMD_EXTENDS = "!#extends"
const CMD_INCLUDE = "!#include"
const CMD_GLOBAL = "!#global"

var templateReplacementRegex = regexp.MustCompile("{{.*?}}")

var generatorTemplate = texttmpl.Must(texttmpl.New("").Parse(`// Generated code, do not edit!!!!
{{ .GlobalCode }}
func {{ .ErrFuncPrefix }}{{ .FuncName }}({{ .ArgsJoined }}) (string, error) {
	__template__ := "{{ .TemplateFile }}"
	_ = __template__
	__escape__ := {{ .EscapeFunc }}
	_ = __escape__
	var result bytes.Buffer
{{ range .Lines }}{{ . }}
{{ end }}
	return result.String(), nil
}


func {{ .NoerrFuncPrefix }}{{ .FuncName }}({{ .ArgsJoined }}) string {
	html, err := {{ .ErrFuncPrefix }}{{ .FuncName }}({{ .ArgNamesJoined }})
	if err != nil {
		os.Stderr.WriteString("Error running template {{ .TemplateFile }}:" + err.Error())
	}
	return html
}`))

type TemplateParams struct {
	GlobalCode      string
	ErrFuncPrefix   string
	NoerrFuncPrefix string
	TemplateFile    string
	EscapeFunc      string
	FuncName        string
	Args            []string
	Lines           []string
}

func (tp TemplateParams) ArgsJoined() string {
	return strings.Join(tp.Args, ", ")
}

func (tp TemplateParams) ArgNamesJoined() string {
	argNames := make([]string, 0)
	for _, arg := range tp.Args {
		argNames = append(argNames, strings.Split(arg, " ")[0])
	}
	return strings.Join(argNames, ", ")
}

func (tp *TemplateParams) addComment(line string) {
	tp.Lines = append(tp.Lines, "/* "+strings.TrimSpace(strings.Replace(line, "*/", "* /", -1))+" */")
}

var initFunctionTemplate = `func init() {
	_ = fmt.Sprintf
	_ = errors.New
	_ = os.Stderr
	_ = html.EscapeString
}`

var defaultQuoteTemplate = `func(str string) string {
	return str
}`

var errorTemplate = texttmpl.Must(texttmpl.New("").Parse(`if {{.Expression}} {
	return "", errors.New({{.Message}})
}`))

type ErrParams struct {
	Expression, Message string
}

var stringTemplate = texttmpl.Must(texttmpl.New("").Parse(`result.WriteString(` + "`" + `{{ . }}` + "`" + `)`))

var patternTemplate = texttmpl.Must(texttmpl.New("").Parse(`result.WriteString(fmt.Sprintf(` + "`" + `{{ .Template }}` + "`" + `, {{ .ArgsJoined }}))`))

var newlineTemplate = `result.WriteString("\\n")`

type PatternTemplateParam struct {
	Template string
	Args     []string
}

func (ptp PatternTemplateParam) ArgsJoined() string {
	return strings.Join(ptp.Args, ", ")
}

func main() {
	var appParams AppParams
	flag.StringVar(&(appParams.targetGoFile), CMDLINE_TARGETGO, "", "Merge the result in this go file")
	flag.StringVar(&(appParams.targetDir), CMDLINE_TARGETDIR, "", "Save the result in this directory")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Need one source directory with templates")
		os.Exit(1)
	}
	sourceDir := flag.Arg(0)

	if len(appParams.targetDir) > 0 && len(appParams.targetGoFile) > 0 {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Only %s or %s (or none) can be used, not both", CMDLINE_TARGETGO, CMDLINE_TARGETDIR))
		os.Exit(1)
	}

	if len(appParams.targetDir) == 0 && len(appParams.targetGoFile) == 0 {
		appParams.targetDir = sourceDir
	}

	var compiledTemplates []compiledTemplate

	fileInfos, err := ioutil.ReadDir(sourceDir)
	handleError(err, "Error listing directory")

	files := make([]string, 0)
	for _, fileInfo := range fileInfos {
		files = append(files, fileInfo.Name())
	}
	sort.Strings(files)

	for _, fileName := range files {
		if strings.HasSuffix(fileName, ".tmpl") {
			compiled := convertTemplate(sourceDir, fileName)
			if len(appParams.targetDir) > 0 {
				_, fileName := getLastPathElements(fileName)
				fileName = strings.Replace(fileName, ".tmpl", ".go", -1)
				fileName = path.Join(appParams.targetDir, fileName)
				if len(compiled.functionCode) > 0 {
					saveTemplates(fileName, compiled)
				}
			} else {
				compiledTemplates = append(compiledTemplates, compiled)
			}
		}
	}

	if len(appParams.targetGoFile) > 0 {
		saveTemplates(appParams.targetGoFile, compiledTemplates...)
	}
}

func saveTemplates(destination string, compiled ...compiledTemplate) {
	fmt.Sprintln("destination=", destination)

	imports := make([]string, 0)
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
	handleError(err, "Error creating file")
	defer fOut.Close()

	fOut.WriteString("package " + packageName + "\n\n")
	fOut.WriteString("import (\n")
	for _, i := range imports {
		fOut.WriteString(i + "\n")
	}
	fOut.WriteString(")\n\n")
	fOut.WriteString(initFunctionTemplate)
	fOut.WriteString("\n\n")

	for _, c := range compiled {
		fOut.WriteString("\n\n")
		fOut.WriteString(c.functionCode)

		fmt.Printf("%-15s -> %-15s\n", c.originalFile, destination)
	}

	cmd := exec.Command("go", "fmt", destination)
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, strings.Join(cmd.Args, " "))
		fmt.Fprintln(os.Stderr, string(bytes))
		handleError(err, "Error formatting resulting go file")
	}
}

func processExtending(pckg string, lines []string) []string {
	extendingTemplate := ""
	for _, line := range lines {
		if strings.HasPrefix(line, CMD_EXTENDS) {
			extendingTemplate = strings.TrimSpace(line[len(CMD_EXTENDS):])
		}
	}

	if len(extendingTemplate) == 0 {
		return lines
	}

	// Load base template:
	mainTemplateLines := loadTemplateAndGetLines(path.Join(pckg, extendingTemplate) + ".tmpl")

	// Parse sub template chunks:
	subTemplateCode, subTemplateChunks := loadTemplateSubChunks(lines)

	// Merge them together:
	result := make([]string, 0)

	subTemplateArgsCode := make([]string, 0)
	for _, line := range subTemplateCode {
		if strings.HasPrefix(line, CMD_ARG) {
			subTemplateArgsCode = append(subTemplateArgsCode, line)
		} else {
			result = append(result, line)
		}
	}

	for _, line := range mainTemplateLines {
		if strings.HasPrefix(line, CMD_INCLUDE) {
			subName := strings.TrimSpace(line[len(CMD_INCLUDE):])
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
		if strings.Contains(line, CMD_SUB) {
			subReached = true
			currentSubName = strings.TrimSpace(line[len(CMD_SUB):])
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

func getLastPathElements(path string) (string, string) {
	absPath, err := filepath.Abs(path)
	handleError(err, "Error getting absolute path from "+path)
	pathParts := strings.Split(absPath, string(os.PathSeparator))

	elements := make([]string, 0)
	for _, part := range pathParts {
		part = strings.TrimSpace(part)
		if len(part) > 0 {
			elements = append(elements, part)
		}
	}

	if len(pathParts) == 0 {
		return "", ""
	}

	if len(elements) == 1 {
		return "", elements[len(elements)-1]
	}

	return elements[len(elements)-2], elements[len(elements)-1]
}

type compiledTemplate struct {
	imports      []string
	originalFile string
	functionCode string
}

const RANDOM_STRING_CHARS = "qwertyuiopasdfghjklzxcvbnm1234567890"

func getRandomString(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result += string(RANDOM_STRING_CHARS[rand.Int()%len(RANDOM_STRING_CHARS)])
	}

	return result
}

// Loads the file and process it to get lines.
//
// Note that "line" is not always a line from the template it is part of the template which will be converted to a line of code.
func loadTemplateAndGetLines(fileName string) []string {
	f, err := os.Open(path.Join(fileName))
	handleError(err, "Error opening "+fileName)
	defer f.Close()

	byts, err := ioutil.ReadAll(f)
	handleError(err, "Error reading "+fileName)
	str := string(byts)

	lineDelimiter := getRandomString(15)
	str = strings.Replace(str, "\n", "\n"+lineDelimiter, -1)

	r := regexp.MustCompile("{{!.*?}}")
	str = r.ReplaceAllStringFunc(str, func(s string) string {
		// In this case there is no \n before the line delimiter:
		result := strings.TrimSpace(s[3 : len(s)-2])

		if (strings.HasPrefix(result, "if ") || strings.HasPrefix(result, "for ")) && !strings.HasSuffix(result, "{") {
			result += "{"
		}
		if result == "else" {
			result = "} else {"
		}
		if result == "end" {
			result = "}"
		}
		if matches, _ := regexp.Match("else\\s+if.*", []byte(result)); matches && !strings.HasSuffix(result, "{") {
			result = fmt.Sprintf("} %s {", result)
		}

		result = lineDelimiter + "!" + result + lineDelimiter

		return result
	})

	return strings.Split(str, lineDelimiter)
}

func convertTemplate(packageDir, file string) compiledTemplate {
	lines := loadTemplateAndGetLines(path.Join(packageDir, file))

	result := compiledTemplate{}
	result.originalFile = file

	_, pckg := getLastPathElements(packageDir)
	lines = processExtending(pckg, lines)

	funcName := strings.Replace(file, ".tmpl", "", -1)
	params := TemplateParams{
		ErrFuncPrefix:   "TE__",
		NoerrFuncPrefix: "T__",
		FuncName:        funcName,
		TemplateFile:    file,
		EscapeFunc:      "html.EscapeString",
	}

	for n, line := range lines {
		_ = n
		if strings.HasPrefix(line, CMD_ARG) {
			params.Args = append(params.Args, strings.TrimSpace(line[len(CMD_ARG):]))
		} else if strings.HasPrefix(line, CMD_GLOBAL) {
			params.GlobalCode += strings.TrimSpace(line[len(CMD_ESCAPE):]) + "\n"
		} else if strings.HasPrefix(line, CMD_ESCAPE) {
			params.EscapeFunc = strings.TrimSpace(line[len(CMD_ESCAPE):])
		} else if strings.HasPrefix(line, CMD_IMPORT) {
			result.imports = append(result.imports, strings.TrimSpace(line[len(CMD_IMPORT):]))
		} else if strings.HasPrefix(line, CMD_ERROIF) {
			params.addComment(line)
			parts := strings.Split(line[len(CMD_ERROIF):], "???")
			expression := parts[0]
			message := fmt.Sprintf("Error checking %s", expression)
			if len(parts) > 1 {
				message = strings.TrimSpace(parts[1])
			}
			var errBuffer bytes.Buffer
			err := errorTemplate.Execute(&errBuffer, ErrParams{Expression: expression, Message: message})
			handleLineError(err, line)
			params.Lines = append(params.Lines, errBuffer.String())
		} else if strings.HasPrefix(line, "!#") {
			params.addComment(line)
		} else if strings.HasPrefix(line, "!!") {
			line = line[1:]
			params.addComment(line)
			params.Lines = append(params.Lines, handleTemplateLine(line))
		} else if strings.HasPrefix(line, "!") {
			params.addComment(line)
			params.Lines = append(params.Lines, line[1:])
		} else {
			params.addComment(line)
			params.Lines = append(params.Lines, handleTemplateLine(line))
		}
	}

	if len(params.EscapeFunc) == 0 {
		params.EscapeFunc = defaultQuoteTemplate
	}

	result.imports = append(result.imports, "\"bytes\"", "\"errors\"", "\"html\"", "\"fmt\"")

	params.Args = rmDoubleImports(params.Args)

	var buf bytes.Buffer
	generatorTemplate.Execute(&buf, params)
	result.functionCode = buf.String()

	return result
}

func handleTemplateLine(line string) string {
	if !strings.Contains(line, "{{") {
		var buf bytes.Buffer
		err := stringTemplate.Execute(&buf, quote(line))
		handleLineError(err, line)
		return buf.String()
	}

	// We need to change "%" to "%%" (otherwise it messes with replacements) but we need to leave them in the template
	// expressions {{d 15 % 3 }}!
	percentageCharReplacement := getRandomString(10)

	params := PatternTemplateParam{}
	str := templateReplacementRegex.ReplaceAllStringFunc(line, func(s string) string {
		placeholder := "%v"
		var valueExpr string

		forceUnquoted := false
		s = strings.TrimSpace(s[2 : len(s)-2])
		if s[0] == '=' {
			forceUnquoted = true
			s = s[1:]
		}

		if len(s) > 0 && s[1] == ' ' {
			valueExpr = s[1:]
			placeholder = "%" + string(s[0])
		} else {
			valueExpr = s
		}

		if !forceUnquoted && placeholder == "%s" {
			valueExpr = "__escape__(" + valueExpr + ")"
		}

		if strings.HasPrefix(s, "/*") && strings.HasSuffix(s, "*/") {
			return ""
		}

		params.Args = append(params.Args, valueExpr)
		return strings.Replace(placeholder, "%", percentageCharReplacement, -1)
	})
	params.Template = quote(str)

	var buf bytes.Buffer
	err := patternTemplate.Execute(&buf, params)
	handleLineError(err, line)

	result := buf.String()
	result = strings.Replace(result, "%", "%%", -1)
	result = strings.Replace(result, percentageCharReplacement, "%", -1)
	return result
}

func quote(s string) string {
	return strings.Replace(s, "`", "`+\"`\"+`", -1)
}

func handleError(err error, message string) {
	if err == nil {
		return
	}
	panic(message + ":" + err.Error())
}

func handleLineError(err error, line string) {
	if err == nil {
		return
	}

	panic("Error " + err.Error() + " in line:" + line)
}

func rmDoubleImports(list []string) []string {
	result := make([]string, 0)
	existing := make(map[string]bool)
	for _, i := range list {
		if _, ok := existing[i]; !ok {
			result = append(result, i)
		}
		existing[i] = true
	}
	return result
}
