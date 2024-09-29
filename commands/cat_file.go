package commands

import (	
	"os"
	"ggit/commands/internal"
)

// Provide contents or details of repository objects
func CatFile(hash string) {
	obj, err := internal.LoadObj(hash)

	if err != nil {
		panic(err)
	}
	os.Stdout.Write(obj.Data)
}
