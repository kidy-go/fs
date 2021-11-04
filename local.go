// local.go kee > 2021/10/31

package storage

import (
	"fmt"
	"github.com/kidy-go/storage/utils"
	"io"
	"os"
	"path/filepath"
)

const (
	LS_MODE_ALL = -1 + iota
	LS_MODE_FILE
	LS_MODE_DIR
)

type Local struct {
	RootPath string
}

var permFiles, permDirs = map[string]os.FileMode{
	"public":  os.FileMode(0644),
	"private": os.FileMode(0600),
}, map[string]os.FileMode{
	"public":  os.FileMode(0755),
	"private": os.FileMode(0700),
}

func NewLocal(rootPath string) *Local {
	return &Local{rootPath}
}

func (lc *Local) applyPathPrefix(path string) string {
	return filepath.Join(lc.RootPath, path)
}

func (lc *Local) Write(filename string, content []byte, options ...interface{}) error {
	path := lc.applyPathPrefix(filename)

	// o := os.O_APPEND | os.O_WRONLY
	o, p := os.O_WRONLY, permFiles["public"] //os.ModePerm
	if _, op := utils.GetTypeOf(options, "FileMode"); op != nil {
		p = op.(os.FileMode)
	}

	if _, op := utils.GetTypeOf(options, "int"); op != nil {
		o = op.(int)
	}

	if _, op := utils.GetTypeOf(options, "string"); op != nil {
		if visibility, ok := permFiles[op.(string)]; ok {
			p = visibility
		}
	}

	if folderPath := filepath.Dir(filename); !lc.Has(folderPath) {
		lc.Mkdir(folderPath, os.ModePerm)
		lc.Chown(folderPath, os.ModePerm)
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
	folderPath = lc.applyPathPrefix(folderPath)
	return os.MkdirAll(folderPath, perm)
}

func (lc *Local) Chown(folderPath string, perm os.FileMode) error {
	folderPath = lc.applyPathPrefix(folderPath)
	return os.Chmod(folderPath, perm)
}

func (lc *Local) Append(filename string, content []byte, options ...interface{}) error {
	if len(options) == 0 {
		options = append(options, os.O_APPEND)
	}

	if i, op := utils.GetTypeOf(options, "FileMode"); op != nil {
		options[i] = op.(os.FileMode)
	}
	if i, op := utils.GetTypeOf(options, "int"); op != nil {
		options[i] = op.(int)
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
	path := lc.applyPathPrefix(filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (lc *Local) Get(filename string) ([]byte, error) {
	filename = lc.applyPathPrefix(filename)
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fd.Close()
	return io.ReadAll(fd)
}

func (lc *Local) SetVisibility(filename string, visibility string) {
	path := lc.applyPathPrefix(filename)

	if stat, err := os.Stat(path); err == nil {
		oPerm, ok := permFiles[visibility]
		if stat.IsDir() {
			oPerm, ok = permDirs[visibility]
		}

		if ok {
			lc.Chown(filename, oPerm)
		}
	}
}

func (lc *Local) GetVisibility(filename string) string {
	filename = lc.applyPathPrefix(filename)
	if stat, err := os.Stat(filename); err == nil {

		perm := fmt.Sprintf("%#o", stat.Mode())
		perm = perm[len(perm)-4:]

		fmt.Println("PERM >>>> ", perm)

		public := fmt.Sprintf("%#o", permFiles["public"])
		if stat.IsDir() {
			public = fmt.Sprintf("%#o", permDirs["public"])
		}

		if perm == public[len(public)-4:] {
			return "public"
		}
	}
	return "private"
}

func (lc *Local) Files(path string) []string {
	return lc.Walk(path, LS_MODE_FILE)
}

func (lc *Local) Directories(path string) []string {
	return lc.Walk(path, LS_MODE_DIR)
}

func (lc *Local) Walk(root string, lsMode int) []string {
	npath := lc.applyPathPrefix(root)
	npath = filepath.ToSlash(npath)
	rootIndex := len(filepath.ToSlash(lc.applyPathPrefix(""))) + 1

	var paths []string
	filepath.Walk(npath, func(path string, fstat os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		mode := LS_MODE_FILE
		if fstat.IsDir() {
			mode = LS_MODE_DIR
		}

		if mode != LS_MODE_ALL && mode != lsMode {
			return nil
		}

		path = filepath.ToSlash(path)
		if path = path[rootIndex:]; path != "" && root != path {
			paths = append(paths, path)
		}

		return nil
	})

	return paths
}

func (lc *Local) Metadata(filename string) Metadata {
	filename = lc.applyPathPrefix(filename)
	finfo, _ := FileInfo(filename)
	return finfo
}
