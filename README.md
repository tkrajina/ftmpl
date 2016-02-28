# FTMPL: fast (and typesafe) templating for Golang

ftmpl is a fast/compiled/typesafe templating "language" for golang.

## Why ftmpl?

 * It's faster (run `go test -v . -run=TestComparisonWithGolangTemplates` for a comparison). The builtin templates involve a lot of reflection, ftmpl "compiles" to simple Golang functions.
 * No need to learn another templating "language", ftmpl is just plain Go
 * Type safety (why have a typesafe language and a nontypesafe templating engine?)

## Installation

    go get github.com/tkrajina/ftmpl

## Templating files and functions

Every template must be saved in a file with extension `.tmpl`.
The template file can be "compiled" in Go code (and this is typically done in the build procedure).

If the file is named `my_template.tmpl` then the `ftmpl` utility will produce a function `T__my_template()`.

Compiling files is done my invoking:

    ftmpl -targetgo target_dir/templates_generated.go source_dir

The resulting code is pretty straightforward, and you will probably need to check (but not edit!) it in case of errors.

For template examples see [here](example/) (`.tmpl` files).
For the "compiled" golang code from those examples [see here](example/compiled.go). The resulting code is always formatted using `go fmt`.

If you prefer using `go generate` add a comment...

    //go:generate ftmpl -targetgo target_dir/templates_generated.go source_dir

...anywhere in your code.

## Arguments

An example template `show_time.tmpl`

    !#import "time"
    <html>
        <body>
            It's {{s time.Now().String() }}
        </body>
    </html>

It will be compiled to the function `func T__show_time()`

If you want your template function to have arguments:

    !#import "time"
    !#arg t time.Time
    <html>
        <body>
            It's {{s t.String() }}
        </body>
    </html>

Now the compiled function will be `func T__show_time(time.Time)`

## Placeholders

Placeholders are expressions to be executed and written when calling the template. 
Since ftmpl is a typesafe templating it needs to know of which type is the expression, and based on that to properly format it:

Use:

 * `{{s expression }}` for string expressions
 * `{{d expression }}` for number expressions
 * `{{t expression }}` for bool expressions
 * ...

It's simple, when compiled `{{s expression }}` will end up as `fmt.Sprintf("...%s...", expression)`, `{{v expression }}` as `fmt.Sprintf("...%v...", expression)`, etc.

`{{ expression }}` is the same as `{{v expression }}`.

### Escape and unsecape

If you use `{{s expresssion }}` the result will be escaped using `html.EscapeString()`. If you want to write *the exact same string* (without escaping), use `{{=s expression }}`.

**Important** that `{{v expression }}` *is not escaped* even when the resulting expression is a string!

## Code

There are two ways to write go code in templates:

### Lines starting with "!"

    !#arg n int
    <html>
        <body>
    !for i := 0; i < n; i++ {
            i is now {{d i }}
    !}
        </body>
    </html>

Basically every line starting with `!` (but not `!#`) will be translated as go code in the resulting function.

If you (really?) need a non-golang line starting with `!` start it with `!!`.

### Embedded code

    !#arg n int
    <html>
        <body>
            {{! for i := 0; i < n; i++ { }}i is now {{d i }}{{! } }}
        </body>
    </html>

...so by using {{! code }} you embed a go command in that place.

Note that (unlike some other templating languages) the `{{! code }}` notation *is not multiline*. It must contain only one go command/line.

For `for` and `if` control structures you can ommit `{`, `}` or `} else {` (open and close block of code).
Use `end` to close and `else` (with `if`) instead:

For example `if`:

    {{! if n > 0 }}{{d n}} biger than 0{{! end }}
    {{! if n > 5 }}{{d n}} biger than 5{{! else }}{{d n}} smaller than 5{{! end }}

...or with `for`...

    {{! for i := 0; i < n; i++ }} i={{d i }} {{! end }}

## Careful with Go code

To prematurely end the execution with an error (depending on an expression), you can use:

    !#errif <expression>, "<error_message>"

...or...

    !#return

## Careful with Go code

Ftmpl allows you to write go code (instead of templating "metalanguages", like other templating engines), but be careful to not write too much of it. 
Ideally, the data should be prepared before the template function, and the only code in the template would be some basic formatting, a couple of `if`s and `for`loops.

## Base templates and extensions

When you need a base website template, and pages "extending it":

    !#arg title string
    <html>
        <head>
            <title>{{s title }}</title>
    !#include head
        </head>
        <body>
    !#include body
    !#include footer
        </body>
    </html>

And then a typical page will be something like:

    !#extends base
    !#arg something int

    !#sub head
    <script>
    alert("included")
    </script>

    !#sub body
    <h1>Body!</h1>

If the name of the extended template is `my_page.tmpl` then the resulting template function will have both the arguments from *that template* and arguments from the *base template*: `func T__my_page(title string, something int)`.

Typically you will have many template arguments, so the best way to deal with them is to pack them all into one "base page structure" and another struct for every page.
The function will then be `func T__my_page(baseParams BasePageParams, pageParams MyPageParams)`

## Returning errors

Every template will result in *two* template functions. Both will execute the same code, but:

 * Use `func TE__my_template(args) (string, error)` if you expect your template to return errors, use this one. Use `!return "", err` to prematurely exit from the template with a proper error value.
 * With `func T__my_template(args) string` if template returns an error, the result is an empty string, and the error will be written in `os.Syserr`.

## Global declarations

In case you need a global declaration (outside of the template function), it can be done with `!#global`:

    !#global type Argument struct { Aaa string; Bbb int }
    !#arg arg Argument
    Aaa is {{s arg.Aaa }}, and Bbb is {{d arg.Bbb }}

The `Argument` struct will be declared outside the function body, and the function is now `T_my_template(Argument)`.

## Special variables

`__template__` is defined in the function and contains the current template name.

# License

ftmpl is licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

