package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/zdjueseservices"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strconv"
	"strings"
)

type Juesecontroller struct {
	beego.Controller
}

func (c *Juesecontroller) Get() {
	zf := zf.Zf{}
	onetest := c.GetString(zf.Id(false))
	id, err := strconv.Atoi(onetest)
	if err != nil {
		log.Println(err)
		c.Data[zf.Json(true)] = map[string]string{zf.Error05(false): baseinits.Cuowus[zf.Error05(false)].Zhi}
		c.ServeJSON()
		return
	}
	juese := zdjueseservices.Chaxunjuese(id)
	c.Data[zf.Json(true)] = juese
	c.ServeJSON()
	return
}
func (c *Juesecontroller) Post() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	sz0 := zfzhi.Shuzi0zhi()
	sz1 := zfzhi.Shuzi1zhi()
	mh := zfzhi.Maohaozhi()

	juese := appmodels.Juese{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &juese)
	retbianma := zdjueseservices.Tianjiajuese(&juese)
	tishi := baseinits.Tishis[retbianma].Zhi
	if tishi == kzf {
		bms := strings.Split(retbianma, xhx)
		c.Data[zf.Json(true)] = baseinits.Tishis[bms[sz0]].Zhi + mh + bms[sz1]
		c.ServeJSON()
		return
	}
	log.Println(retbianma)
	c.Data[zf.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Juesecontroller) Delete() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	sz0 := zfzhi.Shuzi0zhi()
	sz1 := zfzhi.Shuzi1zhi()
	mh := zfzhi.Maohaozhi()

	onetest := c.GetString(zf.Id(false))
	id, err := strconv.Atoi(onetest)
	if err != nil {
		log.Println(err)
		c.Data[zf.Json(true)] = map[string]string{zf.Error05(false): baseinits.Cuowus[zf.Error05(false)].Zhi}
		c.ServeJSON()
		return
	}
	retbianma := zdjueseservices.Shanchujuese(id)
	tishi := baseinits.Tishis[retbianma].Zhi
	if tishi == kzf {
		bms := strings.Split(retbianma, xhx)
		c.Data[zf.Json(true)] = baseinits.Tishis[bms[sz0]].Zhi + mh + bms[sz1]
		c.ServeJSON()
		return
	}
	log.Println(retbianma)
	c.Data[zf.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Juesecontroller) Patch() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	sz0 := zfzhi.Shuzi0zhi()
	sz1 := zfzhi.Shuzi1zhi()
	mh := zfzhi.Maohaozhi()

	juese := appmodels.Juese{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &juese)
	retbianma := zdjueseservices.Xiugaijuese(&juese)
	tishi := baseinits.Tishis[retbianma].Zhi
	if tishi == kzf {
		bms := strings.Split(retbianma, xhx)
		c.Data[zf.Json(true)] = baseinits.Tishis[bms[sz0]].Zhi + mh + bms[sz1]
		c.ServeJSON()
		return
	}
	c.Data[zf.Json(true)] = tishi
	c.ServeJSON()
	return
}
