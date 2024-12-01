package navigation

import (
	"log"
	"os"
)

func Get_dirs() []string {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var res []string
	res = append(res, "Back \n\n")

	for _, v := range files {
		res = append(res, v.Name())
	}

	return res
}
