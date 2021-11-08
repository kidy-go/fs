// filesystem.go kee > 2021/10/31

package fs

import (
	"os"
)

type Filesystem interface {
	Write(filename string, content []byte, options ...interface{}) error

	Append(filename string, content []byte, options ...interface{}) error

	Move(src, dst string) error

	Copy(src, dst string) error

	Delete(src string) error

	Mkdir(folderPath string, perm os.FileMode) error

	Has(filename string) bool

	Get(filename string) ([]byte, error)

	Directories(path string) []string

	Files(path string) []string

	SetVisibility(filename string, visibility string) error

	GetVisibility(filename string) string

	Metadata(filename string) Metadata

	// TODO..
	// PutFileAs(filename string, fd os.File, dst string) error
}
