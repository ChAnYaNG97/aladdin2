package data

import (
	"bufio"
	"fmt"
	"io"
	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
	"os"
)

type Task struct {
	Cpu int `json:"cpu"`
	Mem int `json:"memory"`

}


type Node struct {
	Cpu int `json:"cpu"`
	Mem int `json:"memory"`
}



func ReadDataFromJsonFile() ([]Task, []Node) {
	tasks := make([]Task, 0)
	nodes := make([]Node, 0)

	taskInfoFile, err := os.Open("/Users/yangchen/Desktop/ISCAS/graduation/tianchi/newinstanceinfo.json")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, nil
	}

	br := bufio.NewReader(taskInfoFile)

	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break;
		}
		var task Task
		json.Unmarshal(line, &task)
		if task.Cpu != 0 {

			tasks = append(tasks, task)
		}

	}


	nodeInfoFile, err := os.Open("/Users/yangchen/Desktop/ISCAS/graduation/tianchi/newnodeinfo.json")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, nil
	}

	br = bufio.NewReader(nodeInfoFile)

	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break;
		}
		var node Node
		json.Unmarshal(line, &node)
		nodes = append(nodes, node)
	}
	fmt.Println(len(tasks))
	fmt.Println(len(nodes))


	return tasks, nodes



}