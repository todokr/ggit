package internal

import (
	"os"
	"path/filepath"
	"io"
	"encoding/binary"
	"fmt"
)

type Index struct {
	Header
	Entries []*IndexEntry
}

type Header struct {
	Signature  [4]byte
	Version   uint32
	EntryNum uint32
}

type IndexEntry struct {
	Hash       string
	NameLength uint16
	Path       []byte
}

func NewIndex(repoDir string) (*Index, error) {
	index := &Index{}
	indexPath := filepath.Join(repoDir, ".git/index")
	if _, err := os.Stat(indexPath); !os.IsNotExist(err) {
		if err := index.read(indexPath); err != nil {
			return nil, fmt.Errorf("failed to read index: %w", err)
		}
	}
	/////////////////////////////////////////
	// file, err := os.Open(path)	       //
	// if err != nil {		       //
	// 	return nil, err		       //
	// }				       //
	// defer file.Close()		       //
	// 				       //
	// err = index.ReadHeader(file)	       //
	// if err != nil {		       //
	// 	return nil, err		       //
	// }				       //
	/////////////////////////////////////////
	return index, nil
}

func (index *Index) read(indexPath string) error {
	f, err := os.Open(indexPath)
	if err != nil {
		return fmt.Errorf("failed to open index file: %w", err)
	}
	defer f.Close()

	err = index.readHeader(f)
	if err != nil {
		return fmt.Errorf("failed to read index header: %w", err)
	}
	return nil
}

func (index *Index) readHeader(r io.Reader) error {
	err := binary.Read(r, binary.BigEndian, &index.Header)
	if err != nil {
		return err
	}
	return nil
}
	
func NewIndexEntry(hash string, path []byte) *IndexEntry {
	return &IndexEntry{
		Hash:       hash,
		NameLength: uint16(len(path)),
		Path:       path,
	}

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

type DiffType int
const (
	Deleted DiffType = iota
	New
	Modified
)
func (diff DiffType) String() string {
	switch diff {
	case Deleted:
		return "deleted"
	case New:
		return "new file"
	case Modified:
		return "modified"
	default:
		return ""
	}
}

type DiffEntry struct {
	DiffType
	Entry   *IndexEntry
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
