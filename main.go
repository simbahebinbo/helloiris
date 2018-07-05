package main

import (
	"time"

	"github.com/chrisgoffinet/helloiris/web/controllers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

const (
	DefaultTitle  = "My Awesome Site"
	DefaultLayout = "layouts/main.html"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./templates", ".html").Layout(DefaultLayout).Reload(true))
	mvc.New(app.Party("/hello")).Handle(new(controllers.HelloController))

	app.Use(func(ctx iris.Context) {
		// set the title, current time and a layout in order to be used if and when the next handler(s) calls the .Render function
		ctx.ViewData("Title", DefaultTitle)
		now := time.Now().Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
		ctx.ViewData("CurrentTime", now)
		ctx.ViewLayout(DefaultLayout)

		ctx.Next()
	})
	app.Run(
		// Start the web server at localhost:3000
		iris.Addr("localhost:3000"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
