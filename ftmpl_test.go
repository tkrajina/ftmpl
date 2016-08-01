package main_test

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"strings"
	"testing"
	texttmpl "text/template"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tkrajina/ftmpl/example"
)

func TestBasic(t *testing.T) {
	result := example.TMPLbasic("aaa", 10)
	expected := `String:aaa
Unescaped:aaa
Num:10
This {{ is ignored
So is this }} !`
	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Error("Expected:", expected, "was:", result)
	}
}

func TestWithPercentage(t *testing.T) {
	result := example.TMPLWithPercentage("something")
	expected := `%, str=something
%, str=aaabbbccc`
	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Error("Expected:", expected, "was:", result)
	}
}

func TestBasicEscaped(t *testing.T) {
	result := example.TMPLbasic("<aaa&...", 10)
	expected := `String:&lt;aaa&amp;...
Unescaped:<aaa&...
Num:10
This {{ is ignored
So is this }} !`
	if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
		t.Error(explanation)
	}
}

func TestWithExclamationMark(t *testing.T) {
	result := example.TMPLWithExclamationMark()
	expected := `Something here 5! And something here.
And something here: true! And here, too!! Hey, one more!!!`
	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Error("Expected:", expected, "was:", result)
	}
}

func TestWithDirectWriting(t *testing.T) {
	result := example.TMPLWithDirectWriting()
	expected := `This is Written directly to ftmplresult!`
	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Error("Expected:", expected, "was:", result)
	}
}

func TestBasicCode(t *testing.T) {
	result := example.TMPLBasicCode("aaa", 5)
	expected := `0
1
2
3
4`
	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Error("Expected:", expected, "was:", result)
	}
}

func TestBasicEmbeddedCode(t *testing.T) {
	result := example.TMPLBasicEmbeddedCode(3)
	expected := `i=0 i=1 i=2`
	if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
		t.Error(explanation)
	}
}

func TestBasicIfElse(t *testing.T) {
	{
		result := example.TMPLBasicIfElse(-10)
		expected := `-10 less than 10

-10 smaller than 5`
		if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
			t.Error(explanation)
		}
	}
	{
		result := example.TMPLBasicIfElse(3)
		expected := `3 less than 10
3 biger than 0
3 smaller than 5`
		if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
			t.Error(explanation)
		}
	}
	{
		result := example.TMPLBasicIfElse(100)
		expected := `100 biger than 0
100 biger than 5`
		if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
			t.Error(explanation)
		}
	}
}

func TestBasicIfElseIf(t *testing.T) {
	{
		result := example.TMPLBasicIfElseif(-5)
		expected := `n less than 10`
		if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
			t.Error(explanation)
		}
	}
	{
		result := example.TMPLBasicIfElseif(25)
		expected := `n less than 100`
		if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
			t.Error(explanation)
		}
	}
	{
		result := example.TMPLBasicIfElseif(2500)
		expected := `n bigger than 100`
		if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
			t.Error(explanation)
		}
	}
}

func TestExtendsWithLinePlaceholders(t *testing.T) {
	result := example.TMPLextends("naslov", 12)
	expected := `<html>
    <head>
        <title>naslov</title>
<script>
alert("included")
</script>

    </head>
    <body>
<h1>Body!</h1>
    </body>
</html>`
	if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
		t.Error(explanation)
	}
}

func TestExtendsWithEmbeddedPlaceholders(t *testing.T) {
	result := example.TMPLExtendsEmbedded("naslov", 12)
	expected := `<html>
    <head>
        <title>naslov</title>

<script>
alert("included")
</script>


    </head>
    <body>

<h1>Body!</h1>


    </body>
</html>`
	if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
		t.Error(explanation)
	}
}

var golangTemplate = texttmpl.Must(texttmpl.New("").Parse(`
<html>
    <head>
        <title>{{ .Title }}</title>
    </head>
    <body>
        <h1>{{ .Title }}</h1>
        {{ if .Subtitle }}<h2>{{ .Subtitle }}</h1>{{ end }}
        <ul>
            {{ range .Items }}
                <li> {{ . }}
            {{ end }}
        </ul>
        <p>
            Written {{ len .Items }} items
        </p>
    </body>
</html>
`))

func TestNonCodeStartingWithExclamationMark(t *testing.T) {
	result := example.TMPLNoncodeLineWithExclamationMark()
	expected := `!s1 := "This line is not a code line"
This *is* a line of code`
	if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
		t.Error(explanation)
	}
}

