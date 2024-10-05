package commands

import (
	"fmt"
	"os"

	"ggit/commands/internal"
)

func Add(repoDir, path string, index *internal.Index) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("adding file not found: %s", path)
	}

	object, err := internal.NewObject(internal.Blob, data)

	fmt.Println(object.Hash)
	return nil
}
