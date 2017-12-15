// Package example is generated with ftmpl {{{v0.3.1}}}, do not edit!!!! */
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`
<html>
    <head>
        <title>`)
	_w(fmt.Sprintf(`%s`, _escape(title)))
	_w(`</title>
`)
	_w(`
    </head>
    <body>
`)
	_w(`
`)
	_w(`
    </body>
</html>
`)

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	for i := 0; i < num; i++ {
		_w(fmt.Sprintf(`%d`, i))
		_w(`
`)
	}

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	for i := 0; i < n; i++ {
		_w(`i=`)
		_w(fmt.Sprintf(`%d`, i))
		_w(` `)
	}
	_w(`
`)

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	if n < 10 {
		_w(fmt.Sprintf(`%d`, n))
		_w(` less than 10`)
	}
	_w(`
`)
	if n > 0 {
		_w(fmt.Sprintf(`%d`, n))
		_w(` biger than 0`)
	}
	_w(`
`)
	if n > 5 {
		_w(fmt.Sprintf(`%d`, n))
		_w(` biger than 5`)
	} else {
		_w(fmt.Sprintf(`%d`, n))
		_w(` smaller than 5`)
	}
	_w(`
`)

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`
`)
	if n < 10 {
		_w(`n less than 10`)
	} else if n < 100 {
		_w(`n less than 100`)
	} else {
		_w(`n bigger than 100`)
	}
	_w(`
`)

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`<html>
    <head>
        <title>`)
	_w(fmt.Sprintf(`%s`, _escape(params.Title)))
	_w(`</title>
    </head>
    <body>
        <h1>`)
	_w(fmt.Sprintf(`%s`, _escape(params.Title)))
	_w(`</h1>
        `)
	if len(params.Subtitle) > 0 {
		_w(`<h2>`)
		_w(fmt.Sprintf(`%v`, params.Subtitle))
		_w(`</h1>`)
	}
	_w(`
        <ul>
            `)
	for _, item := range params.Items {
		_w(`
                <li> `)
		_w(fmt.Sprintf(`%s`, _escape(item)))
		_w(`
            `)
	}
	_w(`
        </ul>
        <p>
            Written `)
	_w(fmt.Sprintf(`%d`, len(params.Items)))
	_w(` items
        </p>
    </body>
</html>
`)

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`
<html>
    <head>
        <title>`)
	_w(fmt.Sprintf(`%s`, _escape(title)))
	_w(`</title>
`)
	_w(`
<script>
alert("included")
</script>

`)
	_w(`
    </head>
    <body>
`)
	_w(`
<h1>Body!</h1>
`)
	_w(`
`)
	_w(`
    </body>
</html>
`)

	return _ftmpl.String(), nil
}

