package commands

import (
	"fmt"
	"os"
)

// Initialize git repository
func Init(repo string) {
	os.Mkdir(repo, 0755)
	os.Mkdir(repo+"/.git", 0755)
	os.Mkdir(repo+"/.git/objects", 0755)
	os.Mkdir(repo+"/.git/refs", 0755)
	os.Mkdir(repo+"/.git/refs/heads", 0755)

	f, err := os.Create(repo + "/.git/HEAD")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("ref: refs/heads/master\n")

	fmt.Println("Initialized empty Git repository in", repo+"/.git")
}
