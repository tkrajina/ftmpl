/* Generated code, do not edit!!!! */
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

// TMPLERRbase evaluates a template .TemplateFile
func TMPLERRbase(title string) (string, error) {
	_template := "base.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//base.tmpl: <html>
	_, _ = result.WriteString(`<html>
`)
	//base.tmpl:     <head>
	_, _ = result.WriteString(`    <head>
`)
	//base.tmpl:         <title>{{s title }}</title>
	_, _ = result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, _escape(title)))
	//base.tmpl: !#include head
	//base.tmpl:     </head>
	_, _ = result.WriteString(`    </head>
`)
	//base.tmpl:     <body>
	_, _ = result.WriteString(`    <body>
`)
	//base.tmpl: !#include body
	//base.tmpl: !#include footer
	//base.tmpl:     </body>
	_, _ = result.WriteString(`    </body>
`)
	//base.tmpl: </html>
	_, _ = result.WriteString(`</html>
`)
	//base.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLbase evaluates a template .TemplateFile
func TMPLbase(title string) string {
	html, err := TMPLERRbase(title)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template base.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRbase_embedded evaluates a template .TemplateFile
func TMPLERRbase_embedded(title string) (string, error) {
	_template := "base_embedded.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//base_embedded.tmpl:
	_, _ = result.WriteString(``)
	//base_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//base_embedded.tmpl: <html>
	_, _ = result.WriteString(`<html>
`)
	//base_embedded.tmpl:     <head>
	_, _ = result.WriteString(`    <head>
`)
	//base_embedded.tmpl:         <title>{{s title }}</title>
	_, _ = result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, _escape(title)))
	//base_embedded.tmpl:
	_, _ = result.WriteString(``)
	//base_embedded.tmpl: !#include head
	//base_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//base_embedded.tmpl:     </head>
	_, _ = result.WriteString(`    </head>
`)
	//base_embedded.tmpl:     <body>
	_, _ = result.WriteString(`    <body>
`)
	//base_embedded.tmpl:
	_, _ = result.WriteString(``)
	//base_embedded.tmpl: !#include body
	//base_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//base_embedded.tmpl:
	_, _ = result.WriteString(``)
	//base_embedded.tmpl: !#include footer
	//base_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//base_embedded.tmpl:     </body>
	_, _ = result.WriteString(`    </body>
`)
	//base_embedded.tmpl: </html>
	_, _ = result.WriteString(`</html>
`)
	//base_embedded.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLbase_embedded evaluates a template .TemplateFile
