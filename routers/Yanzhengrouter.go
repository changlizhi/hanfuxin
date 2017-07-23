package routers

import (
	"github.com/astaxie/beego"
	"hanfuxin/controllers"
)

func init() {
	beego.Router("/yanzheng", &controllers.Yanzhengcontroller{})
	beego.Router("/yanzheng/:Id", &controllers.Yanzhengcontroller{})

}
