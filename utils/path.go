package utils

import (
	"log"
	"os"
	"path/filepath"
)

var DirPath string

func Path() {
	var path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	DirPath = path
}
