package ftmplting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLines(t *testing.T) {
	assert.Equal(t, []string{"", "!bbb\n", "ccc\nddd"}, getChunks("!bbb\nccc\nddd"))
	assert.Equal(t, []string{"\n", "!bbb\n", "ccc\nddd"}, getChunks("\n!bbb\nccc\nddd"))
	assert.Equal(t, []string{"000\n", "!bbb\n", "ccc\nddd"}, getChunks("000\n!bbb\nccc\nddd"))
	assert.Equal(t, []string{"000\naaa\n", "!bbb\n", "ccc\nddd"}, getChunks("000\naaa\n!bbb\nccc\nddd"))
	assert.Equal(t, []string{"000\naaa\n", "!bbb", ""}, getChunks("000\naaa\n!bbb"))
}
