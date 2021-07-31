package controller

import (
	"app/service"
	"github.com/kataras/iris/v12"
	"net/http"
)

type HomeController struct {
	Service service.DiskSpaceService
	ctx     *iris.Context
}

func (c *HomeController) Get(ctx iris.Context) {

	if c.Service  == nil {
		print("Service is nil")
	}
	if c.ctx == nil{
		print("Ctx is null")
	}

	message, err := c.Service.GetDiskSpace()

	ctx.ViewLayout("main")

	ctx.ViewData("message", message)

	err = ctx.View("index")
	if err != nil {
		print("Error:" + err.Error())
		http.Error(ctx.ResponseWriter(), err.Error(), http.StatusInternalServerError)
	}

}
