package read

import (
	"os"
)

func read_file(path string) []byte {

	f, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return f
}
