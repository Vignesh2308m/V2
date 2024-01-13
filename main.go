package main

import (
	"fmt"

	"github.com/Vignesh2308m/V2/reader"
)

func main() {

	f, err := reader.Read_file("/home/vickynila/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(f)

	// fmt.Println(reader.Word_count(f, 0))

}
