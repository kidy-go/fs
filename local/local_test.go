// local_test.go kee > 2021/11/05

package local_test

import (
	"github.com/kidy-go/fs/local"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

var rootPath = ".temp/"
var lc = local.NewLocal(rootPath)

func TestLocal(t *testing.T) {
	assert := assert.New(t)
	fn := "text/a/b/c/d/new_file.txt"

	// Testing lc.Write
	text1 := "hello world"
	assert.Equal(nil, lc.Write(fn, []byte(text1)))

	// Testing lc.Get
	getTxt, _ := lc.Get(fn)
	assert.Equal(text1, string(getTxt))

	// Testing lc.Append
	appendTxt := " append"
	lc.Append(fn, []byte(appendTxt))
	getTxt2, _ := lc.Get(fn)
	assert.Equal(text1+appendTxt, string(getTxt2))

	// Testing lc.SetVisibility && lc.GetVisibility
	assert.Equal(nil, lc.SetVisibility(fn, "private"))
	assert.Equal("private", lc.GetVisibility(fn))

	assert.Equal(nil, lc.SetVisibility(fn, "public"))
	assert.Equal("public", lc.GetVisibility(fn))

	// Testing lc.Mkdir
	folderPath, pathPermission := "c/d/e/f", os.ModePerm
	assert.Equal(nil, lc.Mkdir(folderPath, pathPermission))
	assert.Equal(true, lc.Has(folderPath))
	stat, _ := os.Stat(filepath.Join(rootPath, folderPath))
	assert.Equal(true, stat.IsDir())

	// Testing lc.Copy && lc.Move
	dst1, dst2 := "text/copyed.txt", "text/moved.txt"

	assert.Equal(false, lc.Has(dst1))
	assert.Equal(false, lc.Has(dst2))

	assert.Equal(nil, lc.Copy(fn, dst1))
	assert.Equal(nil, lc.Move(fn, dst2))

	assert.Equal(false, lc.Has(fn))
	assert.Equal(true, lc.Has(dst1))
	assert.Equal(true, lc.Has(dst2))

	// Testing lc.Delete
	assert.Equal(nil, lc.Delete(dst1))
	assert.Equal(nil, lc.Delete(dst2))
	assert.Equal(false, lc.Has(dst1))
	assert.Equal(false, lc.Has(dst2))

	files := []string{
		"test/a/b/c/1.txt",
		"test/b/c/d/2.txt",
		"test/c/d/e/3.txt",
	}
	for _, fn := range files {
		lc.Write(fn, []byte(fn))
	}
	// Testing lc.Files
	assert.Equal(files, lc.Files("test"))

	dirs := []string{
		"test/a",
		"test/a/b",
		"test/a/b/c",
		"test/b",
		"test/b/c",
		"test/b/c/d",
		"test/c",
		"test/c/d",
		"test/c/d/e",
	}
	assert.Equal(dirs, lc.Directories("test"))

	// test lc.Metadata
	fn, txt := "test.png", "Apply png File"
	dir := "mkdir_dir"
	mode := os.FileMode(0754)
	lc.Write(fn, []byte(txt), mode)
	lc.Mkdir(dir, mode)
	meta := lc.Metadata(fn)
	metadir := lc.Metadata(dir)

	assert.Equal(filepath.Join(rootPath, fn), meta.Path)
	assert.Equal(mode, meta.PermMode)

	assert.Equal("file", meta.Type)
	assert.Equal("dir", metadir.Type)

	assert.Equal(".png", meta.Extension)
	assert.Equal("", metadir.Extension)

	mime, _ := meta.MIMEType()
	assert.Equal("text/plain; charset=utf-8", mime.String())

	lc.Delete("test")
	lc.Delete("text")
	lc.Delete(fn)
	lc.Delete(dir)
}
