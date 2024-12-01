package helpers

import (
	"log"
	"os"
)

func Go_to(dir_name string) error {
	err := os.Chdir(dir_name)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
