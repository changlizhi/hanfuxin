package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/zdyanzhengservices"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strconv"
	"strings"
)

type Yanzhengcontroller struct {
	beego.Controller
}

func (c *Yanzhengcontroller) Get() {
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
	yanzheng := zdyanzhengservices.Chaxunyanzheng(id)
	c.Data[zf.Zfs.Json(true)] = yanzheng
	c.ServeJSON()
	return
}
func (c *Yanzhengcontroller) Post() {
	yanzheng := appmodels.Yanzheng{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &yanzheng)
	serviceret := zdyanzhengservices.Tianjiayanzheng(&yanzheng)
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
func (c *Yanzhengcontroller) Patch() {
	yanzheng := appmodels.Yanzheng{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &yanzheng)
	serviceret := zdyanzhengservices.Xiugaiyanzheng(&yanzheng)
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
func (c *Yanzhengcontroller) Delete() {
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
	serviceret := zdyanzhengservices.Shanchuyanzheng(id)
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
