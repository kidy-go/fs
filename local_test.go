// local_test.go kee > 2021/11/05

package storage_test

import (
	"github.com/kidy-go/storage"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

var rootPath = ".temp/"
var local = storage.NewLocal(rootPath)

func TestLocal(t *testing.T) {
	assert := assert.New(t)
	fn := "text/a/b/c/d/new_file.txt"

	// Testing local.Write
	text1 := "hello world"
	assert.Equal(nil, local.Write(fn, []byte(text1)))

	// Testing local.Get
	getTxt, _ := local.Get(fn)
	assert.Equal(text1, string(getTxt))

	// Testing local.Append
	appendTxt := " append"
	local.Append(fn, []byte(appendTxt))
	getTxt2, _ := local.Get(fn)
	assert.Equal(text1+appendTxt, string(getTxt2))

	// Testing local.SetVisibility && local.GetVisibility
	assert.Equal(nil, local.SetVisibility(fn, "private"))
	assert.Equal("private", local.GetVisibility(fn))

	assert.Equal(nil, local.SetVisibility(fn, "public"))
	assert.Equal("public", local.GetVisibility(fn))

	// Testing local.Mkdir
	folderPath, pathPermission := "c/d/e/f", os.ModePerm
	assert.Equal(nil, local.Mkdir(folderPath, pathPermission))
	assert.Equal(true, local.Has(folderPath))
	stat, _ := os.Stat(filepath.Join(rootPath, folderPath))
	assert.Equal(true, stat.IsDir())

	// Testing local.Copy && local.Move
	dst1, dst2 := "text/copyed.txt", "text/moved.txt"

	assert.Equal(false, local.Has(dst1))
	assert.Equal(false, local.Has(dst2))

	assert.Equal(nil, local.Copy(fn, dst1))
	assert.Equal(nil, local.Move(fn, dst2))

	assert.Equal(false, local.Has(fn))
	assert.Equal(true, local.Has(dst1))
	assert.Equal(true, local.Has(dst2))

	// Testing local.Delete
	assert.Equal(nil, local.Delete(dst1))
	assert.Equal(nil, local.Delete(dst2))
	assert.Equal(false, local.Has(dst1))
	assert.Equal(false, local.Has(dst2))

	files := []string{
		"test/a/b/c/1.txt",
		"test/b/c/d/2.txt",
		"test/c/d/e/3.txt",
	}
	for _, fn := range files {
		local.Write(fn, []byte(fn))
	}
	// Testing local.Files
	assert.Equal(files, local.Files("test"))

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
	assert.Equal(dirs, local.Directories("test"))

	// test local.Metadata
	fn, txt := "test.png", "Apply png File"
	dir := "mkdir_dir"
	mode := os.FileMode(0754)
	local.Write(fn, []byte(txt), mode)
	local.Mkdir(dir, mode)
	meta := local.Metadata(fn)
	metadir := local.Metadata(dir)

	assert.Equal(filepath.Join(rootPath, fn), meta.Path)
	assert.Equal(mode, meta.PermMode)

	assert.Equal("file", meta.Type)
	assert.Equal("dir", metadir.Type)

	assert.Equal(".png", meta.Extension)
	assert.Equal("", metadir.Extension)

	mime, _ := meta.MIMEType()
	assert.Equal("text/plain; charset=utf-8", mime.String())

	local.Delete("test")
	local.Delete("text")
	local.Delete(fn)
	local.Delete(dir)
}
