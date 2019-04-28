package router

import (
	"../controller/IndexController"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
}
