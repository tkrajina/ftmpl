package ftmplting

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const randomStringChars = "qwertyuiopasdfghjklzxcvbnm1234567890"

func getRandomString(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result += string(randomStringChars[rand.Int()%len(randomStringChars)])
	}

	return result
}

func getLastPathElements(path string) (string, string) {
	absPath, err := filepath.Abs(path)
	handleError(err, "Error getting absolute path from "+path)
	pathParts := strings.Split(absPath, string(os.PathSeparator))

	var elements []string
	for _, part := range pathParts {
		part = strings.TrimSpace(part)
		if len(part) > 0 {
			elements = append(elements, part)
		}
	}

	if len(pathParts) == 0 {
		return "", ""
	}

	if len(elements) == 1 {
		return "", elements[len(elements)-1]
	}

	return elements[len(elements)-2], elements[len(elements)-1]
}

func quote(s string) string {
	return strings.Replace(s, "`", "`+\"`\"+`", -1)
}

func handleError(err error, message string) {
	if err == nil {
		return
	}
	panic(message + ":" + err.Error())
}

func handleLineError(err error, line string) {
	if err == nil {
		return
	}

	panic("Error " + err.Error() + " in line:" + line)
}

func rmDoubleImports(list []string) []string {
	var result []string
	existing := make(map[string]bool)
	for _, i := range list {
		if _, ok := existing[i]; !ok {
			result = append(result, i)
		}
		existing[i] = true
	}
	return result
}

func loadFile(filename string) (string, error) {
	f, err := os.Open(path.Join(filename))
	if err != nil {
		return "", fmt.Errorf("Error reading %s: %s", filename, err.Error())
	}
	defer f.Close()

	byts, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("Error reading %s: %s", filename, err.Error())
	}

	return string(byts), nil
}
