package controller

import (
	"app/service"
	"github.com/kataras/iris/v12"
)

type HomeController struct {
	Service service.DiskSpaceService
	ctx     *iris.Context
}

func (c *HomeController) Get(ctx iris.Context) {


	message, err := c.Service.GetDiskSpace()

	if err == nil{
		ctx.ViewLayout("main")

		ctx.ViewData("message", message)

		err = ctx.View("index")
	}else{
		print("Error:" + err.Error())
	}

}
