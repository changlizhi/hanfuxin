package routers

import (
	"github.com/astaxie/beego"
	"hanfuxin/controllers"
)

func init() {
	beego.Router("/juesequanxian", &controllers.Juesequanxiancontroller{})
	beego.Router("/juesequanxian/:Id", &controllers.Juesequanxiancontroller{})

}
