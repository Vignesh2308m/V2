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

func Word_count(text []byte, prev byte) int {
	var sum int
	for _, i := range text {
		if i == 32 && prev != i {
			sum++
		}
		prev = i
	}

	return sum
}
