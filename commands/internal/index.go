package internal

import (
	"os"
)

type IndexEntry struct {
	ctime_s int
	ctime_n int
	mtime_s int
	mtime_n int
	dev     int
	ino     int
	mode    int
	uid     int
	gid     int
	size    int
	sha1    string
	flags   int
	path    string
}

func ReadIndex() []IndexEntry {

	if !isExist(".git/index") {
		return []IndexEntry{}
	}

	// index, err := os.ReadFile(".git/index")
	// if err != nil {
	// 	panic(err)
	// }

	// digest := index[20:]

	return nil
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