// TMPLExtendsEmbedded evaluates a template ExtendsEmbedded.tmpl
func TMPLExtendsEmbedded(title string, something int) string {
	html, err := TMPLERRExtendsEmbedded(title, something)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template ExtendsEmbedded.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRFmtFormat evaluates a template FmtFormat.tmpl
func TMPLERRFmtFormat() (string, error) {
	_template := "FmtFormat.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`A simple int:`)
	_w(fmt.Sprintf(`%d`, 10))
	_w(`
A number:`)
	_w(fmt.Sprintf(`%1.2f`, 2.3333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333))
	_w(`
A padded string:`)
	_w(fmt.Sprintf(`% 10s`, _escape("padded")))
	_w(`
A padded string #2:`)
	_w(fmt.Sprintf(`% 10s`, _escape("&&&&")))
	_w(`
A padded string #3:`)
	_w(fmt.Sprintf(`% 10s`, "&&&&"))
	_w(`
A +v formatting: `)
	_w(fmt.Sprintf(`%+v`, []string{"aaa"}))
	_w(`
`)

	return _ftmpl.String(), nil
}

// TMPLFmtFormat evaluates a template FmtFormat.tmpl
func TMPLFmtFormat() string {
	html, err := TMPLERRFmtFormat()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template FmtFormat.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRInsertWithSubExtends evaluates a template InsertWithSubExtends.tmpl
func TMPLERRInsertWithSubExtends() (string, error) {
	_template := "InsertWithSubExtends.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`<html>
`)
	_w(`aaa
`)
	_w(`sub-something
`)
	_w(`bbb
`)
	_w(`
</html>
`)

	return _ftmpl.String(), nil
}

// TMPLInsertWithSubExtends evaluates a template InsertWithSubExtends.tmpl
func TMPLInsertWithSubExtends() string {
	html, err := TMPLERRInsertWithSubExtends()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template InsertWithSubExtends.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRNoncodeLineWithExclamationMark evaluates a template NoncodeLineWithExclamationMark.tmpl
func TMPLERRNoncodeLineWithExclamationMark() (string, error) {
	_template := "NoncodeLineWithExclamationMark.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`!s1 := "This line is not a code line"
`)
	s2 := "This *is* a line of code"
	_w(fmt.Sprintf(`%s`, _escape(s2)))
	_w(`
`)

	return _ftmpl.String(), nil
}

// TMPLNoncodeLineWithExclamationMark evaluates a template NoncodeLineWithExclamationMark.tmpl
func TMPLNoncodeLineWithExclamationMark() string {
	html, err := TMPLERRNoncodeLineWithExclamationMark()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template NoncodeLineWithExclamationMark.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRWithDirectWriting evaluates a template WithDirectWriting.tmpl
func TMPLERRWithDirectWriting() (string, error) {
	_template := "WithDirectWriting.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`This is `)
	_ftmpl.WriteString("Written directly to ftmplresult")
	_w("!")
	_w(`
`)

	return _ftmpl.String(), nil
}

// TMPLWithDirectWriting evaluates a template WithDirectWriting.tmpl
func TMPLWithDirectWriting() string {
	html, err := TMPLERRWithDirectWriting()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template WithDirectWriting.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRWithEndInsteadOfBrackets evaluates a template WithEndInsteadOfBrackets.tmpl
func TMPLERRWithEndInsteadOfBrackets() (string, error) {
	_template := "WithEndInsteadOfBrackets.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	for i := 0; i < 5; i++ {
		_w(`i=`)
		_w(fmt.Sprintf(`%d`, i))
		_w(`
`)
	}

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`Something here `)
	_w(fmt.Sprintf(`%d`, 5))
	_w(`! And something `)
	_w(fmt.Sprintf(`%s`, _escape("here")))
	_w(`.
And something here: `)
	_w(fmt.Sprintf(`%t`, true))
	_w(`! And `)
	_w(fmt.Sprintf(`%s`, _escape("here, too")))
	_w(`!! Hey, one `)
	_w(fmt.Sprintf(`%s`, _escape("more")))
	_w(`!!!
`)

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`Aaa=`)
	_w(fmt.Sprintf(`%s`, _escape(arg.Aaa)))
	_w(`
Bbb=`)
	_w(fmt.Sprintf(`%d`, arg.Bbb))
	_w(`
`)

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`Will insert something here: `)
	_w(`a=`)
	_w(fmt.Sprintf(`%d`, a))
	_w(`
`)
	_w(`
`)

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`%, str=`)
	_w(fmt.Sprintf(`%s`, _escape(str)))
	_w(`
%, str=`)
	_w(fmt.Sprintf(`%s`, _escape(fmt.Sprintf("aaa%sccc", "bbb"))))
	_w(`
`)

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`String:`)
	_w(fmt.Sprintf(`%s`, _escape(str)))
	_w(`
Unescaped:`)
	_w(fmt.Sprintf(`%s`, str))
	_w(`
Num:`)
	_w(fmt.Sprintf(`%d`, num))
	_w(`
This {{ is ignored
So is this }} !
`)

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`
`)
	_w(`<html>
    <head>
        <title>`)
	_w(fmt.Sprintf(`%s`, _escape(title)))
	_w(`</title>
`)
	_w(`<script>
alert("included")
</script>

`)
	_w(`    </head>
    <body>
`)
	_w(`<h1>Body!</h1>
`)
	_w(`    </body>
</html>
`)

	return _ftmpl.String(), nil
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
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`a is `)
	_w(fmt.Sprintf(`%d`, a))
	_w(`
`)
	return _ftmpl.String(), nil
	_w(`
This line is ignored
`)

	return _ftmpl.String(), nil
}

// TMPLreturn evaluates a template return.tmpl
func TMPLreturn(a int) string {
	html, err := TMPLERRreturn(a)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template return.tmpl:" + err.Error())
	}
	return html
}
