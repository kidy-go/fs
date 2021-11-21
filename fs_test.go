// fs_test.go kee > 2021/11/17

package fs_test

import (
	"fmt"
	"github.com/kidy-go/fs"
	"github.com/kidy-go/fs/local"
	"testing"
)

func init() {
	fs.Register("local", local.New(".temp"))
}

func TestFs(t *testing.T) {
	//lc := fs.Disk("local")
	//dirs := lc.Directories("c")
	//fmt.Println(dirs)

	dirs := fs.Directories("c/d")
	fmt.Println(dirs)
}
