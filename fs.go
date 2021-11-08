// fs.go kee > 2021/10/31

package fs

type Drivers map[string]Filesystem

var disks Drivers

func Register(driver string, adapter Filesystem) Drivers {
	disks[driver] = adapter
	return disks
}

func Disk(driver string) Filesystem {
	if adapter, ok := disks[driver]; ok {
		return adapter
	}
	return nil
}

// Put | Write (filename, content, ...options)

// Get (filename)

// Exists | Has (filename)

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

// Directories(directory, deepLevel)

// Mkdir(path, permMode)

// MkdirMulti(paths []stirng, permMode)

// RmDir(path)

// RmDirMulti(paths []string)
