package helpers

import (
	"io/ioutil"
	"strings"
)

func ReadFile(filename string) (rows []string, err error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	rows = strings.Split(string(content), "\n")
	return
}
