package yargs

import (
	"fmt"
	"os"
	"path/filepath"
)

func FirstArgFileContent() (fileName, baseName string, content []byte, err error) {
	if len(os.Args) < 2 {
		return "", "", nil, fmt.Errorf("missing first arg - file name")
	}
	fileName = os.Args[1]
	base := filepath.Base(fileName)
	b, err := os.ReadFile(fileName)
	if err != nil {
		return fileName, base, nil, fmt.Errorf("couldn' read file '%+v'", fileName)
	}

	return fileName, base, b, nil
}
