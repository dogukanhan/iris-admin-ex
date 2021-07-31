package main

import (
	"app/controller"
	"app/environment"
	"app/service"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()

	app.RegisterView(iris.Blocks("./views", ".html"))

	app.HandleDir("/", iris.Dir("./assets"))

	env := environment.DEV

	if len(os.Args) > 1 && os.Args[1] == "prod"{
		env = environment.PROD
		print("Environment is Production")
	}else{
		print("Environment is Development")
	}

	mvc.Configure(app.Party("/"), func(app *mvc.Application) {
		app.Register(env,
			service.NewDiskSpaceService,
		)
		app.Handle(new(controller.HomeController))
	})

	app.Listen(":8080", iris.WithLogLevel("debug"))
}
