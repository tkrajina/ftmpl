package ftmplting

import (
	"strings"
	texttmpl "text/template"
)

var generatorTemplate = texttmpl.Must(texttmpl.New("").Parse(`{{ .GlobalCode }}
// {{ .ErrFuncPrefix }}{{ .FuncName }} evaluates a template {{ .TemplateFile }}
func {{ .ErrFuncPrefix }}{{ .FuncName }}({{ .ArgsJoined }}) (string, error) {
	_template := "{{ .TemplateFile }}"
	_escape := {{ .EscapeFunc }}
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

{{ range .Lines }}{{ . }}
{{ end }}
	return _ftmpl.String(), nil
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

var stringTemplate = texttmpl.Must(texttmpl.New("").Parse(`_w(` + "`" + `{{ . }}` + "`" + `)`))

var newlineTemplate = `_w("\\n")`

var patternTemplate = texttmpl.Must(texttmpl.New("").Parse(`_w(fmt.Sprintf(` + "`" + `{{ .Template }}` + "`" + `, {{ .ArgsJoined }}))`))

type patternTemplateParam struct {
	Template string
	Args     []string
}

func (ptp patternTemplateParam) ArgsJoined() string {
	return strings.Join(ptp.Args, ", ")
}