func TMPLbase_embedded(title string) string {
	html, err := TMPLERRbase_embedded(title)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template base_embedded.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRbasic evaluates a template .TemplateFile
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
	//basic.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLbasic evaluates a template .TemplateFile
func TMPLbasic(str string, num int) string {
	html, err := TMPLERRbasic(str, num)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template basic.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRbasic_code evaluates a template .TemplateFile
func TMPLERRbasic_code(s string, num int) (string, error) {
	_template := "basic_code.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//basic_code.tmpl: !for i := 0; i < num; i++ {
	for i := 0; i < num; i++ {
		//basic_code.tmpl: {{d i}}
		_, _ = result.WriteString(fmt.Sprintf(`%d
`, i))
		//basic_code.tmpl: !}
	}
	//basic_code.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLbasic_code evaluates a template .TemplateFile
func TMPLbasic_code(s string, num int) string {
	html, err := TMPLERRbasic_code(s, num)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template basic_code.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRbasic_embedded_code evaluates a template .TemplateFile
func TMPLERRbasic_embedded_code(n int) (string, error) {
	_template := "basic_embedded_code.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//basic_embedded_code.tmpl:
	_, _ = result.WriteString(``)
	//basic_embedded_code.tmpl: !for i := 0; i < n; i++
	for i := 0; i < n; i++ {
		//basic_embedded_code.tmpl: i={{d i }}
		_, _ = result.WriteString(fmt.Sprintf(`i=%d `, i))
		//basic_embedded_code.tmpl: !end
	}
	//basic_embedded_code.tmpl:
	_, _ = result.WriteString(`
`)
	//basic_embedded_code.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLbasic_embedded_code evaluates a template .TemplateFile
func TMPLbasic_embedded_code(n int) string {
	html, err := TMPLERRbasic_embedded_code(n)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template basic_embedded_code.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRbasic_if_else evaluates a template .TemplateFile
func TMPLERRbasic_if_else(n int) (string, error) {
	_template := "basic_if_else.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//basic_if_else.tmpl:
	_, _ = result.WriteString(``)
	//basic_if_else.tmpl: !if n < 10 {
	if n < 10 {
		//basic_if_else.tmpl: {{d n }} less than 10
		_, _ = result.WriteString(fmt.Sprintf(`%d less than 10`, n))
		//basic_if_else.tmpl: !}
	}
	//basic_if_else.tmpl:
	_, _ = result.WriteString(`
`)
	//basic_if_else.tmpl:
	_, _ = result.WriteString(``)
	//basic_if_else.tmpl: !if n > 0
	if n > 0 {
		//basic_if_else.tmpl: {{d n}} biger than 0
		_, _ = result.WriteString(fmt.Sprintf(`%d biger than 0`, n))
		//basic_if_else.tmpl: !end
	}
	//basic_if_else.tmpl:
	_, _ = result.WriteString(`
`)
	//basic_if_else.tmpl:
	_, _ = result.WriteString(``)
	//basic_if_else.tmpl: !if n > 5
	if n > 5 {
		//basic_if_else.tmpl: {{d n}} biger than 5
		_, _ = result.WriteString(fmt.Sprintf(`%d biger than 5`, n))
		//basic_if_else.tmpl: !else
	} else {
		//basic_if_else.tmpl: {{d n}} smaller than 5
		_, _ = result.WriteString(fmt.Sprintf(`%d smaller than 5`, n))
		//basic_if_else.tmpl: !end
	}
	//basic_if_else.tmpl:
	_, _ = result.WriteString(`
`)
	//basic_if_else.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLbasic_if_else evaluates a template .TemplateFile
func TMPLbasic_if_else(n int) string {
	html, err := TMPLERRbasic_if_else(n)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template basic_if_else.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRbasic_if_elseif evaluates a template .TemplateFile
func TMPLERRbasic_if_elseif(n int) (string, error) {
	_template := "basic_if_elseif.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//basic_if_elseif.tmpl:
	_, _ = result.WriteString(``)
	//basic_if_elseif.tmpl:
	_, _ = result.WriteString(`
`)
	//basic_if_elseif.tmpl:
	_, _ = result.WriteString(``)
	//basic_if_elseif.tmpl: !if n < 10
	if n < 10 {
		//basic_if_elseif.tmpl: n less than 10
		_, _ = result.WriteString(`n less than 10`)
		//basic_if_elseif.tmpl: !else if n < 100
	} else if n < 100 {
		//basic_if_elseif.tmpl: n less than 100
		_, _ = result.WriteString(`n less than 100`)
		//basic_if_elseif.tmpl: !else
	} else {
		//basic_if_elseif.tmpl: n bigger than 100
		_, _ = result.WriteString(`n bigger than 100`)
		//basic_if_elseif.tmpl: !end
	}
	//basic_if_elseif.tmpl:
	_, _ = result.WriteString(`
`)
	//basic_if_elseif.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLbasic_if_elseif evaluates a template .TemplateFile
func TMPLbasic_if_elseif(n int) string {
	html, err := TMPLERRbasic_if_elseif(n)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template basic_if_elseif.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRcomparison_with_gotemplates evaluates a template .TemplateFile
func TMPLERRcomparison_with_gotemplates(params TemplateParam) (string, error) {
	_template := "comparison_with_gotemplates.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//comparison_with_gotemplates.tmpl: <html>
	_, _ = result.WriteString(`<html>
`)
	//comparison_with_gotemplates.tmpl:     <head>
	_, _ = result.WriteString(`    <head>
`)
	//comparison_with_gotemplates.tmpl:         <title>{{s params.Title }}</title>
	_, _ = result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, _escape(params.Title)))
	//comparison_with_gotemplates.tmpl:     </head>
	_, _ = result.WriteString(`    </head>
`)
	//comparison_with_gotemplates.tmpl:     <body>
	_, _ = result.WriteString(`    <body>
`)
	//comparison_with_gotemplates.tmpl:         <h1>{{s params.Title }}</h1>
	_, _ = result.WriteString(fmt.Sprintf(`        <h1>%s</h1>
`, _escape(params.Title)))
	//comparison_with_gotemplates.tmpl:
	_, _ = result.WriteString(`        `)
	//comparison_with_gotemplates.tmpl: !if len(params.Subtitle) > 0
	if len(params.Subtitle) > 0 {
		//comparison_with_gotemplates.tmpl: <h2>{{ params.Subtitle }}</h1>
		_, _ = result.WriteString(fmt.Sprintf(`<h2>%v</h1>`, params.Subtitle))
		//comparison_with_gotemplates.tmpl: !end
	}
	//comparison_with_gotemplates.tmpl:
	_, _ = result.WriteString(`
`)
	//comparison_with_gotemplates.tmpl:         <ul>
	_, _ = result.WriteString(`        <ul>
`)
	//comparison_with_gotemplates.tmpl:
	_, _ = result.WriteString(`            `)
	//comparison_with_gotemplates.tmpl: !for _, item := range params.Items
	for _, item := range params.Items {
		//comparison_with_gotemplates.tmpl:
		_, _ = result.WriteString(`
`)
		//comparison_with_gotemplates.tmpl:                 <li> {{s item }}
		_, _ = result.WriteString(fmt.Sprintf(`                <li> %s
`, _escape(item)))
		//comparison_with_gotemplates.tmpl:
		_, _ = result.WriteString(`            `)
		//comparison_with_gotemplates.tmpl: !end
	}
	//comparison_with_gotemplates.tmpl:
	_, _ = result.WriteString(`
`)
	//comparison_with_gotemplates.tmpl:         </ul>
	_, _ = result.WriteString(`        </ul>
`)
	//comparison_with_gotemplates.tmpl:         <p>
	_, _ = result.WriteString(`        <p>
`)
	//comparison_with_gotemplates.tmpl:             Written {{d len(params.Items) }} items
	_, _ = result.WriteString(fmt.Sprintf(`            Written %d items
`, len(params.Items)))
	//comparison_with_gotemplates.tmpl:         </p>
	_, _ = result.WriteString(`        </p>
`)
	//comparison_with_gotemplates.tmpl:     </body>
	_, _ = result.WriteString(`    </body>
`)
	//comparison_with_gotemplates.tmpl: </html>
	_, _ = result.WriteString(`</html>
`)
	//comparison_with_gotemplates.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLcomparison_with_gotemplates evaluates a template .TemplateFile
func TMPLcomparison_with_gotemplates(params TemplateParam) string {
	html, err := TMPLERRcomparison_with_gotemplates(params)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template comparison_with_gotemplates.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRextends evaluates a template .TemplateFile
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
	//extends.tmpl:
	_, _ = result.WriteString(``)
	//extends.tmpl:     </body>
	_, _ = result.WriteString(`    </body>
`)
	//extends.tmpl: </html>
	_, _ = result.WriteString(`</html>
`)
	//extends.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLextends evaluates a template .TemplateFile
func TMPLextends(title string, something int) string {
	html, err := TMPLERRextends(title, something)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template extends.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRextends_embedded evaluates a template .TemplateFile
func TMPLERRextends_embedded(title string, something int) (string, error) {
	_template := "extends_embedded.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//extends_embedded.tmpl:
	_, _ = result.WriteString(``)
	//extends_embedded.tmpl: !#extends base_embedded
	//extends_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(``)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(``)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(``)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//extends_embedded.tmpl: <html>
	_, _ = result.WriteString(`<html>
`)
	//extends_embedded.tmpl:     <head>
	_, _ = result.WriteString(`    <head>
`)
	//extends_embedded.tmpl:         <title>{{s title }}</title>
	_, _ = result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, _escape(title)))
	//extends_embedded.tmpl:
	_, _ = result.WriteString(``)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//extends_embedded.tmpl: <script>
	_, _ = result.WriteString(`<script>
`)
	//extends_embedded.tmpl: alert("included")
	_, _ = result.WriteString(`alert("included")
`)
	//extends_embedded.tmpl: </script>
	_, _ = result.WriteString(`</script>
`)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(``)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//extends_embedded.tmpl:     </head>
	_, _ = result.WriteString(`    </head>
`)
	//extends_embedded.tmpl:     <body>
	_, _ = result.WriteString(`    <body>
`)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(``)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//extends_embedded.tmpl: <h1>Body!</h1>
	_, _ = result.WriteString(`<h1>Body!</h1>
`)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(``)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(``)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(`
`)
	//extends_embedded.tmpl:     </body>
	_, _ = result.WriteString(`    </body>
`)
	//extends_embedded.tmpl: </html>
	_, _ = result.WriteString(`</html>
`)
	//extends_embedded.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLextends_embedded evaluates a template .TemplateFile
func TMPLextends_embedded(title string, something int) string {
	html, err := TMPLERRextends_embedded(title, something)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template extends_embedded.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRnoncode_line_with_exclamation_mark evaluates a template .TemplateFile
func TMPLERRnoncode_line_with_exclamation_mark() (string, error) {
	_template := "noncode_line_with_exclamation_mark.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//noncode_line_with_exclamation_mark.tmpl: !s1 := "This lins is not a code line"
	_, _ = result.WriteString(`!s1 := "This lins is not a code line"
`)
	//noncode_line_with_exclamation_mark.tmpl: !s2 := "This *is* a line of code"
	s2 := "This *is* a line of code"
	//noncode_line_with_exclamation_mark.tmpl: {{s s2 }}
	_, _ = result.WriteString(fmt.Sprintf(`%s
`, _escape(s2)))
	//noncode_line_with_exclamation_mark.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLnoncode_line_with_exclamation_mark evaluates a template .TemplateFile
func TMPLnoncode_line_with_exclamation_mark() string {
	html, err := TMPLERRnoncode_line_with_exclamation_mark()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template noncode_line_with_exclamation_mark.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRreturn evaluates a template .TemplateFile
func TMPLERRreturn(a int) (string, error) {
	_template := "return.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//return.tmpl: a is {{d a }}
	_, _ = result.WriteString(fmt.Sprintf(`a is %d
`, a))
	//return.tmpl:
	_, _ = result.WriteString(``)
	return result.String(), nil
	//return.tmpl:
	_, _ = result.WriteString(`
`)
	//return.tmpl: This line is ignored
	_, _ = result.WriteString(`This line is ignored
`)
	//return.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLreturn evaluates a template .TemplateFile
func TMPLreturn(a int) string {
	html, err := TMPLERRreturn(a)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template return.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRwith_end_instead_of_brackets evaluates a template .TemplateFile
func TMPLERRwith_end_instead_of_brackets() (string, error) {
	_template := "with_end_instead_of_brackets.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//with_end_instead_of_brackets.tmpl: !for i:=0; i < 5; i++
	for i := 0; i < 5; i++ {
		//with_end_instead_of_brackets.tmpl: i={{d i }}
		_, _ = result.WriteString(fmt.Sprintf(`i=%d
`, i))
		//with_end_instead_of_brackets.tmpl: !end
	}
	//with_end_instead_of_brackets.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLwith_end_instead_of_brackets evaluates a template .TemplateFile
func TMPLwith_end_instead_of_brackets() string {
	html, err := TMPLERRwith_end_instead_of_brackets()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template with_end_instead_of_brackets.tmpl:" + err.Error())
	}
	return html
}

type Argument struct {
	Aaa string
	Bbb int
}

// TMPLERRwith_global_declaration evaluates a template .TemplateFile
func TMPLERRwith_global_declaration(arg Argument) (string, error) {
	_template := "with_global_declaration.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//with_global_declaration.tmpl: Aaa={{s arg.Aaa }}
	_, _ = result.WriteString(fmt.Sprintf(`Aaa=%s
`, _escape(arg.Aaa)))
	//with_global_declaration.tmpl: Bbb={{d arg.Bbb }}
	_, _ = result.WriteString(fmt.Sprintf(`Bbb=%d
`, arg.Bbb))
	//with_global_declaration.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLwith_global_declaration evaluates a template .TemplateFile
func TMPLwith_global_declaration(arg Argument) string {
	html, err := TMPLERRwith_global_declaration(arg)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template with_global_declaration.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRwith_percentage evaluates a template .TemplateFile
func TMPLERRwith_percentage(str string) (string, error) {
	_template := "with_percentage.tmpl"
	_ = _template
	_escape := html.EscapeString
	_ = _escape
	var result bytes.Buffer
	//with_percentage.tmpl: %, str={{s str }}
	_, _ = result.WriteString(fmt.Sprintf(`%%, str=%s
`, _escape(str)))
	//with_percentage.tmpl: %, str={{s fmt.Sprintf("aaa%sccc", "bbb") }}
	_, _ = result.WriteString(fmt.Sprintf(`%%, str=%s
`, _escape(fmt.Sprintf("aaa%sccc", "bbb"))))
	//with_percentage.tmpl:
	_, _ = result.WriteString(``)

	return result.String(), nil
}

// TMPLwith_percentage evaluates a template .TemplateFile
func TMPLwith_percentage(str string) string {
	html, err := TMPLERRwith_percentage(str)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template with_percentage.tmpl:" + err.Error())
	}
	return html
}
