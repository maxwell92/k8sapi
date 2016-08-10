package main

import (
	"github.com/kataras/iris"
)

func Index(ctx *iris.Context) {
	ctx.SetCookieKV("name", "hello")
	ctx.Write("cookie has been written")
}

func main() {
	iris.Get("/", Index)
	iris.Listen(":8080")
}
