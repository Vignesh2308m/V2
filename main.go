package main

import (
	"fmt"

	"github.com/Vignesh2308m/V2/reader"
)

func main() {

	f := reader.Read_file("/home/vickynila/DevOps class")

	fmt.Println(reader.Word_count(f))
}
