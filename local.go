// local.go kee > 2021/10/31

package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

type Local struct {
	RootPath string
}

func NewLocal(rootPath string) *Local {
	return &Local{rootPath}
}

func (lc *Local) Write(filename string, content []byte, options ...interface{}) error {
	path := filepath.Join(lc.RootPath, filename)

	// o := os.O_APPEND | os.O_WRONLY
	o, p := os.O_WRONLY, os.ModePerm
	for _, op := range options {
		switch reflect.TypeOf(op).Name() {
		case "FileMode":
			p = op.(os.FileMode)
		case "int":
			o = op.(int)
		}
	}

	if folderPath := filepath.Dir(filename); !lc.Has(folderPath) {
		lc.Mkdir(folderPath, p)
		lc.Chown(folderPath, p)
	}

	// 文件不存在则创建
	if !lc.Has(filename) {
		o = o | os.O_CREATE
	}

	// 写入参数
	if o != o|os.O_WRONLY && o != o|os.O_RDWR {
		o = o | os.O_WRONLY
	}

	// 非追加则清空
	if o != o|os.O_APPEND {
		o = o | os.O_TRUNC
	}

	fd, err := os.OpenFile(path, o, p)
	defer fd.Close()

	_, err = fd.Write(content)
	return err
}

func (lc *Local) Mkdir(folderPath string, perm os.FileMode) error {
	folderPath = filepath.Join(lc.RootPath, folderPath)
	return os.MkdirAll(folderPath, perm)
}

func (lc *Local) Chown(folderPath string, perm os.FileMode) error {
	folderPath = filepath.Join(lc.RootPath, folderPath)
	return os.Chmod(folderPath, perm)
}

func (lc *Local) Append(filename string, content []byte, options ...interface{}) error {
	if len(options) == 0 {
		options = append(options, os.O_APPEND)
	}

	for i, op := range options {
		switch reflect.TypeOf(op).Name() {
		case "FileMode":
			options[i] = op.(os.FileMode)
		case "int":
			options[i] = op.(int) | os.O_APPEND
		}
	}

	return lc.Write(filename, content, options...)
}

func (lc *Local) Rename(oldname, newname string) error {
	return nil
}

func (lc *Local) Copy(src, dst string) error {
	return nil
}

func (lc *Local) Delete(src string) error {
	return nil
}

func (lc *Local) Has(filename string) bool {
	path := filepath.Join(lc.RootPath, filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (lc *Local) Get(filename string) ([]byte, error) {
	filename = filepath.Join(lc.RootPath, filename)
	fl, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer fl.Close()
	return ioutil.ReadAll(fl)
}

func (lc *Local) ListContents(path string, deepLevel int) []string {
	return []string{}
}

func (lc *Local) Metadata() Metadata {
	return Metadata{}
}

func (lc *Local) Size() int {
	return 0
}
