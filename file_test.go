// file_test.go kee > 2021/10/31

package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

func dd(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(1)
}

func TestFile(t *testing.T) {
	//fn := "imgs/3/2/1.jpg"
	fn := "text/9/5/2/7.txt"
	op := os.O_WRONLY

	local := NewLocal(".temp")
	//c, _ := local.Get("tmp/v.jpg")
	s := []byte("Hello world (ccc|13|" + strconv.Itoa(op) + ")")
	local.Write(fn, s, op, "private")
	local.Append(fn, []byte(time.Now().Format("\n2006-01-02 15:04:05 .00000000")))
	local.Append(fn, []byte(time.Now().Format("\n2006-01-02 15:04:05 .0000000")))
	local.Append(fn, []byte(time.Now().Format("\n2006-01-02 15:04:05 .0000")))
	local.Write(fn, append([]byte("\n"), s...), os.O_APPEND)
	c, _ := local.Get(fn)
	fmt.Println(string(c))

	fmt.Println("===============================================================")

	local.SetVisibility(fn, "public")
	m := local.Metadata(fn)
	b, _ := json.Marshal(m)
	perm := fmt.Sprintf("%#o", m.PermMode)
	perm = perm[len(perm)-4:]
	fmt.Printf("METADATA: %s, Perm: %s, %s \n", b, m.PermMode, perm)
	fmt.Printf("Visibility: %s -> %s \n", fn, local.GetVisibility(fn))

	fmt.Println("===============================================================")

	var dir = "text"
	var dirs = local.Directories(dir)
	fmt.Println("fetch directories "+dir, dirs)
	var fils = local.Files(dir)
	fmt.Println("fetch files "+dir, fils)

	// fn := ".temp/tmp/a/b/../b/c/d.jpg"
	// finfo, err := FileInfo(fn)

	// if err == nil {
	// 	fmt.Println("File path: ", finfo.Path)
	// 	fmt.Println("Dirname: ", finfo.Dirname)
	// 	fmt.Printf("MIMEType: %s, extension: %s\n", finfo.MIMEType, finfo.Extension)
	// 	fmt.Println("File name: ", finfo.Filename)
	// 	fmt.Printf("Size: %d bytes | %.2f Kbytes\n", finfo.Size, float64(finfo.Size/1000))
	// 	fmt.Println("Direcory: ", finfo.Stat().IsDir())
	// 	fmt.Printf("Permission: (9-bits) %s | (4-digits) %#o | (3-digits) %o\n", finfo.PermMode, finfo.PermMode, finfo.PermMode)
	// 	fmt.Println("File Last Modified: ", finfo.ModTime.Format(time.RFC3339))
	// } else {
	// 	fmt.Println("Invalid file!")
	// }
}
