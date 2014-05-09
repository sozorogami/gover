package gover

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	Extension = ".coverprofile"
)

func Gover(root, out string) {
	var buffer bytes.Buffer

	walkFn := func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != Extension {
			return err
		}

		readBytes, readErr := ioutil.ReadFile(path)
		if readErr == nil {
			buffer.WriteString(string(readBytes))
		} else {
			log.Println("gover: Could not read file:", path)
		}

		return err
	}

	filepath.Walk(root, walkFn)
	err := ioutil.WriteFile(out, buffer.Bytes(), 0666)
	if err != nil {
		log.Fatal("gover: Could not write to out:", out)
	}
}
