package routers

import (
	"github.com/astaxie/beego"
	"hanfuxin/controllers"
)

func init() {
	beego.Router("/juese", &controllers.Juesecontroller{})
}
