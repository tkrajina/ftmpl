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
