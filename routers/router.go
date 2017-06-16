package routers

import (
	"bank/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.ListController{})
	beego.Router("/list/:bankId", &controllers.ListController{}, "get:ShowList")
}
