package gover

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
			readStr := string(readBytes)

			re, _ := regexp.Compile("^mode: [a-z]+\n")
			if re.Match(buffer.Bytes()) {
				readStr = re.ReplaceAllString(readStr, "")
			}

			if strings.HasPrefix(buffer.String(), "mode:") {

			}
			buffer.WriteString(readStr)
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
