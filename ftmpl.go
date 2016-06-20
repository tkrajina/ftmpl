package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	texttmpl "text/template"
)

const (
	cmdlineTargetGo      = "targetgo"
	cmdlineTargetDir     = "targetdir"
	cmdlineFuncPrefix    = "prefix"
	cmdlineFuncPrefixErr = "prefixerr"
)

type appParams struct {
	targetDir     string
	targetGoFile  string
	funcPrefix    string
	funcPrefixErr string
}

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

const randomStringChars = "qwertyuiopasdfghjklzxcvbnm1234567890"

var templateReplacementRegex = regexp.MustCompile("{{.*?}}")

var generatorTemplate = texttmpl.Must(texttmpl.New("").Parse(`{{ .GlobalCode }}
// {{ .ErrFuncPrefix }}{{ .FuncName }} evaluates a template {{ .TemplateFile }}
func {{ .ErrFuncPrefix }}{{ .FuncName }}({{ .ArgsJoined }}) (string, error) {
	_template := "{{ .TemplateFile }}"
	_ = _template
	_escape := {{ .EscapeFunc }}
	_ = _escape
	var result bytes.Buffer
{{ range .Lines }}{{ . }}
{{ end }}
	return result.String(), nil
}

// {{ .NoerrFuncPrefix }}{{ .FuncName }} evaluates a template {{ .TemplateFile }}
func {{ .NoerrFuncPrefix }}{{ .FuncName }}({{ .ArgsJoined }}) string {
	html, err := {{ .ErrFuncPrefix }}{{ .FuncName }}({{ .ArgNamesJoined }})
	if err != nil {
		_ ,_ = os.Stderr.WriteString("Error running template {{ .TemplateFile }}:" + err.Error())
	}
	return html
}`))

type templateParams struct {
	GlobalCode      string
	ErrFuncPrefix   string
	NoerrFuncPrefix string
	TemplateFile    string
	EscapeFunc      string
	FuncName        string
	Args            []string
	Lines           []string
}

func (tp templateParams) ArgsJoined() string {
	return strings.Join(tp.Args, ", ")
}

func (tp templateParams) ArgNamesJoined() string {
	var argNames []string
	for _, arg := range tp.Args {
		argNames = append(argNames, strings.Split(arg, " ")[0])
	}
	return strings.Join(argNames, ", ")
}

