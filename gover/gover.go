/*
Functions for walking a file tree and concatenating all files with the
".coverprofile" extension. Designed for use with the `gover` command-line tool.
*/
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
	/*
		The file extension to be concatenated. Currently gover assumes this
		will always be '.coverprofile'.
	*/
	Extension = ".coverprofile"
)

/*
Walks the file tree at `root`, concatenates all files ending with `Extension`,
then writes those files to `out`.

If `root` is an invalid path or does not contain any relevant files, an empty
string is written to `out`.

If `out` already exists, this function appends the concatenation to `out`. If
`out` does not exist, a new file is created with readwrite permissions (0666).
*/
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
