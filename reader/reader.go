package reader

import (
	"os"
)

func Read_file(path string) []byte {

	f, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return f
}
