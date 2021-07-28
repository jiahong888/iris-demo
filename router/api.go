package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-demo/controller/portal"
	"iris-demo/middleware"
)

func InitApp(app *iris.Application)  {
	api := app.Party("/jh", middleware.SignatureAuth)
	mvc.New(api.Party("/demo")).Handle(portal.NewDemoController())
}
