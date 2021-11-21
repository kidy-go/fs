// fs.go kee > 2021/10/31

package fs

type Drivers map[string]Filesystem

var disks = Drivers{}

func Register(driver string, adapter Filesystem) Drivers {
	disks[driver] = adapter

	if 1 == len(disks) {
		SetDefault(driver)
	}
	return disks
}

func Disk(driver string) Filesystem {
	if adapter, ok := disks[driver]; ok {
		return adapter
	}
	return nil
}

func GetDefault() Filesystem {
	adapter := Disk("default")
	if adapter == nil {
		panic("Invalid default driver.")
	}
	return adapter
}

func SetDefault(driver string) {
	adapter := Disk(driver)
	if adapter != nil && "default" != driver {
		Register("default", adapter)
	}
}

// Put | Write (filename, content, ...options)
func Write(filename string, contents []byte, options ...interface{}) error {
	return GetDefault().Write(filename, contents, options...)
}

// Get (filename)
func Get(filename string) ([]byte, error) {
	return GetDefault().Get(filename)
}

// Exists | Has (filename)
func Has(filename string) bool {
	return GetDefault().Has(filename)
}

// Download (filename, name, headers)

// Url(filename) | TempUrl(filename, expires)

// Size / LastModified (filename)

// Metadata | Stat (filename)

// PutFile(filename, os.file) | PutFileAs(filename, os.file, newname)

// Append (filename, content)

// Copy(src, dst)

// Move(src, dst)

// Delete | Remove (filename)

// MultiDelete (filename []string)

// GetVisibility(name) returns visibility == 'public' || 'private'

// SetVisibility(name, permMode)

// Files(directory, deepLevel)
func Files(directory string) []string {
	return GetDefault().Files(directory)
}

// Directories(directory, deepLevel)
func Directories(directory string) []string {
	return GetDefault().Directories(directory)
}

// Mkdir(path, permMode)

// MkdirMulti(paths []stirng, permMode)

// RmDir(path)

// RmDirMulti(paths []string)
