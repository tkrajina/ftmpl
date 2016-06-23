// package example is generated, do not edit!!!! */
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
	//BaseEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//BaseEmbedded.tmpl: <html>
	_, _ = result.WriteString(`<html>
`)
	//BaseEmbedded.tmpl:     <head>
	_, _ = result.WriteString(`    <head>
`)
	//BaseEmbedded.tmpl:         <title>{{s title }}</title>
	_, _ = result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, _escape(title)))
	//BaseEmbedded.tmpl: !#include head
	//BaseEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//BaseEmbedded.tmpl:     </head>
	_, _ = result.WriteString(`    </head>
`)
	//BaseEmbedded.tmpl:     <body>
	_, _ = result.WriteString(`    <body>
`)
	//BaseEmbedded.tmpl: !#include body
	//BaseEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//BaseEmbedded.tmpl: !#include footer
	//BaseEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//BaseEmbedded.tmpl:     </body>
	_, _ = result.WriteString(`    </body>
`)
	//BaseEmbedded.tmpl: </html>
	_, _ = result.WriteString(`</html>
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
	//BasicCode.tmpl: !for i := 0; i < num; i++ {
	for i := 0; i < num; i++ {
		//BasicCode.tmpl: {{d i}}
		_, _ = result.WriteString(fmt.Sprintf(`%d
`, i))
		//BasicCode.tmpl: !}
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
	//BasicEmbeddedCode.tmpl: !for i := 0; i < n; i++
	for i := 0; i < n; i++ {
		//BasicEmbeddedCode.tmpl: i={{d i }}
		_, _ = result.WriteString(fmt.Sprintf(`i=%d `, i))
		//BasicEmbeddedCode.tmpl: !end
	}
	//BasicEmbeddedCode.tmpl:
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
	//BasicIfElse.tmpl: !if n < 10 {
	if n < 10 {
		//BasicIfElse.tmpl: {{d n }} less than 10
		_, _ = result.WriteString(fmt.Sprintf(`%d less than 10`, n))
		//BasicIfElse.tmpl: !}
	}
	//BasicIfElse.tmpl:
	_, _ = result.WriteString(`
`)
	//BasicIfElse.tmpl: !if n > 0
	if n > 0 {
		//BasicIfElse.tmpl: {{d n}} biger than 0
		_, _ = result.WriteString(fmt.Sprintf(`%d biger than 0`, n))
		//BasicIfElse.tmpl: !end
	}
	//BasicIfElse.tmpl:
	_, _ = result.WriteString(`
`)
	//BasicIfElse.tmpl: !if n > 5
	if n > 5 {
		//BasicIfElse.tmpl: {{d n}} biger than 5
		_, _ = result.WriteString(fmt.Sprintf(`%d biger than 5`, n))
		//BasicIfElse.tmpl: !else
	} else {
		//BasicIfElse.tmpl: {{d n}} smaller than 5
		_, _ = result.WriteString(fmt.Sprintf(`%d smaller than 5`, n))
		//BasicIfElse.tmpl: !end
	}
	//BasicIfElse.tmpl:
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
	//BasicIfElseif.tmpl:
	_, _ = result.WriteString(`
`)
	//BasicIfElseif.tmpl: !if n < 10
	if n < 10 {
		//BasicIfElseif.tmpl: n less than 10
		_, _ = result.WriteString(`n less than 10`)
		//BasicIfElseif.tmpl: !else if n < 100
	} else if n < 100 {
		//BasicIfElseif.tmpl: n less than 100
		_, _ = result.WriteString(`n less than 100`)
		//BasicIfElseif.tmpl: !else
	} else {
		//BasicIfElseif.tmpl: n bigger than 100
		_, _ = result.WriteString(`n bigger than 100`)
		//BasicIfElseif.tmpl: !end
	}
	//BasicIfElseif.tmpl:
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
	//ComparisonWithGoTemplates.tmpl: <html>
	_, _ = result.WriteString(`<html>
`)
	//ComparisonWithGoTemplates.tmpl:     <head>
	_, _ = result.WriteString(`    <head>
`)
	//ComparisonWithGoTemplates.tmpl:         <title>{{s params.Title }}</title>
	_, _ = result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, _escape(params.Title)))
	//ComparisonWithGoTemplates.tmpl:     </head>
	_, _ = result.WriteString(`    </head>
`)
	//ComparisonWithGoTemplates.tmpl:     <body>
	_, _ = result.WriteString(`    <body>
`)
	//ComparisonWithGoTemplates.tmpl:         <h1>{{s params.Title }}</h1>
	_, _ = result.WriteString(fmt.Sprintf(`        <h1>%s</h1>
`, _escape(params.Title)))
	//ComparisonWithGoTemplates.tmpl:
	_, _ = result.WriteString(`        `)
	//ComparisonWithGoTemplates.tmpl: !if len(params.Subtitle) > 0
	if len(params.Subtitle) > 0 {
		//ComparisonWithGoTemplates.tmpl: <h2>{{ params.Subtitle }}</h1>
		_, _ = result.WriteString(fmt.Sprintf(`<h2>%v</h1>`, params.Subtitle))
		//ComparisonWithGoTemplates.tmpl: !end
	}
	//ComparisonWithGoTemplates.tmpl:
	_, _ = result.WriteString(`
`)
	//ComparisonWithGoTemplates.tmpl:         <ul>
	_, _ = result.WriteString(`        <ul>
`)
	//ComparisonWithGoTemplates.tmpl:
	_, _ = result.WriteString(`            `)
	//ComparisonWithGoTemplates.tmpl: !for _, item := range params.Items
	for _, item := range params.Items {
		//ComparisonWithGoTemplates.tmpl:
		_, _ = result.WriteString(`
`)
		//ComparisonWithGoTemplates.tmpl:                 <li> {{s item }}
		_, _ = result.WriteString(fmt.Sprintf(`                <li> %s
`, _escape(item)))
		//ComparisonWithGoTemplates.tmpl:
		_, _ = result.WriteString(`            `)
		//ComparisonWithGoTemplates.tmpl: !end
	}
	//ComparisonWithGoTemplates.tmpl:
	_, _ = result.WriteString(`
`)
	//ComparisonWithGoTemplates.tmpl:         </ul>
	_, _ = result.WriteString(`        </ul>
`)
	//ComparisonWithGoTemplates.tmpl:         <p>
	_, _ = result.WriteString(`        <p>
`)
	//ComparisonWithGoTemplates.tmpl:             Written {{d len(params.Items) }} items
	_, _ = result.WriteString(fmt.Sprintf(`            Written %d items
`, len(params.Items)))
	//ComparisonWithGoTemplates.tmpl:         </p>
	_, _ = result.WriteString(`        </p>
`)
	//ComparisonWithGoTemplates.tmpl:     </body>
	_, _ = result.WriteString(`    </body>
`)
	//ComparisonWithGoTemplates.tmpl: </html>
	_, _ = result.WriteString(`</html>
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
	//ExtendsEmbedded.tmpl: !#extends BaseEmbedded
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//ExtendsEmbedded.tmpl: <html>
	_, _ = result.WriteString(`<html>
`)
	//ExtendsEmbedded.tmpl:     <head>
	_, _ = result.WriteString(`    <head>
`)
	//ExtendsEmbedded.tmpl:         <title>{{s title }}</title>
	_, _ = result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, _escape(title)))
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//ExtendsEmbedded.tmpl: <script>
	_, _ = result.WriteString(`<script>
`)
	//ExtendsEmbedded.tmpl: alert("included")
	_, _ = result.WriteString(`alert("included")
`)
	//ExtendsEmbedded.tmpl: </script>
	_, _ = result.WriteString(`</script>
`)
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//ExtendsEmbedded.tmpl:     </head>
	_, _ = result.WriteString(`    </head>
`)
	//ExtendsEmbedded.tmpl:     <body>
	_, _ = result.WriteString(`    <body>
`)
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//ExtendsEmbedded.tmpl: <h1>Body!</h1>
	_, _ = result.WriteString(`<h1>Body!</h1>
`)
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//ExtendsEmbedded.tmpl:     </body>
	_, _ = result.WriteString(`    </body>
`)
	//ExtendsEmbedded.tmpl: </html>
	_, _ = result.WriteString(`</html>
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
	//NoncodeLineWithExclamationMark.tmpl: !s1 := "This lins is not a code line"
	_, _ = result.WriteString(`!s1 := "This lins is not a code line"
`)
	//NoncodeLineWithExclamationMark.tmpl: !s2 := "This *is* a line of code"
	s2 := "This *is* a line of code"
	//NoncodeLineWithExclamationMark.tmpl: {{s s2 }}
	_, _ = result.WriteString(fmt.Sprintf(`%s
`, _escape(s2)))

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
	//WithEndInsteadOfBrackets.tmpl: !for i:=0; i < 5; i++
	for i := 0; i < 5; i++ {
		//WithEndInsteadOfBrackets.tmpl: i={{d i }}
		_, _ = result.WriteString(fmt.Sprintf(`i=%d
`, i))
		//WithEndInsteadOfBrackets.tmpl: !end
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
	//WithGlobalDeclaration.tmpl: Aaa={{s arg.Aaa }}
	_, _ = result.WriteString(fmt.Sprintf(`Aaa=%s
`, _escape(arg.Aaa)))
	//WithGlobalDeclaration.tmpl: Bbb={{d arg.Bbb }}
	_, _ = result.WriteString(fmt.Sprintf(`Bbb=%d
`, arg.Bbb))

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
	//WithInsert.tmpl: Will insert something here:
	_, _ = result.WriteString(`Will insert something here: `)
	//WithInsert.tmpl: a={{d a }}
	_, _ = result.WriteString(fmt.Sprintf(`a=%d
`, a))
	//WithInsert.tmpl:
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
	//WithPercentage.tmpl: %, str={{s str }}
	_, _ = result.WriteString(fmt.Sprintf(`%%, str=%s
`, _escape(str)))
	//WithPercentage.tmpl: %, str={{s fmt.Sprintf("aaa%sccc", "bbb") }}
	_, _ = result.WriteString(fmt.Sprintf(`%%, str=%s
`, _escape(fmt.Sprintf("aaa%sccc", "bbb"))))

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
	//basic.tmpl: String:{{s str}}
	_, _ = result.WriteString(fmt.Sprintf(`String:%s
`, _escape(str)))
	//basic.tmpl: Unescaped:{{=s str}}
	_, _ = result.WriteString(fmt.Sprintf(`Unescaped:%s
`, str))
	//basic.tmpl: Num:{{d num}}
	_, _ = result.WriteString(fmt.Sprintf(`Num:%d
`, num))

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
	//extends.tmpl: !#extends base
	//extends.tmpl:
	_, _ = result.WriteString(`
`)
	//extends.tmpl: <html>
	_, _ = result.WriteString(`<html>
`)
	//extends.tmpl:     <head>
	_, _ = result.WriteString(`    <head>
`)
	//extends.tmpl:         <title>{{s title }}</title>
	_, _ = result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, _escape(title)))
	//extends.tmpl: <script>
	_, _ = result.WriteString(`<script>
`)
	//extends.tmpl: alert("included")
	_, _ = result.WriteString(`alert("included")
`)
	//extends.tmpl: </script>
	_, _ = result.WriteString(`</script>
`)
	//extends.tmpl:
	_, _ = result.WriteString(`
`)
	//extends.tmpl:     </head>
	_, _ = result.WriteString(`    </head>
`)
	//extends.tmpl:     <body>
	_, _ = result.WriteString(`    <body>
`)
	//extends.tmpl: <h1>Body!</h1>
	_, _ = result.WriteString(`<h1>Body!</h1>
`)
	//extends.tmpl:     </body>
	_, _ = result.WriteString(`    </body>
`)
	//extends.tmpl: </html>
	_, _ = result.WriteString(`</html>
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
	//return.tmpl: a is {{d a }}
	_, _ = result.WriteString(fmt.Sprintf(`a is %d
`, a))
	return result.String(), nil
	//return.tmpl:
	_, _ = result.WriteString(`
`)
	//return.tmpl: This line is ignored
	_, _ = result.WriteString(`This line is ignored
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
