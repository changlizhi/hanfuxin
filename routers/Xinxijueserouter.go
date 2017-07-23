package routers

import (
	"github.com/astaxie/beego"
	"hanfuxin/controllers"
)

func init() {
	beego.Router("/xinxijuese", &controllers.Xinxijuesecontroller{})
	beego.Router("/xinxijuese/:Id", &controllers.Xinxijuesecontroller{})

}
