package gover_test

import (
	. "github.com/modocache/gover/gover"

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
				It("creates a file at out with its content", func() {
					Gover(root, out)
					Expect(readFile(out)).To(Equal("mario\n"))
				})
			})

			Context("and it contains several .coverprofile files at the root level", func() {
				BeforeEach(func() {
					root = filepath.Join(fixturesDir(), "root_cover_profiles")
				})
				It("creates a file at out with their contents (in filename alphabetical order)", func() {
					Gover(root, out)
					Expect(readFile(out)).To(Equal("mario\na link to the past\n"))
				})
			})

			Context("and it contains several .coverprofile files at several levels", func() {
				BeforeEach(func() {
					root = filepath.Join(fixturesDir(), "nested_cover_profiles")
				})
				It("creates a file at out with their contents", func() {
					Gover(root, out)
					Expect(readFile(out)).To(Equal("mario\na link to the past\nsonic\n"))
				})
			})
		})
	})
})
