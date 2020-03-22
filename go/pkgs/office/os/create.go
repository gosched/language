package explore

import (
	"os"
	"syscall"
)

// the default permissions value is 777 for a directory
// the default permissions value is 666 for a regular file
// the default umask       value is 022
// 777 - 022 == 755
// 666 - 022 == 644

// 0777 == drwxrwxrwx
// 0755 == drwxr-xr-x
// 0644 == -rw-r--r--

// CreateFile .
func CreateFile(name string, perm os.FileMode) {
	os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
}

// CreateFileSpecifyUmask .
func CreateFileSpecifyUmask() {

}

// CreateDir .
func CreateDir(name string, perm os.FileMode) {
	os.Mkdir(name, perm)
}

// CreateDirSpecifyUmask .
func CreateDirSpecifyUmask(name string, perm os.FileMode, newmask int) {
	oldmask := syscall.Umask(newmask)
	defer syscall.Umask(oldmask)
	os.Mkdir(name, perm)
}
