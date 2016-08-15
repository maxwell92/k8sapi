package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func Id(ctx *iris.Context) {
	id := ctx.Param(id)
	fmt.Printf("%s\n", id)
}

func main() {
	iris.Get("/api/:id", Id)
	iris.Listen(":10000")
}
