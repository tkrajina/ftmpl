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
func TE__base(title string) (string, error) {
	__template__ := "base.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* <html> */
	result.WriteString(`<html>
`)
	/* <head> */
	result.WriteString(`    <head>
`)
	/* <title>{{s title }}</title> */
	result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, __escape__(title)))
	/* !#include head */
	/* </head> */
	result.WriteString(`    </head>
`)
	/* <body> */
	result.WriteString(`    <body>
`)
	/* !#include body */
	/* !#include footer */
	/* </body> */
	result.WriteString(`    </body>
`)
	/* </html> */
	result.WriteString(`</html>
`)
	/*  */
	result.WriteString(``)

	return result.String(), nil
}

func T__base(title string) string {
	html, err := TE__base(title)
	if err != nil {
		os.Stderr.WriteString("Error running template base.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__base_embedded(title string) (string, error) {
	__template__ := "base_embedded.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/*  */
	result.WriteString(``)
	/*  */
	result.WriteString(`
`)
	/* <html> */
	result.WriteString(`<html>
`)
	/* <head> */
	result.WriteString(`    <head>
`)
	/* <title>{{s title }}</title> */
	result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, __escape__(title)))
	/*  */
	result.WriteString(``)
	/* !#include head */
	/*  */
	result.WriteString(`
`)
	/* </head> */
	result.WriteString(`    </head>
`)
	/* <body> */
	result.WriteString(`    <body>
`)
	/*  */
	result.WriteString(``)
	/* !#include body */
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(``)
	/* !#include footer */
	/*  */
	result.WriteString(`
`)
	/* </body> */
	result.WriteString(`    </body>
`)
	/* </html> */
	result.WriteString(`</html>
`)
	/*  */
	result.WriteString(``)

	return result.String(), nil
}

