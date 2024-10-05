package internal

import (
	"bufio"
	"compress/zlib"
	"io"
	"os"
	"strconv"
	"strings"
	"crypto/sha1"
	"fmt"
	"encoding/hex"
)

func LoadObj(hash string) (*GitObject, error) {
	objf, err := os.Open(".git/objects/" + hash[:2] + "/" + hash[2:])
	if err != nil {
		return nil, err
	}
	defer objf.Close()

	zr, err := zlib.NewReader(objf)
	if err != nil {
		return nil, err
	}

	br := bufio.NewReader(zr)

	// format: (commit|tree|blob) <byte_size>\x00<body>
	tpe, err := br.ReadString(' ')
	if err != nil {
		return nil, err
	}
	tpe = strings.Trim(tpe, " ")
	objType, err := NewObjType(tpe)
	if err != nil {
		return nil, err
	}

	sizeStr, err := br.ReadString(0x00)
	if err != nil {
		return nil, err
	}
	sizeStr = strings.Trim(sizeStr, "\x00")
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, size)
	if _, err := io.ReadFull(br, buf); err != nil {
		return nil, err
	}

	return &GitObject{
		Type: objType,
		Hash: hash,
		Size: size,
		Data: buf,
	}, nil

}

func NewObject(tpe ObjType, data []byte) (*GitObject, error) {
	size := int64(len(data))
	sha1 := sha1.New()
	header := []byte(fmt.Sprintf("%s %d\x00", tpe, size))
	_, err := sha1.Write(append(header, data...))
	if err != nil {
		return nil, err
	}
	hash := hex.EncodeToString(sha1.Sum(nil))
	return &GitObject{
		Type: tpe,
		Hash: hash,
		Size: size,
		Data: data,
	}, nil
}

type ObjType string
func (t ObjType) String() string {
	return string(t)
}

const (
	Unknown ObjType = ""
	Commit          = "commit"
	Tree            = "tree"
	Blob            = "blob"
)

func NewObjType(tpe string) (ObjType, error) {
	switch tpe {
	case "commit":
		return Commit, nil
	case "tree":
		return Tree, nil
	case "blob":
		return Blob, nil
	default:
		return Unknown, fmt.Errorf("Unknown object type: %s", tpe)
	}
}

type GitObject struct {
	Type ObjType
	Hash string
	Size int64
	Data []byte
}
