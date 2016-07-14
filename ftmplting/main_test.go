package ftmplting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLines(t *testing.T) {
	assert.Equal(t, []string{"", "!bbb\n", "ccc\nddd"}, getLines("!bbb\nccc\nddd"))
	assert.Equal(t, []string{"\n", "!bbb\n", "ccc\nddd"}, getLines("\n!bbb\nccc\nddd"))
	assert.Equal(t, []string{"000\n", "!bbb\n", "ccc\nddd"}, getLines("000\n!bbb\nccc\nddd"))
	assert.Equal(t, []string{"000\naaa\n", "!bbb\n", "ccc\nddd"}, getLines("000\naaa\n!bbb\nccc\nddd"))
	assert.Equal(t, []string{"000\naaa\n", "!bbb", ""}, getLines("000\naaa\n!bbb"))
}
