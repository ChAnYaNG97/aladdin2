package data

import (
	"fmt"
	"testing"
)

func TestReadDataFromJsonFile(t *testing.T) {
	tasks, nodes := ReadDataFromJsonFile()
	fmt.Println(tasks)
	fmt.Println(nodes)
}