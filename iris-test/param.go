package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type AbcController struct {
	*iris.Context
}

func (ac AbcController) Get() {
	id := ac.Param("id")
	fmt.Printf("Get from /abc/%s\n", id)
	ac.Write("Get from /abc/" + id)
}

func main() {
	abc := new(AbcController)
	iris.API("/abc/:id", *abc)
	iris.Listen(":10000")
}
