package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	render := func(ctx *iris.Context) {
		ctx.Render("/index.html", nil)
		fmt.Println("render")
	}

	iris.Get("/", render)
	iris.Listen(":10000")
}
