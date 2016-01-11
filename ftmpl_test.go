package main_test

import (
	"strings"
	"testing"

	"github.com/tkrajina/ftmpl/example"
)

func TestBasic(t *testing.T) {
	result := example.T__basic("aaa", 10)
	expected := `String:aaa
Unescaped:aaa
Num:10`
	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Error("Expected:", expected, "was:", result)
	}
}

func TestBasicEscaped(t *testing.T) {
	result := example.T__basic("<aaa&...", 10)
	expected := `String:&lt;aaa&amp;...
Unescaped:<aaa&...
Num:10`
	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Error("Expected:", expected, "was:", result)
	}
}

func TestBasicCode(t *testing.T) {
	result := example.T__basic_code("aaa", 5)
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
	result := example.T__basic_embedded_code(3)
	expected := `i=0 i=1 i=2`
	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Error("Expected:", expected, "was:", result)
	}
}

func TestBasicIfElse(t *testing.T) {
	{
		result := example.T__basic_if_else(-10)
		expected := `-10 less than 10

-10 smaller than 5`
		if strings.TrimSpace(expected) != strings.TrimSpace(result) {
			t.Error("Expected:", expected, "was:", result)
		}
	}
	{
		result := example.T__basic_if_else(3)
		expected := `3 less than 10
3 biger than 0
3 smaller than 5`
		if strings.TrimSpace(expected) != strings.TrimSpace(result) {
			t.Error("Expected:", expected, "was:", result)
		}
	}
	{
		result := example.T__basic_if_else(100)
		expected := `100 biger than 0
100 biger than 5`
		if strings.TrimSpace(expected) != strings.TrimSpace(result) {
			t.Error("Expected:", expected, "was:", result)
		}
	}
}

func TestExtendsWithLinePlaceholders(t *testing.T) {
	result := example.T__extends("naslov", 12)
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
	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Error("Expected:", expected, "was:", result)
	}
}

func TestExtendsWithEmbeddedPlaceholders(t *testing.T) {
	result := example.T__extends_embedded("naslov", 12)
	expected := `



<html>
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
	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Error("Expected:", expected, "was:", result)
	}
}
