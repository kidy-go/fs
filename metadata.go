// metadata.go kee > 2021/10/31

package storage

import (
	"github.com/gabriel-vasile/mimetype"
	"os"
	"path/filepath"
	"time"
)

type Metadata struct {
	Path      string      `json:"path"`
	Dirname   string      `json:"dirname"`
	Extension string      `json:"extension"`
	Filename  string      `json:"filename"`
	ModTime   time.Time   `json:"modified_at"`
	Size      int64       `json:"size"`
	PermMode  os.FileMode `json:"perm_mode"`
	Type      string      `json:"type"`
	Stat      os.FileInfo `json:"-"`
}

func FileInfo(filename string) (Metadata, error) {
	finfo, err := os.Stat(filename)
	if err != nil {
		return Metadata{}, err
	}

	fType := "file"
	if finfo.IsDir() {
		fType = "dir"
	}

	return Metadata{
		Path:      filepath.Clean(filename),
		Dirname:   filepath.Dir(filename),
		Extension: filepath.Ext(filename),
		Filename:  finfo.Name(),
		ModTime:   finfo.ModTime(),
		Size:      finfo.Size(),
		PermMode:  finfo.Mode(),
		Type:      fType,
		Stat:      finfo,
	}, nil
}

func (m Metadata) MIMEType() (*mimetype.MIME, error) {
	return mimetype.DetectFile(m.Path)
}
