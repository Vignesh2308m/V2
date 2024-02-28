package main

import (
	"fmt"
	"strings"
)

func main() {

	sample_data := map[int]string{
		1: "Vijayanila",
		2: "Vignesh",
	}

	data := NewDataSet(sample_data)

	data.apply(strings.ToUpper)

	fmt.Println(data)
}
