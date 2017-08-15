package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hanfuxin/appmodels"
	"hanfuxin/appinits"
	"hanfuxin/zdjueseservices"
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"strconv"
	"strings"
)

type Juesecontroller struct {
	beego.Controller
}

func (c *Juesecontroller) Get() {
	canshu := c.GetString(zfzhi.Zhi.Mh() + zf.Zfs.Id(false))
	id, err := strconv.Atoi(canshu)
	if err != nil {
		log.Println(err)
		c.Data[zf.Zfs.Json(true)] = map[string]string{
			zf.Zfs.Error05(false): appinits.Cuowus[zf.Zfs.Error05(false)].Zhi,
		}
		c.ServeJSON()
		return
	}
	juese := zdjueseservices.Chaxunjuese(id)
	c.Data[zf.Zfs.Json(true)] = juese
	c.ServeJSON()
	return
}
func (c *Juesecontroller) Post() {
	juese := appmodels.Juese{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &juese)
	serviceret := zdjueseservices.Tianjiajuese(&juese)
	tishi := appinits.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = appinits.Tishis[splitret[zfzhi.Zhi.Shuzi0()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Juesecontroller) Patch() {
	juese := appmodels.Juese{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &juese)
	serviceret := zdjueseservices.Xiugaijuese(&juese)
	tishi := appinits.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = appinits.Tishis[splitret[zfzhi.Zhi.Shuzi0()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Juesecontroller) Delete() {
	canshu := c.GetString(zfzhi.Zhi.Mh() + zf.Zfs.Id(false))
	id, err := strconv.Atoi(canshu)
	if err != nil {
		log.Println(err)
		c.Data[zf.Zfs.Json(true)] = map[string]string{
			zf.Zfs.Error05(false): appinits.Cuowus[zf.Zfs.Error05(false)].Zhi,
		}
		c.ServeJSON()
		return
	}
	serviceret := zdjueseservices.Shanchujuese(id)
	tishi := appinits.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = appinits.Tishis[splitret[zfzhi.Zhi.Shuzi0()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
