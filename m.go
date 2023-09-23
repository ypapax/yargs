package yargs

import (
	"fmt"
	"os"
)

func FirstArgFileContent() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("missing first arg - file name")
	}
	fileName := os.Args[1]
	b, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("couldn' read file '%+v'", fileName)
	}
	return string(b), nil
}
