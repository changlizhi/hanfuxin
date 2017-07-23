package routers

import (
	"github.com/astaxie/beego"
	"hanfuxin/controllers"
)

func init() {
	beego.Router("/yanzhengleixing", &controllers.Yanzhengleixingcontroller{})
	beego.Router("/yanzhengleixing/:Id", &controllers.Yanzhengleixingcontroller{})

}
