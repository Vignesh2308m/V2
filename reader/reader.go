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

func Word_count(text []byte) int {
	var prev int
	var sum int
	for i := range text {
		if i == 32 ||
			(prev >= 41 && prev <= 90) ||
			(prev >= 97 && prev <= 121) ||
			(prev >= 30 && prev <= 39) {
			sum++
		}
		prev = i
	}

	return sum
}
