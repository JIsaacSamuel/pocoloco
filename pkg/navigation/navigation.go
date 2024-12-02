package navigation

import (
	"io/fs"
	"log"
	"os"
)

func Get_dirs() []fs.DirEntry {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	return files
}
