package controllers

import "github.com/kataras/iris/mvc"

type HelloController struct{}

func (c *HelloController) Get() mvc.Result {
	return mvc.View{
		Name: "views/hello/index.html",
		Data: map[string]interface{}{
			"Title":     "Hello Page",
			"MyMessage": "Welcome to my awesome website",
		},
	}
}
