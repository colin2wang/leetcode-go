package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var (
		number int
		name   string
	)
	fmt.Print("Input problem number: ")
	fmt.Scanf("%d\n", &number)
	fmt.Print("Input problem name: ")
	fmt.Scanf("%v\n", &name)

	fileName := fmt.Sprintf("./problems/%04d-%v_test.go", number, name)
	content, err := ioutil.ReadFile("./tools/problem-template.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(content))
	ioutil.WriteFile(fileName, content, 0777)
}
