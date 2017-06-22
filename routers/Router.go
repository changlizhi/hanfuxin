package routers

import (
	"github.com/astaxie/beego"
	"hanfuxin/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
