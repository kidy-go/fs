// storage.go kee > 2021/10/31

package storage

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
