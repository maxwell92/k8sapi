package main

import (
	"encoding/xml"

	"github.com/kataras/iris"
)

type ExampleXml struct {
	XMLName xml.Name `xml:"example"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}

func Data(ctx *iris.Context) {
	ctx.Data(iris.StatusOK, []byte("Hello World"))
}

func Text(ctx *iris.Context) {
	ctx.Text(iris.StatusOK, "Hello World")
}

func Json(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, map[string]string{"hello": "world"})
}

func Jsonp(ctx *iris.Context) {
	ctx.JSONP(iris.StatusOK, "callbackName", map[string]string{"hello": "world"})
}

func Xml(ctx *iris.Context) {
	ctx.XML(iris.StatusOK, ExampleXml{One: "hello", Two: "world"})
}

func Markdown(ctx *iris.Context) {
	ctx.Markdown(iris.StatusOK, " ##Hello World ")
}

func main() {

	iris.Get("/data", Data)

	iris.Get("/text", Text)

	iris.Get("/json", Json)

	iris.Get("/jsonp", Jsonp)

	iris.Get("/xml", Xml)

	iris.Get("/markdown", Markdown)

	iris.Listen(":8080")
}
