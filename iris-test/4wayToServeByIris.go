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
	u.Write("Get from /users")
}

func (u UserAPI) GetBy(id string) {
	u.Write("Get from /users/%s", id)
}

type TempController struct {
	*iris.Context
}

func (t TempController) Get() {
	//id := u.Param("id")
	t.Write("Get from /temp")
	fmt.Println("this is Get()")
}

func (t TempController) GetByName(id string) {
	t.Write("Get from /temp/%s", id)
	fmt.Println("this is GetBy")
}

func (t TempController) GetById() {
	id := t.Param("id")
	t.Write("Get from /temp/%s", id)
	fmt.Println("this is GetById")
}

type AbcController struct {
	*iris.Context
}

func (t AbcController) GetBy() {
	id := t.Param("id")
	t.Write("Get from /abc/%s", id)
	fmt.Println("this is GetBy")
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
	iris.API("/temp", TempController{})
	iris.API("/tmp", TempController{})
	iris.API("/abc/:id", AbcController{})
	iris.Handle("GET", "/magic", &myHandler{"hello mush"})
	iris.Listen(":8080")
}
