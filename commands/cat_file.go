package commands

import (
	"archive/zip"
	"io"
)

// Provide contents or details of repository objects
func CatFile(hash string) {
	zf, err := zip.OpenReader(".git/objects/" + hash[:2] + "/" + hash[2:])
	if err != nil {
		panic(err)
	}
	defer zf.Close()

	bs, err := zf.File[0].Open()
	if err != nil {
		panic(err)
	}
	defer bs.Close()
	io.ReadAll(bs)
}
