package main

import (
	"github.com/kataras/iris"
)

type myHandleGet struct {
}

func (m myHandlerGet) Serve(c *iris.Context) {
	c.Write("From %s", c.PathString())
}

func main() {
	iris.Handle("GET", "/get", myHandlerGet())
	iris.Handle("POST", "/post", post)
	iris.Handle("PUT", "/put", put)
	iris.Handle("DELETE", "/delete", del)
	iris.Listen(":8080")
}
