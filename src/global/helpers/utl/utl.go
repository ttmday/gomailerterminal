package utl

import (
	"io/ioutil"
	"os"
	"strings"
)

func Mapping(start int, args []string, f func([]string)) {
	var m []string
	for _, r := range args[start:] {
		if strings.HasPrefix(r, "-") {
			break
		}

		m = append(m, r)
	}

	if len(m) > 0 {
		f(m)
	}
}

func LoadFile(filename string) (string, error) {

	file, err := os.OpenFile(filename, os.O_RDWR, 0o600)
	if err != nil {
		return "", err
	}

	defer file.Close()
	b, err := ioutil.ReadAll(file)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(b)), nil
}
