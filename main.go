package main

import (
	"flag"
	"fmt"
	"ggit/commands"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		cmd := flag.NewFlagSet("init", flag.ExitOnError)
		cmd.Parse(os.Args[2:])
		path := cmd.Arg(0)
		if path == "" {
			path = "."
		}
		commands.Init(path)
	case "hash-object":
		cmd := flag.NewFlagSet("hash-object", flag.ExitOnError)
		write := cmd.Bool("w", false, "Actually write the object into the database")
		objType := cmd.String("t", "blob", "Specify the type")
		cmd.Parse(os.Args[2:])
		file := cmd.Arg(0)
		bs, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		hash := commands.HashObject(bs, *objType, *write)
		fmt.Println(hash)
	case "cat-file":
		cmd := flag.NewFlagSet("cat-file", flag.ExitOnError)
		cmd.Parse(os.Args[2:])
		hash := cmd.Arg(0)
		commands.CatFile(hash)
	}

}
