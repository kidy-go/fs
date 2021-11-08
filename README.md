# kidy-go/fs

`fs`是包含在`kidy-go`项目中的一个文件系统中间件(Golan), 负责统一本地存储、云储存的接口组件

### 目前已完成:
- 基础文件接口定义
- 本地文件存储相关功能

### 接下来待完成:
- FTP/SFTP 协议的网络存储驱动
- Qiniu 云储协议驱动
- Aliyun OSS云存储协议驱动
- Amazon S3云存储协议驱动
- Dropbox 自建云存储协议驱动
- 其它协议持续扩展

## 接口定义:

```golang
type Filesystem interface {
	Write(filename string, content []byte, options ...interface{}) error

	Append(filename string, content []byte, options ...interface{}) error

	Move(src, dst string) error

	Copy(src, dst string) error

	Delete(src string) error

	Mkdir(folderPath string, perm os.FileMode) error

	Has(filename string) bool

	Get(filename string) ([]byte, error)

	Directories(path string) []string

	Files(path string) []string

	SetVisibility(filename string, visibility string) error

	GetVisibility(filename string) string

	Metadata(filename string) Metadata

	// TODO..
	// PutFileAs(filename string, fd os.File, dst string) error
}
```
