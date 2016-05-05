package invalid

import (
"bytes"
"errors"
"html"
"fmt"
"os"
)

func init() {
	_ = fmt.Sprintf
	_ = errors.New
	_ = os.Stderr
	_ = html.EscapeString
}



// Generated code, do not edit!!!!

func TE__invalid(s int) (string, error) {
	__template__ := "invalid.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
//invalid.tmpl: <html>
result.WriteString(`<html>
`)
//invalid.tmpl:     <head>
result.WriteString(`    <head>
`)
//invalid.tmpl:     </head>
result.WriteString(`    </head>
`)
//invalid.tmpl:     <body>
result.WriteString(`    <body>
`)
//invalid.tmpl:         {{s "\ " }}
result.WriteString(fmt.Sprintf(`        %s
`, __escape__( "\ ")))
//invalid.tmpl:     </body>
result.WriteString(`    </body>
`)
//invalid.tmpl: </html>
result.WriteString(`</html>
`)
//invalid.tmpl: 
result.WriteString(``)

	return result.String(), nil
}


func T__invalid(s int) string {
	html, err := TE__invalid(s)
	if err != nil {
		os.Stderr.WriteString("Error running template invalid.tmpl:" + err.Error())
	}
	return html
}