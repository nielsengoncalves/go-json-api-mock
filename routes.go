package main

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
	"io/ioutil"
)

var gets = map[string]string{
	"/get": "get-response.json", // [GET] /mock/get
}

var posts = map[string]string{
	"/post": "post-response.json", // [POST] /mock/post
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadRoutes(app iris.Application) {
	mockRoutes := app.Party("/mock", logThisMiddleware)
	{
		loadGetRoutes(mockRoutes)
		loadPostRoutes(mockRoutes)
	}
}

func loadGetRoutes(app router.Party) {
	for route, file := range gets {
		app.Get(route, func(ctx context.Context) {
			ctx.JSON(getJsonFromFile(file))
		})
	}
}

func loadPostRoutes(app router.Party) {
	for route, file := range posts {
		app.Post(route, func(ctx context.Context) {
			ctx.JSON(getJsonFromFile(file))
		})
	}
}

func getJsonFromFile(response string) interface{} {
	file, err := ioutil.ReadFile("./responses/" + response)
	check(err)
	var f interface{}
	json.Unmarshal(file, &f)
	return f
}

func logThisMiddleware(ctx context.Context) {
	ctx.Application().Log("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())
	ctx.Next()
}
