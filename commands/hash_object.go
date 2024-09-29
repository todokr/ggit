package commands

import (
	"crypto/sha1"
	"encoding/hex"
	"compress/zlib"
	"fmt"
	"os"
)

// Compute hash of object data of given type
// Return SHA-1 object hash as string
func HashObject(data []byte, objType string, write bool) string {
	header := objType + " " + fmt.Sprint(len(data))
	fullData := append([]byte(header+"\x00"), data...)

	h := sha1.New()
	h.Write(fullData)
	hash := hex.EncodeToString(h.Sum(nil))

	if write {
		dirPath := ".git/objects/" + string(hash[:2])
		os.MkdirAll(dirPath, 0755)
		filePath := dirPath + "/" + string(hash[2:])
		f, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		zw := zlib.NewWriter(f)
		defer zw.Close()
		zw.Write(fullData)
	}
	return hash
}
