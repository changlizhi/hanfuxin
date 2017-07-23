package routers

import (
	"github.com/astaxie/beego"
	"hanfuxin/controllers"
)

func init() {
	beego.Router("/quanxian", &controllers.Quanxiancontroller{})
	beego.Router("/quanxian/:Id", &controllers.Quanxiancontroller{})

}
