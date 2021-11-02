// file_test.go kee > 2021/10/31

package storage

import (
	"fmt"
	"testing"
	"time"
)

func TestFile(t *testing.T) {
	fn := "tmp/a/b/../b/c/d.jpg"
	finfo, err := FileInfo(fn)

	if err == nil {
		fmt.Println("File path: ", finfo.Path)
		fmt.Println("Dirname: ", finfo.Dirname)
		fmt.Printf("MIMEType: %s, extension: %s\n", finfo.MIMEType, finfo.Extension)
		fmt.Println("File name: ", finfo.Filename)
		fmt.Printf("Size: %d bytes | %.2f Kbytes\n", finfo.Size, float64(finfo.Size/1000))
		fmt.Println("Direcory: ", finfo.Stat().IsDir())
		fmt.Printf("Permission: (9-bits) %s | (4-digits) %#o | (3-digits) %o\n", finfo.PermMode, finfo.PermMode, finfo.PermMode)
		fmt.Println("File Last Modified: ", finfo.ModTime.Format(time.RFC3339))
	} else {
		fmt.Println("Invalid file!")
	}
}
