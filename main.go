package main

import (
	"fmt"
	"strings"
	"time"
)

func letter_check(s string) bool {
	return strings.Contains(s, "V")
}

func concat(a string, b string) string {
	return a + b
}

func main() {

	start := time.Now()

	sample_data := map[int]string{
		1: "Vijayanila",
		2: "Vignesh",
	}

	data := NewDataSet(sample_data)

	data.apply(strings.ToUpper).filter(letter_check).reduce(concat)

	fmt.Println(data)
	fmt.Println(time.Since(start))
}
