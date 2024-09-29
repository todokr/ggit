package internal

import (
	"bufio"
	"compress/zlib"
	"io"
	"os"
	"strconv"
	"strings"
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
	objType := newObjType(tpe)

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
		Size: size,
		Data: buf,
	}, nil

}

type ObjType string

const (
	Unknown ObjType = ""
	Commit          = "commit"
	Tree            = "tree"
	Blob            = "blob"
)

func newObjType(tpe string) ObjType {
	switch tpe {
	case "commit":
		return Commit
	case "tree":
		return Tree
	case "blob":
		return Blob
	default:
		return Unknown
	}
}

type GitObject struct {
	Type ObjType
	Size int64
	Data []byte
}
