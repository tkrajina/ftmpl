package ftmplting

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseGoFmtErrorOutput(out []byte) {
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		r := regexp.MustCompile("^(.*?):(\\d+):(\\d+):(.*)$")
		groups := r.FindStringSubmatch(line)
		fmt.Fprint(os.Stderr, line, "\n")
		if len(groups) > 0 {
			filename := groups[1]
			line, _ := strconv.ParseInt(groups[2], 10, 64)
			column, _ := strconv.ParseInt(groups[3], 10, 64)
			msg := strings.TrimSpace(groups[4])
			printTemplateErrDetails(filename, int(line), int(column), msg)
		}
	}
}

func printTemplateErrDetails(filename string, line, column int, msg string) {
	contents, err := loadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "> Cannot open %s: %s", filename, err.Error())
		return
	}

	lines := strings.Split(contents, "\n")

	context := 5
	fromLine := int(math.Max(float64(0), float64(line-context)))
	toLine := int(math.Min(float64(line+context), float64(len(lines)-1)))
	for i := fromLine; i < toLine; i++ {
		currLine := "   "
		if i+1 == line {
			currLine = "-->"
		}
		prefix := fmt.Sprintf("%s %5d:    ", currLine, i+1)
		fmt.Fprintln(os.Stderr, prefix+lines[i])
		if i+1 == line {
			fmt.Fprintln(os.Stderr, strings.Repeat(" ", len(prefix)+column-1)+"^")
		}
	}
}
