// metadata.go kee > 2021/10/31

package storage

import (
	"github.com/gabriel-vasile/mimetype"
	"os"
	"path/filepath"
	"time"
)

type MIMEType struct {
	mime *mimetype.MIME
}

func DetectFile(filename string) (MIMEType, error) {
	mime, err := mimetype.DetectFile(filename)
	return MIMEType{mime}, err
}

func (m MIMEType) String() string {
	return m.mime.String()
}

func (m MIMEType) Extension() string {
	return m.mime.Extension()
}

func (m MIMEType) GetMIMEType() *mimetype.MIME {
	return m.mime
}

type Metadata struct {
	stat      os.FileInfo
	Path      string      `json:"path"`
	Dirname   string      `json:"dirname"`
	Extension string      `json:"extension"`
	Filename  string      `json:"filename"`
	ModTime   time.Time   `json:"modified_at"`
	Size      int64       `json:"size"`
	MIMEType  string      `json:"mimetype"`
	PermMode  os.FileMode `json:"perm_mode"`
}

func FileInfo(filename string) (Metadata, error) {
	finfo, err := os.Stat(filename)
	if err != nil {
		return Metadata{}, err
	}
	mime, _ := DetectFile(filename)

	return Metadata{
		Path:      filepath.Clean(filename),
		Dirname:   filepath.Dir(filename),
		Extension: mime.Extension(),
		Filename:  finfo.Name(),
		ModTime:   finfo.ModTime(),
		Size:      finfo.Size(),
		MIMEType:  mime.String(),
		PermMode:  finfo.Mode(),
		stat:      finfo,
	}, nil
}

func (m Metadata) Stat() os.FileInfo {
	return m.stat
}
