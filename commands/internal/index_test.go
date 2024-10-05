package internal

import (
	"bytes"
	"testing"
)

func TestReadHeader(t *testing.T) {
	index := &Index{}
	input := []byte{
		0x44, 0x49, 0x52, 0x43, // DIRC
		0x00, 0x00, 0x00, 0x02, // 2
		0x00, 0x00, 0x00, 0x29, // 41
	}
	err := index.readHeader(bytes.NewReader(input))
	if err != nil {
		t.Error("Error loading index", err)
	}

	if string(index.Signature[:]) != "DIRC" {
		t.Error("Invalid signature")
	}

	if index.Version != 2 {
		t.Error("Invalid version")
	}

	if index.EntryNum != 41 {
		t.Error("Invalid entry number")
	}
}
