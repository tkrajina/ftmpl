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
	//BaseEmbedded.tmpl:
	//BaseEmbedded.tmpl: <html>
	//BaseEmbedded.tmpl:     <head>
	//BaseEmbedded.tmpl:         <title>{{s title }}</title>
	//BaseEmbedded.tmpl:
	_, _ = result.WriteString(fmt.Sprintf(`
<html>
    <head>
        <title>%s</title>
`, _escape(title)))
	//BaseEmbedded.tmpl: !#include head
	//BaseEmbedded.tmpl:
	//BaseEmbedded.tmpl:     </head>
	//BaseEmbedded.tmpl:     <body>
	//BaseEmbedded.tmpl:
	_, _ = result.WriteString(`
    </head>
    <body>
`)
	//BaseEmbedded.tmpl: !#include body
	//BaseEmbedded.tmpl:
	//BaseEmbedded.tmpl:
	_, _ = result.WriteString(`
`)
	//BaseEmbedded.tmpl: !#include footer
	//BaseEmbedded.tmpl:
	//BaseEmbedded.tmpl:     </body>
	//BaseEmbedded.tmpl: </html>
	//BaseEmbedded.tmpl:
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
	//BasicCode.tmpl: !for i := 0; i < num; i++ {
	for i := 0; i < num; i++ {
		//BasicCode.tmpl: {{d i}}
		_, _ = result.WriteString(fmt.Sprintf(`%d
`, i))
		//BasicCode.tmpl: !}
	}
	//BasicCode.tmpl:

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
	//BasicEmbeddedCode.tmpl:
	//BasicEmbeddedCode.tmpl: !for i := 0; i < n; i++
	for i := 0; i < n; i++ {
		//BasicEmbeddedCode.tmpl: i={{d i }}
		_, _ = result.WriteString(fmt.Sprintf(`i=%d `, i))
		//BasicEmbeddedCode.tmpl: !end
	}
	//BasicEmbeddedCode.tmpl:
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
	//BasicIfElse.tmpl:
	//BasicIfElse.tmpl: !if n < 10 {
	if n < 10 {
		//BasicIfElse.tmpl: {{d n }} less than 10
		_, _ = result.WriteString(fmt.Sprintf(`%d less than 10`, n))
		//BasicIfElse.tmpl: !}
	}
	//BasicIfElse.tmpl:
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
	//BasicIfElseif.tmpl:
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
	//ComparisonWithGoTemplates.tmpl:     <head>
	//ComparisonWithGoTemplates.tmpl:         <title>{{s params.Title }}</title>
	//ComparisonWithGoTemplates.tmpl:     </head>
	//ComparisonWithGoTemplates.tmpl:     <body>
	//ComparisonWithGoTemplates.tmpl:         <h1>{{s params.Title }}</h1>
	//ComparisonWithGoTemplates.tmpl:
	_, _ = result.WriteString(fmt.Sprintf(`<html>
    <head>
        <title>%s</title>
    </head>
    <body>
        <h1>%s</h1>
        `, _escape(params.Title), _escape(params.Title)))
	//ComparisonWithGoTemplates.tmpl: !if len(params.Subtitle) > 0
	if len(params.Subtitle) > 0 {
		//ComparisonWithGoTemplates.tmpl: <h2>{{ params.Subtitle }}</h1>
		_, _ = result.WriteString(fmt.Sprintf(`<h2>%v</h1>`, params.Subtitle))
		//ComparisonWithGoTemplates.tmpl: !end
	}
	//ComparisonWithGoTemplates.tmpl:
	//ComparisonWithGoTemplates.tmpl:         <ul>
	//ComparisonWithGoTemplates.tmpl:
	_, _ = result.WriteString(`
        <ul>
            `)
	//ComparisonWithGoTemplates.tmpl: !for _, item := range params.Items
	for _, item := range params.Items {
		//ComparisonWithGoTemplates.tmpl:
		//ComparisonWithGoTemplates.tmpl:                 <li> {{s item }}
		//ComparisonWithGoTemplates.tmpl:
		_, _ = result.WriteString(fmt.Sprintf(`
                <li> %s
            `, _escape(item)))
		//ComparisonWithGoTemplates.tmpl: !end
	}
	//ComparisonWithGoTemplates.tmpl:
	//ComparisonWithGoTemplates.tmpl:         </ul>
	//ComparisonWithGoTemplates.tmpl:         <p>
	//ComparisonWithGoTemplates.tmpl:             Written {{d len(params.Items) }} items
	//ComparisonWithGoTemplates.tmpl:         </p>
	//ComparisonWithGoTemplates.tmpl:     </body>
	//ComparisonWithGoTemplates.tmpl: </html>
	//ComparisonWithGoTemplates.tmpl:
	_, _ = result.WriteString(fmt.Sprintf(`
        </ul>
        <p>
            Written %d items
        </p>
    </body>
</html>
`, len(params.Items)))

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
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl: !#extends BaseEmbedded
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(`


`)
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl: <html>
	//ExtendsEmbedded.tmpl:     <head>
	//ExtendsEmbedded.tmpl:         <title>{{s title }}</title>
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl: <script>
	//ExtendsEmbedded.tmpl: alert("included")
	//ExtendsEmbedded.tmpl: </script>
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:     </head>
	//ExtendsEmbedded.tmpl:     <body>
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl: <h1>Body!</h1>
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:
	//ExtendsEmbedded.tmpl:     </body>
	//ExtendsEmbedded.tmpl: </html>
	//ExtendsEmbedded.tmpl:
	_, _ = result.WriteString(fmt.Sprintf(`
<html>
    <head>
        <title>%s</title>

<script>
alert("included")
</script>


    </head>
    <body>

<h1>Body!</h1>


    </body>
</html>
`, _escape(title)))

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
	//NoncodeLineWithExclamationMark.tmpl:
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
	//WithEndInsteadOfBrackets.tmpl:

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
	//WithGlobalDeclaration.tmpl: Bbb={{d arg.Bbb }}
	//WithGlobalDeclaration.tmpl:
	_, _ = result.WriteString(fmt.Sprintf(`Aaa=%s
Bbb=%d
`, _escape(arg.Aaa), arg.Bbb))

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
	//WithInsert.tmpl: a={{d a }}
	//WithInsert.tmpl:
	//WithInsert.tmpl:
	//WithInsert.tmpl:
	_, _ = result.WriteString(fmt.Sprintf(`Will insert something here: a=%d

`, a))

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
	//WithPercentage.tmpl: %, str={{s fmt.Sprintf("aaa%sccc", "bbb") }}
	//WithPercentage.tmpl:
	_, _ = result.WriteString(fmt.Sprintf(`%%, str=%s
%%, str=%s
`, _escape(str), _escape(fmt.Sprintf("aaa%sccc", "bbb"))))

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
	//basic.tmpl: Unescaped:{{=s str}}
	//basic.tmpl: Num:{{d num}}
	//basic.tmpl:
	_, _ = result.WriteString(fmt.Sprintf(`String:%s
Unescaped:%s
Num:%d
`, _escape(str), str, num))

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
	//extends.tmpl:     <head>
	//extends.tmpl:         <title>{{s title }}</title>
	//extends.tmpl: <script>
	//extends.tmpl: alert("included")
	//extends.tmpl: </script>
	//extends.tmpl:
	//extends.tmpl:     </head>
	//extends.tmpl:     <body>
	//extends.tmpl: <h1>Body!</h1>
	//extends.tmpl:
	//extends.tmpl:     </body>
	//extends.tmpl: </html>
	//extends.tmpl:
	_, _ = result.WriteString(fmt.Sprintf(`<html>
    <head>
        <title>%s</title>
<script>
alert("included")
</script>

    </head>
    <body>
<h1>Body!</h1>
    </body>
</html>
`, _escape(title)))

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
	//return.tmpl:
	_, _ = result.WriteString(fmt.Sprintf(`a is %d
`, a))
	return result.String(), nil
	//return.tmpl:
	//return.tmpl: This line is ignored
	//return.tmpl:
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
