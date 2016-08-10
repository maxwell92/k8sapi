package main

import "github.com/kataras/iris"

// nothing to do, defaults to ./templates and .html extension, no need to import any template engine because HTML engine is the default
// if anything else has been registered
func main() {
	iris.Config.IsDevelopment = true // this will reload the templates on each request, defaults to false

	iris.Get("/hi", hi)
	iris.Listen(":8080")
}

func hi(ctx *iris.Context) {
	ctx.MustRender("hi.html", struct {
		Name string
		Age  int
	}{Name: "iris", Age: 24})
}
