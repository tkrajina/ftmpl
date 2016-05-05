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

// Generated code, do not edit!!!!

func TMPLERRbase(title string) (string, error) {
	__template__ := "base.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//base.tmpl: <html>
	result.WriteString(`<html>
`)
	//base.tmpl:     <head>
	result.WriteString(`    <head>
`)
	//base.tmpl:         <title>{{s title }}</title>
	result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, __escape__(title)))
	//base.tmpl: !#include head
	//base.tmpl:     </head>
	result.WriteString(`    </head>
`)
	//base.tmpl:     <body>
	result.WriteString(`    <body>
`)
	//base.tmpl: !#include body
	//base.tmpl: !#include footer
	//base.tmpl:     </body>
	result.WriteString(`    </body>
`)
	//base.tmpl: </html>
	result.WriteString(`</html>
`)
	//base.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLbase(title string) string {
	html, err := TMPLERRbase(title)
	if err != nil {
		os.Stderr.WriteString("Error running template base.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRbase_embedded(title string) (string, error) {
	__template__ := "base_embedded.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//base_embedded.tmpl:
	result.WriteString(``)
	//base_embedded.tmpl:
	result.WriteString(`
`)
	//base_embedded.tmpl: <html>
	result.WriteString(`<html>
`)
	//base_embedded.tmpl:     <head>
	result.WriteString(`    <head>
`)
	//base_embedded.tmpl:         <title>{{s title }}</title>
	result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, __escape__(title)))
	//base_embedded.tmpl:
	result.WriteString(``)
	//base_embedded.tmpl: !#include head
	//base_embedded.tmpl:
	result.WriteString(`
`)
	//base_embedded.tmpl:     </head>
	result.WriteString(`    </head>
`)
	//base_embedded.tmpl:     <body>
	result.WriteString(`    <body>
`)
	//base_embedded.tmpl:
	result.WriteString(``)
	//base_embedded.tmpl: !#include body
	//base_embedded.tmpl:
	result.WriteString(`
`)
	//base_embedded.tmpl:
	result.WriteString(``)
	//base_embedded.tmpl: !#include footer
	//base_embedded.tmpl:
	result.WriteString(`
`)
	//base_embedded.tmpl:     </body>
	result.WriteString(`    </body>
`)
	//base_embedded.tmpl: </html>
	result.WriteString(`</html>
`)
	//base_embedded.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLbase_embedded(title string) string {
	html, err := TMPLERRbase_embedded(title)
	if err != nil {
		os.Stderr.WriteString("Error running template base_embedded.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRbasic(str string, num int) (string, error) {
	__template__ := "basic.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//basic.tmpl: String:{{s str}}
	result.WriteString(fmt.Sprintf(`String:%s
`, __escape__(str)))
	//basic.tmpl: Unescaped:{{=s str}}
	result.WriteString(fmt.Sprintf(`Unescaped:%s
`, str))
	//basic.tmpl: Num:{{d num}}
	result.WriteString(fmt.Sprintf(`Num:%d
`, num))
	//basic.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLbasic(str string, num int) string {
	html, err := TMPLERRbasic(str, num)
	if err != nil {
		os.Stderr.WriteString("Error running template basic.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRbasic_code(s string, num int) (string, error) {
	__template__ := "basic_code.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//basic_code.tmpl: !for i := 0; i < num; i++ {
	for i := 0; i < num; i++ {
		//basic_code.tmpl: {{d i}}
		result.WriteString(fmt.Sprintf(`%d
`, i))
		//basic_code.tmpl: !}
	}
	//basic_code.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLbasic_code(s string, num int) string {
	html, err := TMPLERRbasic_code(s, num)
	if err != nil {
		os.Stderr.WriteString("Error running template basic_code.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRbasic_embedded_code(n int) (string, error) {
	__template__ := "basic_embedded_code.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//basic_embedded_code.tmpl:
	result.WriteString(``)
	//basic_embedded_code.tmpl: !for i := 0; i < n; i++
	for i := 0; i < n; i++ {
		//basic_embedded_code.tmpl: i={{d i }}
		result.WriteString(fmt.Sprintf(`i=%d `, i))
		//basic_embedded_code.tmpl: !end
	}
	//basic_embedded_code.tmpl:
	result.WriteString(`
`)
	//basic_embedded_code.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLbasic_embedded_code(n int) string {
	html, err := TMPLERRbasic_embedded_code(n)
	if err != nil {
		os.Stderr.WriteString("Error running template basic_embedded_code.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRbasic_if_else(n int) (string, error) {
	__template__ := "basic_if_else.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//basic_if_else.tmpl:
	result.WriteString(``)
	//basic_if_else.tmpl: !if n < 10 {
	if n < 10 {
		//basic_if_else.tmpl: {{d n }} less than 10
		result.WriteString(fmt.Sprintf(`%d less than 10`, n))
		//basic_if_else.tmpl: !}
	}
	//basic_if_else.tmpl:
	result.WriteString(`
`)
	//basic_if_else.tmpl:
	result.WriteString(``)
	//basic_if_else.tmpl: !if n > 0
	if n > 0 {
		//basic_if_else.tmpl: {{d n}} biger than 0
		result.WriteString(fmt.Sprintf(`%d biger than 0`, n))
		//basic_if_else.tmpl: !end
	}
	//basic_if_else.tmpl:
	result.WriteString(`
`)
	//basic_if_else.tmpl:
	result.WriteString(``)
	//basic_if_else.tmpl: !if n > 5
	if n > 5 {
		//basic_if_else.tmpl: {{d n}} biger than 5
		result.WriteString(fmt.Sprintf(`%d biger than 5`, n))
		//basic_if_else.tmpl: !else
	} else {
		//basic_if_else.tmpl: {{d n}} smaller than 5
		result.WriteString(fmt.Sprintf(`%d smaller than 5`, n))
		//basic_if_else.tmpl: !end
	}
	//basic_if_else.tmpl:
	result.WriteString(`
`)
	//basic_if_else.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLbasic_if_else(n int) string {
	html, err := TMPLERRbasic_if_else(n)
	if err != nil {
		os.Stderr.WriteString("Error running template basic_if_else.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRbasic_if_elseif(n int) (string, error) {
	__template__ := "basic_if_elseif.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//basic_if_elseif.tmpl:
	result.WriteString(``)
	//basic_if_elseif.tmpl:
	result.WriteString(`
`)
	//basic_if_elseif.tmpl:
	result.WriteString(``)
	//basic_if_elseif.tmpl: !if n < 10
	if n < 10 {
		//basic_if_elseif.tmpl: n less than 10
		result.WriteString(`n less than 10`)
		//basic_if_elseif.tmpl: !else if n < 100
	} else if n < 100 {
		//basic_if_elseif.tmpl: n less than 100
		result.WriteString(`n less than 100`)
		//basic_if_elseif.tmpl: !else
	} else {
		//basic_if_elseif.tmpl: n bigger than 100
		result.WriteString(`n bigger than 100`)
		//basic_if_elseif.tmpl: !end
	}
	//basic_if_elseif.tmpl:
	result.WriteString(`
`)
	//basic_if_elseif.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLbasic_if_elseif(n int) string {
	html, err := TMPLERRbasic_if_elseif(n)
	if err != nil {
		os.Stderr.WriteString("Error running template basic_if_elseif.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRcomparison_with_gotemplates(params TemplateParam) (string, error) {
	__template__ := "comparison_with_gotemplates.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//comparison_with_gotemplates.tmpl: <html>
	result.WriteString(`<html>
`)
	//comparison_with_gotemplates.tmpl:     <head>
	result.WriteString(`    <head>
`)
	//comparison_with_gotemplates.tmpl:         <title>{{s params.Title }}</title>
	result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, __escape__(params.Title)))
	//comparison_with_gotemplates.tmpl:     </head>
	result.WriteString(`    </head>
`)
	//comparison_with_gotemplates.tmpl:     <body>
	result.WriteString(`    <body>
`)
	//comparison_with_gotemplates.tmpl:         <h1>{{s params.Title }}</h1>
	result.WriteString(fmt.Sprintf(`        <h1>%s</h1>
`, __escape__(params.Title)))
	//comparison_with_gotemplates.tmpl:
	result.WriteString(`        `)
	//comparison_with_gotemplates.tmpl: !if len(params.Subtitle) > 0
	if len(params.Subtitle) > 0 {
		//comparison_with_gotemplates.tmpl: <h2>{{ params.Subtitle }}</h1>
		result.WriteString(fmt.Sprintf(`<h2>%v</h1>`, params.Subtitle))
		//comparison_with_gotemplates.tmpl: !end
	}
	//comparison_with_gotemplates.tmpl:
	result.WriteString(`
`)
	//comparison_with_gotemplates.tmpl:         <ul>
	result.WriteString(`        <ul>
`)
	//comparison_with_gotemplates.tmpl:
	result.WriteString(`            `)
	//comparison_with_gotemplates.tmpl: !for _, item := range params.Items
	for _, item := range params.Items {
		//comparison_with_gotemplates.tmpl:
		result.WriteString(`
`)
		//comparison_with_gotemplates.tmpl:                 <li> {{s item }}
		result.WriteString(fmt.Sprintf(`                <li> %s
`, __escape__(item)))
		//comparison_with_gotemplates.tmpl:
		result.WriteString(`            `)
		//comparison_with_gotemplates.tmpl: !end
	}
	//comparison_with_gotemplates.tmpl:
	result.WriteString(`
`)
	//comparison_with_gotemplates.tmpl:         </ul>
	result.WriteString(`        </ul>
`)
	//comparison_with_gotemplates.tmpl:         <p>
	result.WriteString(`        <p>
`)
	//comparison_with_gotemplates.tmpl:             Written {{d len(params.Items) }} items
	result.WriteString(fmt.Sprintf(`            Written %d items
`, len(params.Items)))
	//comparison_with_gotemplates.tmpl:         </p>
	result.WriteString(`        </p>
`)
	//comparison_with_gotemplates.tmpl:     </body>
	result.WriteString(`    </body>
`)
	//comparison_with_gotemplates.tmpl: </html>
	result.WriteString(`</html>
`)
	//comparison_with_gotemplates.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLcomparison_with_gotemplates(params TemplateParam) string {
	html, err := TMPLERRcomparison_with_gotemplates(params)
	if err != nil {
		os.Stderr.WriteString("Error running template comparison_with_gotemplates.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRextends(title string, something int) (string, error) {
	__template__ := "extends.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//extends.tmpl: !#extends base
	//extends.tmpl:
	result.WriteString(`
`)
	//extends.tmpl: <html>
	result.WriteString(`<html>
`)
	//extends.tmpl:     <head>
	result.WriteString(`    <head>
`)
	//extends.tmpl:         <title>{{s title }}</title>
	result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, __escape__(title)))
	//extends.tmpl: <script>
	result.WriteString(`<script>
`)
	//extends.tmpl: alert("included")
	result.WriteString(`alert("included")
`)
	//extends.tmpl: </script>
	result.WriteString(`</script>
`)
	//extends.tmpl:
	result.WriteString(`
`)
	//extends.tmpl:     </head>
	result.WriteString(`    </head>
`)
	//extends.tmpl:     <body>
	result.WriteString(`    <body>
`)
	//extends.tmpl: <h1>Body!</h1>
	result.WriteString(`<h1>Body!</h1>
`)
	//extends.tmpl:
	result.WriteString(``)
	//extends.tmpl:     </body>
	result.WriteString(`    </body>
`)
	//extends.tmpl: </html>
	result.WriteString(`</html>
`)
	//extends.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLextends(title string, something int) string {
	html, err := TMPLERRextends(title, something)
	if err != nil {
		os.Stderr.WriteString("Error running template extends.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRextends_embedded(title string, something int) (string, error) {
	__template__ := "extends_embedded.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//extends_embedded.tmpl:
	result.WriteString(``)
	//extends_embedded.tmpl: !#extends base_embedded
	//extends_embedded.tmpl:
	result.WriteString(`
`)
	//extends_embedded.tmpl:
	result.WriteString(``)
	//extends_embedded.tmpl:
	result.WriteString(`
`)
	//extends_embedded.tmpl:
	result.WriteString(`
`)
	//extends_embedded.tmpl:
	result.WriteString(``)
	//extends_embedded.tmpl:
	result.WriteString(``)
	//extends_embedded.tmpl:
	result.WriteString(`
`)
	//extends_embedded.tmpl: <html>
	result.WriteString(`<html>
`)
	//extends_embedded.tmpl:     <head>
	result.WriteString(`    <head>
`)
	//extends_embedded.tmpl:         <title>{{s title }}</title>
	result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, __escape__(title)))
	//extends_embedded.tmpl:
	result.WriteString(``)
	//extends_embedded.tmpl:
	result.WriteString(`
`)
	//extends_embedded.tmpl: <script>
	result.WriteString(`<script>
`)
	//extends_embedded.tmpl: alert("included")
	result.WriteString(`alert("included")
`)
	//extends_embedded.tmpl: </script>
	result.WriteString(`</script>
`)
	//extends_embedded.tmpl:
	result.WriteString(`
`)
	//extends_embedded.tmpl:
	result.WriteString(``)
	//extends_embedded.tmpl:
	result.WriteString(`
`)
	//extends_embedded.tmpl:     </head>
	result.WriteString(`    </head>
`)
	//extends_embedded.tmpl:     <body>
	result.WriteString(`    <body>
`)
	//extends_embedded.tmpl:
	result.WriteString(``)
	//extends_embedded.tmpl:
	result.WriteString(`
`)
	//extends_embedded.tmpl: <h1>Body!</h1>
	result.WriteString(`<h1>Body!</h1>
`)
	//extends_embedded.tmpl:
	result.WriteString(``)
	//extends_embedded.tmpl:
	result.WriteString(`
`)
	//extends_embedded.tmpl:
	result.WriteString(``)
	//extends_embedded.tmpl:
	result.WriteString(`
`)
	//extends_embedded.tmpl:     </body>
	result.WriteString(`    </body>
`)
	//extends_embedded.tmpl: </html>
	result.WriteString(`</html>
`)
	//extends_embedded.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLextends_embedded(title string, something int) string {
	html, err := TMPLERRextends_embedded(title, something)
	if err != nil {
		os.Stderr.WriteString("Error running template extends_embedded.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRnoncode_line_with_exclamation_mark() (string, error) {
	__template__ := "noncode_line_with_exclamation_mark.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//noncode_line_with_exclamation_mark.tmpl: !s1 := "This lins is not a code line"
	result.WriteString(`!s1 := "This lins is not a code line"
`)
	//noncode_line_with_exclamation_mark.tmpl: !s2 := "This *is* a line of code"
	s2 := "This *is* a line of code"
	//noncode_line_with_exclamation_mark.tmpl: {{s s2 }}
	result.WriteString(fmt.Sprintf(`%s
`, __escape__(s2)))
	//noncode_line_with_exclamation_mark.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLnoncode_line_with_exclamation_mark() string {
	html, err := TMPLERRnoncode_line_with_exclamation_mark()
	if err != nil {
		os.Stderr.WriteString("Error running template noncode_line_with_exclamation_mark.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRreturn(a int) (string, error) {
	__template__ := "return.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//return.tmpl: a is {{d a }}
	result.WriteString(fmt.Sprintf(`a is %d
`, a))
	//return.tmpl:
	result.WriteString(``)
	return result.String(), nil
	//return.tmpl:
	result.WriteString(`
`)
	//return.tmpl: This line is ignored
	result.WriteString(`This line is ignored
`)
	//return.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLreturn(a int) string {
	html, err := TMPLERRreturn(a)
	if err != nil {
		os.Stderr.WriteString("Error running template return.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRwith_end_instead_of_brackets() (string, error) {
	__template__ := "with_end_instead_of_brackets.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//with_end_instead_of_brackets.tmpl: !for i:=0; i < 5; i++
	for i := 0; i < 5; i++ {
		//with_end_instead_of_brackets.tmpl: i={{d i }}
		result.WriteString(fmt.Sprintf(`i=%d
`, i))
		//with_end_instead_of_brackets.tmpl: !end
	}
	//with_end_instead_of_brackets.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLwith_end_instead_of_brackets() string {
	html, err := TMPLERRwith_end_instead_of_brackets()
	if err != nil {
		os.Stderr.WriteString("Error running template with_end_instead_of_brackets.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
type Argument struct {
	Aaa string
	Bbb int
}

func TMPLERRwith_global_declaration(arg Argument) (string, error) {
	__template__ := "with_global_declaration.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//with_global_declaration.tmpl: Aaa={{s arg.Aaa }}
	result.WriteString(fmt.Sprintf(`Aaa=%s
`, __escape__(arg.Aaa)))
	//with_global_declaration.tmpl: Bbb={{d arg.Bbb }}
	result.WriteString(fmt.Sprintf(`Bbb=%d
`, arg.Bbb))
	//with_global_declaration.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLwith_global_declaration(arg Argument) string {
	html, err := TMPLERRwith_global_declaration(arg)
	if err != nil {
		os.Stderr.WriteString("Error running template with_global_declaration.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!

func TMPLERRwith_percentage(str string) (string, error) {
	__template__ := "with_percentage.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	//with_percentage.tmpl: %, str={{s str }}
	result.WriteString(fmt.Sprintf(`%%, str=%s
`, __escape__(str)))
	//with_percentage.tmpl: %, str={{s fmt.Sprintf("aaa%sccc", "bbb") }}
	result.WriteString(fmt.Sprintf(`%%, str=%s
`, __escape__(fmt.Sprintf("aaa%sccc", "bbb"))))
	//with_percentage.tmpl:
	result.WriteString(``)

	return result.String(), nil
}

func TMPLwith_percentage(str string) string {
	html, err := TMPLERRwith_percentage(str)
	if err != nil {
		os.Stderr.WriteString("Error running template with_percentage.tmpl:" + err.Error())
	}
	return html
}
