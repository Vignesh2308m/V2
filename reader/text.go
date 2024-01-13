package reader

import (
	"bufio"
	"os"
)

func Read_file(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)
	var val string
	for r.Scan() {
		val += r.Text()
	}
	return val, nil
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
