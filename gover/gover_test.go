package gover_test

import (
	. "github.com/modocache/gover/gover"

	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"path/filepath"
)

func fixturesDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return filepath.Join(dir, "_fixtures")
}

func readFile(path string) string {
	bytes, _ := ioutil.ReadFile(path)
	return string(bytes)
}

var _ = Describe("Gover", func() {
	Describe("Gover()", func() {
		var root, out string

		BeforeEach(func() {
			tmpfile, _ := ioutil.TempFile("", "")
			out = tmpfile.Name()
		})
		AfterEach(func() {
			os.Remove(out)
		})

		Context("when the root directory does not exist", func() {
			BeforeEach(func() {
				root = ".!this$file^shouldn't\"exist."
			})
			It("creates an empty file at out", func() {
				Gover(root, out)
				Expect(readFile(out)).To(Equal(""))
			})
		})

		Context("when the root directory does exist", func() {
			Context("and it contains no .coverprofile files", func() {
				BeforeEach(func() {
					root = filepath.Join(fixturesDir(), "no_cover_profiles")
				})
				It("creates an empty file at out", func() {
					Gover(root, out)
					Expect(readFile(out)).To(Equal(""))
				})
			})

			Context("and it contains one .coverprofile file at the root level", func() {
				BeforeEach(func() {
					root = filepath.Join(fixturesDir(), "root_cover_profile")
				})
				It("writes its content to out", func() {
					Gover(root, out)
					Expect(readFile(out)).To(Equal("mode: set\nmario\n"))
				})
			})

			Context("and it contains several .coverprofile files at the root level", func() {
				BeforeEach(func() {
					root = filepath.Join(fixturesDir(), "root_cover_profiles")
				})

				It("writes their content to out, using the mode of the first alphabetically", func() {
					Gover(root, out)
					Expect(readFile(out)).To(Equal("mode: set\nmario\na link to the past\n"))
				})
			})

			Context("and it contains several .coverprofile files at several levels", func() {
				BeforeEach(func() {
					root = filepath.Join(fixturesDir(), "nested_cover_profiles")
				})
				It("writes their content to out", func() {
					Gover(root, out)
					Expect(readFile(out)).To(Equal("mode: set\nmario\na link to the past\nsonic\n"))
				})
			})

			Context("and it contains some .coverprofile files that can't be read", func() {
				var tmpPath string
				BeforeEach(func() {
					root = filepath.Join(fixturesDir(), "nested_cover_profiles")
					tmpPath = filepath.Join(root, "tmp.coverprofile")
					ioutil.WriteFile(tmpPath, []byte("unreadable!"), os.ModeAppend)
				})
				AfterEach(func() {
					os.Remove(tmpPath)
				})
				It("writes their content to out", func() {
					fmt.Println("\ngover_test.go: should display warning:")
					Gover(root, out)
					Expect(readFile(out)).To(Equal("mode: set\nmario\na link to the past\nsonic\n"))
				})
			})
		})
	})
})
