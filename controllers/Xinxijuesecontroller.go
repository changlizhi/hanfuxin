package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/zdxinxijueseservices"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strconv"
	"strings"
)

type Xinxijuesecontroller struct {
	beego.Controller
}

func (c *Xinxijuesecontroller) Get() {
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
	xinxijuese := zdxinxijueseservices.Chaxunxinxijuese(id)
	c.Data[zf.Zfs.Json(true)] = xinxijuese
	c.ServeJSON()
	return
}
func (c *Xinxijuesecontroller) Post() {
	xinxijuese := appmodels.Xinxijuese{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &xinxijuese)
	serviceret := zdxinxijueseservices.Tianjiaxinxijuese(&xinxijuese)
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
func (c *Xinxijuesecontroller) Patch() {
	xinxijuese := appmodels.Xinxijuese{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &xinxijuese)
	serviceret := zdxinxijueseservices.Xiugaixinxijuese(&xinxijuese)
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
func (c *Xinxijuesecontroller) Delete() {
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
	serviceret := zdxinxijueseservices.Shanchuxinxijuese(id)
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
