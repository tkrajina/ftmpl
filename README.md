# FTMPL a fast (and typesafe) templating for Golang

ftmpl is a fast/compiled/typesafe templating "language" for golang.

## Templating files and functions

Every template must be saved in a file with extension `.tmpl`.
The template file can be "compiled" in Go code (and this is typically done in the build procedure).

If the file is names `my_template.tmpl` then the `ftmpl` utility will produce a function function `T__my_template()`.

Compiling the files is done my invoking:

    ftmpl -targetgo target_dir/templates_generated.go source_dir

The resulting code is pretty straightforward, and you will probably need to check (but not edit!) it in case of errors.

For template examples see [here](example/) (`.tmpl` files).
For the "compiled" golang code from those examples [see here](example/compiled.go).

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

Placeholders are expressions to be executed and shown when calling the template. 
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

Note that `{{v expression }}` *is not escaped* even when the resulting expression is a string!

## Code

There are two ways to write go code in templates:

## Lines starting with "!"

    !#arg n int
    <html>
        <body>
    !for i := 0; i < n; i++ {
            i is now {{d i }}
    !}
        </body>
    </html>

Basically every line starting with `!` (but not `!#`) will be translated as go code in the resulting function.

## Embedded code

    !#arg n int
    <html>
        <body>
    {{! for i := 0; i < n; i++ { }}i is now {{d i }}{{! } }}
        </body>
    </html>

...so by using {{! code }} you embed a go command in that place.

Note that (unlike some other templating languages the `{{! code }}` notation *is not multiline*. It must contain only one go command/line.

For `for` and `if` control structures you can ommit `{`, `}` or `} else {` (open and close block of code).
Use `end` to close and `else` (with `if`) instead:

For example `if`:

    {{! if n > 0 }}{{d n}} biger than 0{{! end }}
    {{! if n > 5 }}{{d n}} biger than 5{{! else }}{{d n}} smaller than 5{{! end }}

...or with `for`...

    {{! for i := 0; i < n; i++ }} i={{d i }} {{! end }}

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

And then various pages can be:

    !#extends base
    !#arg something int

    !#sub head
    <script>
    alert("included")
    </script>

    !#sub body
    <h1>Body!</h1>

If the extended template is names `my_page.tmpl` then the resulting template function will have both the arguments from *that template* and arguments from the *base template: `func T__my_page(title string, something int)`.

Typically you will have many template arguments, so the best way to deal with them is to pack them all into one "base page structure" and another struct for every page.
The function will then be `func T__my_page(baseParams BasePageParams, pageParams MyPageParams)`

## Throwing errors

Every template will result in *two* methods. Both will execute the same code, but:

 * Use `func TE__my_template(args) (string, error)` if you expect your template to return errors, use this one. Use `!return "", err` to prematurely exit from the template with a propper error value.
 * With `func T__my_template(args) string` if template returns an error, the result is an empty string, and the error will be written in `os.Syserr`.

## Special variables

`__template__` is defined in the function and contains the current template name.

# License

ftmpl is licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

