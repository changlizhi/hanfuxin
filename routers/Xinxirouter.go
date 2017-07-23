package routers

import (
	"github.com/astaxie/beego"
	"hanfuxin/controllers"
)

func init() {
	beego.Router("/xinxi", &controllers.Xinxicontroller{})
	beego.Router("/xinxi/:Id", &controllers.Xinxicontroller{})

}
