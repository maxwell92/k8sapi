package main

import (
	iris "github.com/kataras/iris"
)

func Index(ctx *iris.Context) {
	ctx.Write("Hello World")
}

func main() {
	iris.Get("/", Index)
	iris.Listen(":8080")
}
