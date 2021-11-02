// filesystem.go kee > 2021/10/31

package storage

import (
	"os"
)

type Filesystem interface {
	Write(filename string, content []byte, options ...interface{}) error

	Append(filename string, content []byte, options ...interface{}) error

	Rename(oldname string, newname string) error

	Copy(src string, dst string) error

	Delete(src string) error

	Mkdir(folderPath string, perm os.FileMode) error

	Has(filename string) bool

	Get(filename string) ([]byte, error)

	ListContents(path string, deepLevel int) []string

	Metadata() map[string]string

	Size() int64
}
