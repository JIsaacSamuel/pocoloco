package helpers

import (
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
