package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hanfuxin/appmodels"
	"hanfuxin/appinits"
	"hanfuxin/zdxinxiservices"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strconv"
	"strings"
)

type Xinxicontroller struct {
	beego.Controller
}

func (c *Xinxicontroller) Get() {
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
	xinxi := zdxinxiservices.Chaxunxinxi(id)
	c.Data[zf.Zfs.Json(true)] = xinxi
	c.ServeJSON()
	return
}
func (c *Xinxicontroller) Post() {
	xinxi := appmodels.Xinxi{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &xinxi)
	serviceret := zdxinxiservices.Tianjiaxinxi(&xinxi)
	tishi := appinits.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = appinits.Tishis[splitret[zfzhi.Zhi.Shuzi0zhi()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1zhi()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Xinxicontroller) Patch() {
	xinxi := appmodels.Xinxi{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &xinxi)
	serviceret := zdxinxiservices.Xiugaixinxi(&xinxi)
	tishi := appinits.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = appinits.Tishis[splitret[zfzhi.Zhi.Shuzi0zhi()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1zhi()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Xinxicontroller) Delete() {
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
	serviceret := zdxinxiservices.Shanchuxinxi(id)
	tishi := appinits.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = appinits.Tishis[splitret[zfzhi.Zhi.Shuzi0zhi()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1zhi()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