func TestWithGLobalDeclaration(t *testing.T) {
	result := example.TMPLWithGlobalDeclaration(example.Argument{Aaa: "a", Bbb: 10})
	expected := `Aaa=a
Bbb=10`
	if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
		t.Error(explanation)
	}
}

func TestWithReturn(t *testing.T) {
	result := example.TMPLreturn(17)
	expected := `a is 17
`
	if explanation, ok := linesEquals(strings.TrimSpace(expected), strings.TrimSpace(result)); !ok {
		t.Error(explanation)
	}
}

func TestComparisonWithGolangTemplates(t *testing.T) {
	param := example.TemplateParam{
		Title:    "titl",
		Subtitle: "subtitl",
		Items:    []string{"qwe", "asd", "jkl", "1", "2"},
	}
	var buf bytes.Buffer
	goTmplStarted := time.Now().Nanosecond()
	err := golangTemplate.Execute(&buf, param)
	assert.Nil(t, err)
	goTmplFinished := time.Now().Nanosecond()

	ftmplStarted := time.Now().Nanosecond()
	withFtmpl := example.TMPLComparisonWithGoTemplates(param)
	ftmplFinished := time.Now().Nanosecond()

	goTemplateTime := goTmplFinished - goTmplStarted
	ftmplTime := ftmplFinished - ftmplStarted
	log.Printf("goTemplateTime=%d, ftmplTime=%d (nanoseconds) ... ftmpl %f times faster\n", goTemplateTime, ftmplTime, float64(goTemplateTime)/float64(ftmplTime))
	// This is not exactly deterministic, but since ftmp is 2-10x faster than go templates, it is safe enough:
	if goTemplateTime < ftmplTime/2 {
		t.Error(fmt.Sprintf("golang templates %d, fmt templates %d", goTemplateTime, ftmplTime))
	}

	if explanation, ok := linesEquals(buf.String(), withFtmpl); !ok {
		t.Error(explanation)
	}
}

func TestInsert(t *testing.T) {
	withInsert := example.TMPLWithInsert(5)
	expected := `Will insert something here: a=5`
	if explanation, ok := linesEquals(withInsert, expected); !ok {
		t.Error(explanation)
	}
}

func TestFmtFormat(t *testing.T) {
	withInsert := example.TMPLFmtFormat()
	expected := `A simple int:10
A number:2.33
A padded string:    padded
A padded string #2:&amp;&amp;&amp;&amp;
A padded string #3:      &&&&
A +v formatting: [aaa]`
	if explanation, ok := linesEquals(withInsert, expected); !ok {
		t.Error(explanation)
	}
}

func TestNoBrackets(t *testing.T) {
	expected := `i=0
i=1
i=2
i=3
i=4`

	if explanation, ok := linesEquals(expected, example.TMPLWithEndInsteadOfBrackets()); !ok {
		t.Error(explanation)
	}
}

func linesEquals(str1, str2 string) (explanation string, equals bool) {
	if str1 == str2 {
		return "", true
	}

	str1 = strings.Replace(str1, " ", "[space]", -1)
	str1 = strings.Replace(str1, "\t", "[tab]", -1)
	str2 = strings.Replace(str2, " ", "[space]", -1)
	str2 = strings.Replace(str2, "\t", "[tab]", -1)

	lines1 := strings.Split(strings.TrimSpace(str1), "\n")
	lines2 := strings.Split(strings.TrimSpace(str2), "\n")

	if len(lines1) != len(lines2) {
		fmt.Println("----------------------------------------------------------------------------------------------------")
		fmt.Println(str1)
		fmt.Println("----------------------------------------------------------------------------------------------------")
		fmt.Println(str2)
		fmt.Println("----------------------------------------------------------------------------------------------------")
		return fmt.Sprintf("Lines count don't match %d!=%d", len(lines1), len(lines2)), false
	}

	for i := 0; i < len(lines1); i++ {
		line1 := lines1[i]
		line2 := lines2[i]
		for j := 0; j < int(math.Min(float64(len(line1)), float64(len(line2)))); j++ {
			ch1 := line1[j]
			ch2 := line2[j]
			if ch1 != ch2 {
				return fmt.Sprintf("Line #%d don't match \"%s\"!=\"%s\" in character #%d: %c!=%c", i, line1, line2, i, ch1, ch2), false
			}
		}
		if line1 != line2 {
			return fmt.Sprintf("Line #%d don't match \"%s\"!=\"%s\"", i, line1, line2), false
		}
	}

	return "", true
}
