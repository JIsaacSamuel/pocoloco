package helpers

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
)

func Go_to(dir_name string) error {
	err := os.Chdir(dir_name)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func Start_coding() {
	cmd := exec.Command("code", ".")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func Open_nano(file_name string) {
	cmd := exec.Command("nano", file_name)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func Filer_files(list []fs.DirEntry, param string) []fs.DirEntry {
	var result []fs.DirEntry
	for i := 0; i < len(list); i++ {
		if param == list[i].Name()[:min(len(param), len(list[i].Name()))] {
			result = append(result, list[i])
		}
	}

	return result
}
