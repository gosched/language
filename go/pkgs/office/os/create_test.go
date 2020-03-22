package explore

import (
	"testing"
)

func TestCreateFile(t *testing.T) {
	CreateFile("file1.txt", 0777) // 0777 >  0755, 0777 - 0022 = 0755
	CreateFile("file2.txt", 0644) // 0644 <= 0644, 0644 - 0000 = 0644
}

func TestCreateDir(t *testing.T) {
	CreateDir("dir1", 0777) // 0777 >  0755, 0777 - 0022 = 0755
	CreateDir("dir2", 0755) // 0755 <= 0755, 0755 - 0000 = 0755
}

func TestCreateDirSpecifyUmask(t *testing.T) {
	CreateDirSpecifyUmask("dir3", 0777, 0) // 0777 - 0000 = 0777
	CreateDirSpecifyUmask("dir4", 0755, 0) // 0755 - 0000 = 0755
}
