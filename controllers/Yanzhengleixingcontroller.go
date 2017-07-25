package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/zdyanzhengleixingservices"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strconv"
	"strings"
)

type Yanzhengleixingcontroller struct {
	beego.Controller
}

func (c *Yanzhengleixingcontroller) Get() {
	canshu := c.GetString(zfzhi.Zhi.Mh() + zf.Zfs.Id(false))
	id, err := strconv.Atoi(canshu)
	if err != nil {
		log.Println(err)
		c.Data[zf.Zfs.Json(true)] = map[string]string{
			zf.Zfs.Error05(false): baseinits.Cuowus[zf.Zfs.Error05(false)].Zhi,
		}
		c.ServeJSON()
		return
	}
	yanzhengleixing := zdyanzhengleixingservices.Chaxunyanzhengleixing(id)
	c.Data[zf.Zfs.Json(true)] = yanzhengleixing
	c.ServeJSON()
	return
}
func (c *Yanzhengleixingcontroller) Post() {
	yanzhengleixing := appmodels.Yanzhengleixing{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &yanzhengleixing)
	serviceret := zdyanzhengleixingservices.Tianjiayanzhengleixing(&yanzhengleixing)
	tishi := baseinits.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = baseinits.Tishis[splitret[zfzhi.Zhi.Shuzi0zhi()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1zhi()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Yanzhengleixingcontroller) Patch() {
	yanzhengleixing := appmodels.Yanzhengleixing{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &yanzhengleixing)
	serviceret := zdyanzhengleixingservices.Xiugaiyanzhengleixing(&yanzhengleixing)
	tishi := baseinits.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = baseinits.Tishis[splitret[zfzhi.Zhi.Shuzi0zhi()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1zhi()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Yanzhengleixingcontroller) Delete() {
	canshu := c.GetString(zfzhi.Zhi.Mh() + zf.Zfs.Id(false))
	id, err := strconv.Atoi(canshu)
	if err != nil {
		log.Println(err)
		c.Data[zf.Zfs.Json(true)] = map[string]string{
			zf.Zfs.Error05(false): baseinits.Cuowus[zf.Zfs.Error05(false)].Zhi,
		}
		c.ServeJSON()
		return
	}
	serviceret := zdyanzhengleixingservices.Shanchuyanzhengleixing(id)
	tishi := baseinits.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = baseinits.Tishis[splitret[zfzhi.Zhi.Shuzi0zhi()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1zhi()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
