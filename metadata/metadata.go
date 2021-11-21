// metadata.go kee > 2021/10/31

package fs

import (
	"github.com/gabriel-vasile/mimetype"
	//	"os"
	//	"time"
)

type Metadata interface {
	MIMEType() (*mimetype.MIME, error)
	//	GetPath() string
	//	GetDirname() string
	//	GetExtension() string
	//	GetFilename() string
	//	GetModTime() time.Time
	//	GetSize() int64
	//	GetPermMode() os.FileMode
	//	GetType() string
}
