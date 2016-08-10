package main

import (
	"github.com/kataras/iris"
)

func main() {
	render := func(ctx *iris.Context) {
		ctx.Render("index.html", nil)
	}

	iris.Get("/", render)("home")
	iris.Get("/about", render)("about")
	iris.Get("/page/:id", render)("page")

	iris.Listen(":8080")
}
