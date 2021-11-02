// filesystem.go kee > 2021/10/31

package storage

import (
	"time"
)

type Filesystem interface {
	Write(filename string, content []byte, options ...interface{}) error

	Append(filename string, content []byte, options ...interface{}) error

	Rename(oldname string, newname string) error

	Copy(src string, dst string) error

	Delete(src string) error

	Has(filename string) bool

	Read(filename string) []byte

	ListContents(path string, deepLevel int) []string

	Metadata() map[string]string

	Mimetype() string

	Size() int64

	Timestamp() time.Time
}