func (tp *templateParams) addComment(filename, line string) {
	comment := fmt.Sprintf("//%s: %s", filename, strings.TrimRight(line, "\n\r \t"))
	tp.Lines = append(tp.Lines, comment)
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

type errParams struct {
	Expression, Message string
}

var stringTemplate = texttmpl.Must(texttmpl.New("").Parse(`_, _ = result.WriteString(` + "`" + `{{ . }}` + "`" + `)`))

var patternTemplate = texttmpl.Must(texttmpl.New("").Parse(`_, _ = result.WriteString(fmt.Sprintf(` + "`" + `{{ .Template }}` + "`" + `, {{ .ArgsJoined }}))`))

var newlineTemplate = `_, _ = result.WriteString("\\n")`

type patternTemplateParam struct {
	Template string
	Args     []string
}

func (ptp patternTemplateParam) ArgsJoined() string {
	return strings.Join(ptp.Args, ", ")
}

func main() {
	var ap appParams
	flag.StringVar(&(ap.targetGoFile), cmdlineTargetGo, "", "Merge the result in this go file")
	flag.StringVar(&(ap.targetDir), cmdlineTargetDir, "", "Save the result in this directory")
	flag.StringVar(&(ap.funcPrefix), cmdlineFuncPrefix, "TMPL", "Prefix to be used with generated functions")
	flag.StringVar(&(ap.funcPrefixErr), cmdlineFuncPrefixErr, "TMPLERR", "Prefix to be used with generated functions (returning error)")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Need one source directory with templates")
		os.Exit(1)
	}
	sourceDir := flag.Arg(0)

	if len(ap.targetDir) > 0 && len(ap.targetGoFile) > 0 {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Only %s or %s (or none) can be used, not both", cmdlineTargetGo, cmdlineTargetDir))
		os.Exit(1)
	}

	if len(ap.targetDir) == 0 && len(ap.targetGoFile) == 0 {
		ap.targetDir = sourceDir
	}

	var compiledTemplates []compiledTemplate

	fileInfos, err := ioutil.ReadDir(sourceDir)
	handleError(err, "Error listing directory")

	var files []string
	for _, fileInfo := range fileInfos {
		files = append(files, fileInfo.Name())
	}
	sort.Strings(files)

	for _, fileName := range files {
		if strings.HasSuffix(fileName, ".tmpl") {

			params := convertTemplateParams{
				NoerrFuncPrefix: ap.funcPrefix,
				ErrFuncPrefix:   ap.funcPrefixErr,
			}
			compiled := convertTemplate(sourceDir, fileName, params)
			if len(ap.targetDir) > 0 {
				_, fileName := getLastPathElements(fileName)
				fileName = strings.Replace(fileName, ".tmpl", ".go", -1)
				fileName = path.Join(ap.targetDir, fileName)
				if len(compiled.functionCode) > 0 {
					saveTemplates(fileName, compiled)
				}
			} else {
				compiledTemplates = append(compiledTemplates, compiled)
			}
		}
	}

	if len(ap.targetGoFile) > 0 {
		saveTemplates(ap.targetGoFile, compiledTemplates...)
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
	handleError(err, "Error creating file")
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

func parseGoFmtErrorOutput(out []byte) {
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		r := regexp.MustCompile("^(.*?):(\\d+):(\\d+):(.*)$")
		groups := r.FindStringSubmatch(line)
		fmt.Fprint(os.Stderr, line, "\n")
		if len(groups) > 0 {
			filename := groups[1]
			line, _ := strconv.ParseInt(groups[2], 10, 64)
			column, _ := strconv.ParseInt(groups[3], 10, 64)
			msg := strings.TrimSpace(groups[4])
			printTemplateErrDetails(filename, int(line), int(column), msg)
		}
	}
}

func printTemplateErrDetails(filename string, line, column int, msg string) {
	contents, err := loadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "> Cannot open %s: %s", filename, err.Error())
		return
	}

	lines := strings.Split(contents, "\n")

	context := 5
	fromLine := int(math.Max(float64(0), float64(line-context)))
	toLine := int(math.Min(float64(line+context), float64(len(lines)-1)))
	for i := fromLine; i < toLine; i++ {
		currLine := "   "
		if i+1 == line {
			currLine = "-->"
		}
		prefix := fmt.Sprintf("%s %5d:    ", currLine, i+1)
		fmt.Fprintln(os.Stderr, prefix+lines[i])
		if i+1 == line {
			fmt.Fprintln(os.Stderr, strings.Repeat(" ", len(prefix)+column-1)+"^")
		}
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

func getLastPathElements(path string) (string, string) {
	absPath, err := filepath.Abs(path)
	handleError(err, "Error getting absolute path from "+path)
	pathParts := strings.Split(absPath, string(os.PathSeparator))

	var elements []string
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

func getRandomString(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result += string(randomStringChars[rand.Int()%len(randomStringChars)])
	}

	return result
}

// Loads the file and process it to get lines.
//
// Note that "line" is not always a line from the template it is part of the template which will be converted to a line of code.
func loadTemplateAndGetLines(fileName string, keepNoCompile bool) []string {
	str, err := loadFile(fileName)
	handleError(err, "Error reading "+fileName)

	lineDelimiter := getRandomString(15)
	str = strings.Replace(str, "\n", "\n"+lineDelimiter, -1)

	r := regexp.MustCompile("{{!.*?}}")
	str = r.ReplaceAllStringFunc(str, func(s string) string {
		// In this case there is no \n before the line delimiter:
		result := strings.TrimSpace(s[3 : len(s)-2])

		result = lineDelimiter + "!" + result + lineDelimiter

		return result
	})

	lines := strings.Split(str, lineDelimiter)

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
			tmplParams.addComment(file, line)
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
			tmplParams.addComment(file, line)
		} else if strings.HasPrefix(line, "!!") {
			line = line[1:]
			tmplParams.addComment(file, line)
			tmplParams.Lines = append(tmplParams.Lines, handleTemplateLine(line))
		} else if strings.HasPrefix(line, "!") {
			tmplParams.addComment(file, line)
			tmplParams.Lines = append(tmplParams.Lines, prepareCommandLine(line[1:]))
		} else {
			tmplParams.addComment(file, line)
			tmplParams.Lines = append(tmplParams.Lines, handleTemplateLine(line))
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

	params := patternTemplateParam{}
	str := templateReplacementRegex.ReplaceAllStringFunc(line, func(s string) string {
		s = strings.TrimSpace(s[2 : len(s)-2])
		if len(s) == 0 {
			return ""
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

		params.Args = append(params.Args, strings.Replace(valueExpr, "%", percentageCharReplacement, -1))
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
	var result []string
	existing := make(map[string]bool)
	for _, i := range list {
		if _, ok := existing[i]; !ok {
			result = append(result, i)
		}
		existing[i] = true
	}
	return result
}

func loadFile(filename string) (string, error) {
	f, err := os.Open(path.Join(filename))
	if err != nil {
		return "", fmt.Errorf("Error reading %s: %s", filename, err.Error())
	}
	defer f.Close()

	byts, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("Error reading %s: %s", filename, err.Error())
	}

	return string(byts), nil
}
