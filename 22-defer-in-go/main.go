package main

import (
	"fmt"
)

type Engineer struct {
	Name string
}

func (e *Engineer) UpdateName(name string) {
	e.Name = name
}

func doStuff(e *Engineer) {
	name := "Ralph Emmason"
	defer e.UpdateName(name)
	fmt.Println("doing other exciting stuff")

	name = "Ralph O Emmason"
}

func main() {
	fmt.Println("defer in go")
	elliot := &Engineer{
		Name: "Ralph",
	}

	fmt.Printf("%+v\n", elliot)
	doStuff(elliot)
	fmt.Printf("%+v\n", elliot)
}
