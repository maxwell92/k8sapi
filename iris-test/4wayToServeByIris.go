package main

import (
	"fmt"
	iris "github.com/kataras/iris"
)

// Way1: iris.Get(Route, Handler)
func Index(ctx *iris.Context) {
	ctx.Write("Hello World")
}

// Way2: iris.HandleFunc(Method, Routes, HandleFunc)
func Handle(ctx *iris.Context) {
	fmt.Println("This is handle")
}

// Way3: iris.API(Routes, API)

type UserAPI struct {
	*iris.Context
}

func (u UserAPI) Get() {
	u.Write("Get from /Users")
}

func (u UserAPI) GetBy(id string) {
	u.Write("Get from /Users/%s", id)
}

// Way 4: iris.Handle(Method, Route, iris.Handler)
type myHandler struct {
	data string
}

func (m *myHandler) Serve(ctx *iris.Context) { // Realise the iris.Handle.Serve() Method
	fmt.Println(m.data)
}

func main() {
	iris.Get("/", Index)
	iris.HandleFunc("GET", "/handle", Handle)
	iris.API("/users", UserAPI{})
	iris.Handle("GET", "/magic", &myHandler{"hello mush"})
	iris.Listen(":8080")
}
