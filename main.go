package main

import (
	"app/controller"
	"app/environment"
	"app/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()

	app.RegisterView(iris.Blocks("./views", ".html"))

	app.HandleDir("/", iris.Dir("./assets"))

	mvc.Configure(app.Party("/"), func(app *mvc.Application) {
		app.Register(environment.DEV,
			service.NewDiskSpaceService,
		)
		app.Handle(new(controller.HomeController))
	})

	app.Listen(":8080", iris.WithLogLevel("debug"))
}