func T__base_embedded(title string) string {
	html, err := TE__base_embedded(title)
	if err != nil {
		os.Stderr.WriteString("Error running template base_embedded.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__basic(str string, num int) (string, error) {
	__template__ := "basic.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* String:{{s str}} */
	result.WriteString(fmt.Sprintf(`String:%s
`, __escape__(str)))
	/* Unescaped:{{=s str}} */
	result.WriteString(fmt.Sprintf(`Unescaped:%s
`, str))
	/* Num:{{d num}} */
	result.WriteString(fmt.Sprintf(`Num:%d
`, num))
	/*  */
	result.WriteString(``)

	return result.String(), nil
}

func T__basic(str string, num int) string {
	html, err := TE__basic(str, num)
	if err != nil {
		os.Stderr.WriteString("Error running template basic.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__basic_code(s string, num int) (string, error) {
	__template__ := "basic_code.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* !for i := 0; i < num; i++ { */
	for i := 0; i < num; i++ {

		/* {{d i}} */
		result.WriteString(fmt.Sprintf(`%d
`, i))
		/* !} */
	}

	/*  */
	result.WriteString(``)

	return result.String(), nil
}

func T__basic_code(s string, num int) string {
	html, err := TE__basic_code(s, num)
	if err != nil {
		os.Stderr.WriteString("Error running template basic_code.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__basic_embedded_code(n int) (string, error) {
	__template__ := "basic_embedded_code.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/*  */
	result.WriteString(``)
	/* !for i := 0; i < n; i++{ */
	for i := 0; i < n; i++ {
		/* i={{d i }} */
		result.WriteString(fmt.Sprintf(`i=%d `, i))
		/* !} */
	}
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(``)

	return result.String(), nil
}

func T__basic_embedded_code(n int) string {
	html, err := TE__basic_embedded_code(n)
	if err != nil {
		os.Stderr.WriteString("Error running template basic_embedded_code.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__basic_if_else(n int) (string, error) {
	__template__ := "basic_if_else.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/*  */
	result.WriteString(``)
	/* !if n < 10 { */
	if n < 10 {
		/* {{d n }} less than 10 */
		result.WriteString(fmt.Sprintf(`%d less than 10`, n))
		/* !} */
	}
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(``)
	/* !if n > 0{ */
	if n > 0 {
		/* {{d n}} biger than 0 */
		result.WriteString(fmt.Sprintf(`%d biger than 0`, n))
		/* !} */
	}
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(``)
	/* !if n > 5{ */
	if n > 5 {
		/* {{d n}} biger than 5 */
		result.WriteString(fmt.Sprintf(`%d biger than 5`, n))
		/* !} else { */
	} else {
		/* {{d n}} smaller than 5 */
		result.WriteString(fmt.Sprintf(`%d smaller than 5`, n))
		/* !} */
	}
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(``)

	return result.String(), nil
}

func T__basic_if_else(n int) string {
	html, err := TE__basic_if_else(n)
	if err != nil {
		os.Stderr.WriteString("Error running template basic_if_else.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__basic_if_elseif(n int) (string, error) {
	__template__ := "basic_if_elseif.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/*  */
	result.WriteString(``)
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(``)
	/* !if n < 10{ */
	if n < 10 {
		/* n less than 10 */
		result.WriteString(`n less than 10`)
		/* !} else if n < 100 { */
	} else if n < 100 {
		/* n less than 100 */
		result.WriteString(`n less than 100`)
		/* !} else { */
	} else {
		/* n bigger than 100 */
		result.WriteString(`n bigger than 100`)
		/* !} */
	}
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(``)

	return result.String(), nil
}

func T__basic_if_elseif(n int) string {
	html, err := TE__basic_if_elseif(n)
	if err != nil {
		os.Stderr.WriteString("Error running template basic_if_elseif.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__comparison_with_gotemplates(params TemplateParam) (string, error) {
	__template__ := "comparison_with_gotemplates.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* <html> */
	result.WriteString(`<html>
`)
	/* <head> */
	result.WriteString(`    <head>
`)
	/* <title>{{s params.Title }}</title> */
	result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, __escape__(params.Title)))
	/* </head> */
	result.WriteString(`    </head>
`)
	/* <body> */
	result.WriteString(`    <body>
`)
	/* <h1>{{s params.Title }}</h1> */
	result.WriteString(fmt.Sprintf(`        <h1>%s</h1>
`, __escape__(params.Title)))
	/*  */
	result.WriteString(`        `)
	/* !if len(params.Subtitle) > 0{ */
	if len(params.Subtitle) > 0 {
		/* <h2>{{ params.Subtitle }}</h1> */
		result.WriteString(fmt.Sprintf(`<h2>%v</h1>`, params.Subtitle))
		/* !} */
	}
	/*  */
	result.WriteString(`
`)
	/* <ul> */
	result.WriteString(`        <ul>
`)
	/*  */
	result.WriteString(`            `)
	/* !for _, item := range params.Items{ */
	for _, item := range params.Items {
		/*  */
		result.WriteString(`
`)
		/* <li> {{s item }} */
		result.WriteString(fmt.Sprintf(`                <li> %s
`, __escape__(item)))
		/*  */
		result.WriteString(`            `)
		/* !} */
	}
	/*  */
	result.WriteString(`
`)
	/* </ul> */
	result.WriteString(`        </ul>
`)
	/* <p> */
	result.WriteString(`        <p>
`)
	/* Written {{d len(params.Items) }} items */
	result.WriteString(fmt.Sprintf(`            Written %d items
`, len(params.Items)))
	/* </p> */
	result.WriteString(`        </p>
`)
	/* </body> */
	result.WriteString(`    </body>
`)
	/* </html> */
	result.WriteString(`</html>
`)
	/*  */
	result.WriteString(``)

	return result.String(), nil
}

func T__comparison_with_gotemplates(params TemplateParam) string {
	html, err := TE__comparison_with_gotemplates(params)
	if err != nil {
		os.Stderr.WriteString("Error running template comparison_with_gotemplates.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__extends(title string, something int) (string, error) {
	__template__ := "extends.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* !#extends base */
	/*  */
	result.WriteString(`
`)
	/* <html> */
	result.WriteString(`<html>
`)
	/* <head> */
	result.WriteString(`    <head>
`)
	/* <title>{{s title }}</title> */
	result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, __escape__(title)))
	/* <script> */
	result.WriteString(`<script>
`)
	/* alert("included") */
	result.WriteString(`alert("included")
`)
	/* </script> */
	result.WriteString(`</script>
`)
	/*  */
	result.WriteString(`
`)
	/* </head> */
	result.WriteString(`    </head>
`)
	/* <body> */
	result.WriteString(`    <body>
`)
	/* <h1>Body!</h1> */
	result.WriteString(`<h1>Body!</h1>
`)
	/*  */
	result.WriteString(``)
	/* </body> */
	result.WriteString(`    </body>
`)
	/* </html> */
	result.WriteString(`</html>
`)
	/*  */
	result.WriteString(``)

	return result.String(), nil
}

func T__extends(title string, something int) string {
	html, err := TE__extends(title, something)
	if err != nil {
		os.Stderr.WriteString("Error running template extends.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__extends_embedded(title string, something int) (string, error) {
	__template__ := "extends_embedded.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/*  */
	result.WriteString(``)
	/* !#extends base_embedded */
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(``)
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(``)
	/*  */
	result.WriteString(``)
	/*  */
	result.WriteString(`
`)
	/* <html> */
	result.WriteString(`<html>
`)
	/* <head> */
	result.WriteString(`    <head>
`)
	/* <title>{{s title }}</title> */
	result.WriteString(fmt.Sprintf(`        <title>%s</title>
`, __escape__(title)))
	/*  */
	result.WriteString(``)
	/*  */
	result.WriteString(`
`)
	/* <script> */
	result.WriteString(`<script>
`)
	/* alert("included") */
	result.WriteString(`alert("included")
`)
	/* </script> */
	result.WriteString(`</script>
`)
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(``)
	/*  */
	result.WriteString(`
`)
	/* </head> */
	result.WriteString(`    </head>
`)
	/* <body> */
	result.WriteString(`    <body>
`)
	/*  */
	result.WriteString(``)
	/*  */
	result.WriteString(`
`)
	/* <h1>Body!</h1> */
	result.WriteString(`<h1>Body!</h1>
`)
	/*  */
	result.WriteString(``)
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(``)
	/*  */
	result.WriteString(`
`)
	/* </body> */
	result.WriteString(`    </body>
`)
	/* </html> */
	result.WriteString(`</html>
`)
	/*  */
	result.WriteString(``)

	return result.String(), nil
}

func T__extends_embedded(title string, something int) string {
	html, err := TE__extends_embedded(title, something)
	if err != nil {
		os.Stderr.WriteString("Error running template extends_embedded.tmpl:" + err.Error())
	}
	return html
}
