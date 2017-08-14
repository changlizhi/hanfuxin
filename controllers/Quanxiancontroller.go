package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hanfuxin/appmodels"
	"hanfuxin/appinits"
	"hanfuxin/zdquanxianservices"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strconv"
	"strings"
)

type Quanxiancontroller struct {
	beego.Controller
}

func (c *Quanxiancontroller) Get() {
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
	quanxian := zdquanxianservices.Chaxunquanxian(id)
	c.Data[zf.Zfs.Json(true)] = quanxian
	c.ServeJSON()
	return
}
func (c *Quanxiancontroller) Post() {
	quanxian := appmodels.Quanxian{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &quanxian)
	serviceret := zdquanxianservices.Tianjiaquanxian(&quanxian)
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
func (c *Quanxiancontroller) Patch() {
	quanxian := appmodels.Quanxian{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &quanxian)
	serviceret := zdquanxianservices.Xiugaiquanxian(&quanxian)
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
func (c *Quanxiancontroller) Delete() {
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
	serviceret := zdquanxianservices.Shanchuquanxian(id)
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
