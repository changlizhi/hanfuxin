package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hanfuxin/appmodels"
	"hanfuxin/jueseservices"
	"log"
)

type Juesecontroller struct {
	beego.Controller
}

func (c *Juesecontroller) Get() {
	c.TplName = "index.tpl"
}
func (c *Juesecontroller) Post() {
	juese := appmodels.Juese{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &juese)
	log.Println("juese----------", juese)
	jueseservices.Tianjiajuese(&juese)
	c.Data["json"] = juese
	c.ServeJSON()
	return
}
