package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type Test struct {
	*iris.Context
}

// func (t *Test) Get() DOESN'T work
/*
func (t Test) Get() {
	fmt.Println("This is get")
	t.Write("this is get")
}
*/
// func (t *Test) Get() DOESN'T work
func (t Test) Get() {
	id := t.Param("id")
	fmt.Printf("%s\n", id)
	t.Write("Get from /api/%s/end", id)
}

func main() {
	t := new(Test)
	//iris.API("/api", Test{})
	iris.API("/api/:id/end", *t)
	iris.Listen(":10000")
}
