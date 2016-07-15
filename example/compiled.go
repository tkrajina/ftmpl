// package example is generated with ftmpl {{{v0.2.0}}}, do not edit!!!! */
package example

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"os"
)

func init() {
	_ = fmt.Sprintf
	_ = errors.New
	_ = os.Stderr
	_ = html.EscapeString
}

// TMPLERRBaseEmbedded evaluates a template BaseEmbedded.tmpl
func TMPLERRBaseEmbedded(title string) (string, error) {
	_template := "BaseEmbedded.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`
<html>
    <head>
        <title>`)
	_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape(title)))
	_, _ = result.WriteString(`</title>
`)
	_, _ = result.WriteString(`
    </head>
    <body>
`)
	_, _ = result.WriteString(`
`)
	_, _ = result.WriteString(`
    </body>
</html>
`)

	return result.String(), nil
}

// TMPLBaseEmbedded evaluates a template BaseEmbedded.tmpl
func TMPLBaseEmbedded(title string) string {
	html, err := TMPLERRBaseEmbedded(title)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template BaseEmbedded.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRBasicCode evaluates a template BasicCode.tmpl
func TMPLERRBasicCode(s string, num int) (string, error) {
	_template := "BasicCode.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	for i := 0; i < num; i++ {
		_, _ = result.WriteString(fmt.Sprintf(`%d`, i))
		_, _ = result.WriteString(`
`)
	}

	return result.String(), nil
}

// TMPLBasicCode evaluates a template BasicCode.tmpl
func TMPLBasicCode(s string, num int) string {
	html, err := TMPLERRBasicCode(s, num)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template BasicCode.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRBasicEmbeddedCode evaluates a template BasicEmbeddedCode.tmpl
func TMPLERRBasicEmbeddedCode(n int) (string, error) {
	_template := "BasicEmbeddedCode.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	for i := 0; i < n; i++ {
		_, _ = result.WriteString(`i=`)
		_, _ = result.WriteString(fmt.Sprintf(`%d`, i))
		_, _ = result.WriteString(` `)
	}
	_, _ = result.WriteString(`
`)

	return result.String(), nil
}

// TMPLBasicEmbeddedCode evaluates a template BasicEmbeddedCode.tmpl
func TMPLBasicEmbeddedCode(n int) string {
	html, err := TMPLERRBasicEmbeddedCode(n)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template BasicEmbeddedCode.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRBasicIfElse evaluates a template BasicIfElse.tmpl
func TMPLERRBasicIfElse(n int) (string, error) {
	_template := "BasicIfElse.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	if n < 10 {
		_, _ = result.WriteString(fmt.Sprintf(`%d`, n))
		_, _ = result.WriteString(` less than 10`)
	}
	_, _ = result.WriteString(`
`)
	if n > 0 {
		_, _ = result.WriteString(fmt.Sprintf(`%d`, n))
		_, _ = result.WriteString(` biger than 0`)
	}
	_, _ = result.WriteString(`
`)
	if n > 5 {
		_, _ = result.WriteString(fmt.Sprintf(`%d`, n))
		_, _ = result.WriteString(` biger than 5`)
	} else {
		_, _ = result.WriteString(fmt.Sprintf(`%d`, n))
		_, _ = result.WriteString(` smaller than 5`)
	}
	_, _ = result.WriteString(`
`)

	return result.String(), nil
}

// TMPLBasicIfElse evaluates a template BasicIfElse.tmpl
func TMPLBasicIfElse(n int) string {
	html, err := TMPLERRBasicIfElse(n)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template BasicIfElse.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRBasicIfElseif evaluates a template BasicIfElseif.tmpl
func TMPLERRBasicIfElseif(n int) (string, error) {
	_template := "BasicIfElseif.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`
`)
	if n < 10 {
		_, _ = result.WriteString(`n less than 10`)
	} else if n < 100 {
		_, _ = result.WriteString(`n less than 100`)
	} else {
		_, _ = result.WriteString(`n bigger than 100`)
	}
	_, _ = result.WriteString(`
`)

	return result.String(), nil
}

// TMPLBasicIfElseif evaluates a template BasicIfElseif.tmpl
func TMPLBasicIfElseif(n int) string {
	html, err := TMPLERRBasicIfElseif(n)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template BasicIfElseif.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRComparisonWithGoTemplates evaluates a template ComparisonWithGoTemplates.tmpl
func TMPLERRComparisonWithGoTemplates(params TemplateParam) (string, error) {
	_template := "ComparisonWithGoTemplates.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`<html>
    <head>
        <title>`)
	_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape(params.Title)))
	_, _ = result.WriteString(`</title>
    </head>
    <body>
        <h1>`)
	_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape(params.Title)))
	_, _ = result.WriteString(`</h1>
        `)
	if len(params.Subtitle) > 0 {
		_, _ = result.WriteString(`<h2>`)
		_, _ = result.WriteString(fmt.Sprintf(`%v`, params.Subtitle))
		_, _ = result.WriteString(`</h1>`)
	}
	_, _ = result.WriteString(`
        <ul>
            `)
	for _, item := range params.Items {
		_, _ = result.WriteString(`
                <li> `)
		_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape(item)))
		_, _ = result.WriteString(`
            `)
	}
	_, _ = result.WriteString(`
        </ul>
        <p>
            Written `)
	_, _ = result.WriteString(fmt.Sprintf(`%d`, len(params.Items)))
	_, _ = result.WriteString(` items
        </p>
    </body>
</html>
`)

	return result.String(), nil
}

// TMPLComparisonWithGoTemplates evaluates a template ComparisonWithGoTemplates.tmpl
func TMPLComparisonWithGoTemplates(params TemplateParam) string {
	html, err := TMPLERRComparisonWithGoTemplates(params)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template ComparisonWithGoTemplates.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRExtendsEmbedded evaluates a template ExtendsEmbedded.tmpl
func TMPLERRExtendsEmbedded(title string, something int) (string, error) {
	_template := "ExtendsEmbedded.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`
`)
	_, _ = result.WriteString(`

`)
	_, _ = result.WriteString(`
<html>
    <head>
        <title>`)
	_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape(title)))
	_, _ = result.WriteString(`</title>
`)
	_, _ = result.WriteString(`
<script>
alert("included")
</script>

`)
	_, _ = result.WriteString(`
    </head>
    <body>
`)
	_, _ = result.WriteString(`
<h1>Body!</h1>
`)
	_, _ = result.WriteString(`
`)
	_, _ = result.WriteString(`
    </body>
</html>
`)

	return result.String(), nil
}

// TMPLExtendsEmbedded evaluates a template ExtendsEmbedded.tmpl
func TMPLExtendsEmbedded(title string, something int) string {
	html, err := TMPLERRExtendsEmbedded(title, something)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template ExtendsEmbedded.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRNoncodeLineWithExclamationMark evaluates a template NoncodeLineWithExclamationMark.tmpl
func TMPLERRNoncodeLineWithExclamationMark() (string, error) {
	_template := "NoncodeLineWithExclamationMark.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`!s1 := "This line is not a code line"
`)
	s2 := "This *is* a line of code"
	_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape(s2)))
	_, _ = result.WriteString(`
`)

	return result.String(), nil
}

// TMPLNoncodeLineWithExclamationMark evaluates a template NoncodeLineWithExclamationMark.tmpl
func TMPLNoncodeLineWithExclamationMark() string {
	html, err := TMPLERRNoncodeLineWithExclamationMark()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template NoncodeLineWithExclamationMark.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRWithEndInsteadOfBrackets evaluates a template WithEndInsteadOfBrackets.tmpl
func TMPLERRWithEndInsteadOfBrackets() (string, error) {
	_template := "WithEndInsteadOfBrackets.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	for i := 0; i < 5; i++ {
		_, _ = result.WriteString(`i=`)
		_, _ = result.WriteString(fmt.Sprintf(`%d`, i))
		_, _ = result.WriteString(`
`)
	}

	return result.String(), nil
}

// TMPLWithEndInsteadOfBrackets evaluates a template WithEndInsteadOfBrackets.tmpl
func TMPLWithEndInsteadOfBrackets() string {
	html, err := TMPLERRWithEndInsteadOfBrackets()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template WithEndInsteadOfBrackets.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRWithExclamationMark evaluates a template WithExclamationMark.tmpl
func TMPLERRWithExclamationMark() (string, error) {
	_template := "WithExclamationMark.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`Something here `)
	_, _ = result.WriteString(fmt.Sprintf(`%d`, 5))
	_, _ = result.WriteString(`! And something `)
	_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape("here")))
	_, _ = result.WriteString(`.
And something here: `)
	_, _ = result.WriteString(fmt.Sprintf(`%t`, true))
	_, _ = result.WriteString(`!
`)

	return result.String(), nil
}

// TMPLWithExclamationMark evaluates a template WithExclamationMark.tmpl
func TMPLWithExclamationMark() string {
	html, err := TMPLERRWithExclamationMark()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template WithExclamationMark.tmpl:" + err.Error())
	}
	return html
}

type Argument struct {
	Aaa string
	Bbb int
}

// TMPLERRWithGlobalDeclaration evaluates a template WithGlobalDeclaration.tmpl
func TMPLERRWithGlobalDeclaration(arg Argument) (string, error) {
	_template := "WithGlobalDeclaration.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`Aaa=`)
	_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape(arg.Aaa)))
	_, _ = result.WriteString(`
Bbb=`)
	_, _ = result.WriteString(fmt.Sprintf(`%d`, arg.Bbb))
	_, _ = result.WriteString(`
`)

	return result.String(), nil
}

// TMPLWithGlobalDeclaration evaluates a template WithGlobalDeclaration.tmpl
func TMPLWithGlobalDeclaration(arg Argument) string {
	html, err := TMPLERRWithGlobalDeclaration(arg)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template WithGlobalDeclaration.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRWithInsert evaluates a template WithInsert.tmpl
func TMPLERRWithInsert(a int) (string, error) {
	_template := "WithInsert.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`Will insert something here: `)
	_, _ = result.WriteString(`a=`)
	_, _ = result.WriteString(fmt.Sprintf(`%d`, a))
	_, _ = result.WriteString(`
`)
	_, _ = result.WriteString(`
`)

	return result.String(), nil
}

// TMPLWithInsert evaluates a template WithInsert.tmpl
func TMPLWithInsert(a int) string {
	html, err := TMPLERRWithInsert(a)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template WithInsert.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRWithPercentage evaluates a template WithPercentage.tmpl
func TMPLERRWithPercentage(str string) (string, error) {
	_template := "WithPercentage.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`%, str=`)
	_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape(str)))
	_, _ = result.WriteString(`
%, str=`)
	_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape(fmt.Sprintf("aaa%sccc", "bbb"))))
	_, _ = result.WriteString(`
`)

	return result.String(), nil
}

// TMPLWithPercentage evaluates a template WithPercentage.tmpl
func TMPLWithPercentage(str string) string {
	html, err := TMPLERRWithPercentage(str)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template WithPercentage.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRbasic evaluates a template basic.tmpl
func TMPLERRbasic(str string, num int) (string, error) {
	_template := "basic.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`String:`)
	_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape(str)))
	_, _ = result.WriteString(`
Unescaped:`)
	_, _ = result.WriteString(fmt.Sprintf(`%s`, str))
	_, _ = result.WriteString(`
Num:`)
	_, _ = result.WriteString(fmt.Sprintf(`%d`, num))
	_, _ = result.WriteString(`
`)

	return result.String(), nil
}

// TMPLbasic evaluates a template basic.tmpl
func TMPLbasic(str string, num int) string {
	html, err := TMPLERRbasic(str, num)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template basic.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRextends evaluates a template extends.tmpl
func TMPLERRextends(title string, something int) (string, error) {
	_template := "extends.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`
`)
	_, _ = result.WriteString(`<html>
    <head>
        <title>`)
	_, _ = result.WriteString(fmt.Sprintf(`%s`, _escape(title)))
	_, _ = result.WriteString(`</title>
`)
	_, _ = result.WriteString(`<script>
alert("included")
</script>

`)
	_, _ = result.WriteString(`    </head>
    <body>
`)
	_, _ = result.WriteString(`<h1>Body!</h1>
`)
	_, _ = result.WriteString(`    </body>
</html>
`)

	return result.String(), nil
}

// TMPLextends evaluates a template extends.tmpl
func TMPLextends(title string, something int) string {
	html, err := TMPLERRextends(title, something)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template extends.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRreturn evaluates a template return.tmpl
func TMPLERRreturn(a int) (string, error) {
	_template := "return.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	_, _ = result.WriteString(`a is `)
	_, _ = result.WriteString(fmt.Sprintf(`%d`, a))
	_, _ = result.WriteString(`
`)
	return result.String(), nil
	_, _ = result.WriteString(`
This line is ignored
`)

	return result.String(), nil
}

// TMPLreturn evaluates a template return.tmpl
func TMPLreturn(a int) string {
	html, err := TMPLERRreturn(a)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template return.tmpl:" + err.Error())
	}
	return html
}
