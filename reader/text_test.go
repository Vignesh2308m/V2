package reader

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	path := "/home/vickynila/test.txt"

	data, err := Read_file(path)

	assert.Nil(t, err)
	fmt.Println(data)
}
