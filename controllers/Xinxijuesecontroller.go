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
	zfzhi := zfzhi.Zfzhi{}
	mh := zfzhi.Maohaozhi()
	canshu := c.GetString(mh + zf.zfs.Id(false))
	id, err := strconv.Atoi(canshu)
	if err != nil {
		log.Println(err)
		c.Data[zf.Zfs.Json(true)] = map[string]string{
			zf.Zfs.Error05(false): baseinits.Cuowus[zf.Error05(false)].Zhi,
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
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	sz0 := zfzhi.Shuzi0zhi()
	sz1 := zfzhi.Shuzi1zhi()
	mh := zfzhi.Maohaozhi()
	xinxijuese := appmodels.Xinxijuese{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &xinxijuese)
	serviceret := zdxinxijueseservices.Tianjiaxinxijuese(&xinxijuese)
	tishi := baseinits.Tishis[serviceret].Zhi
	if tishi == kzf {
		splitret := strings.Split(serviceret, xhx)
		c.Data[zf.Zfs.Json(true)] = baseinits.Tishis[splitret[sz0]].Zhi + mh + splitret[sz1]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Xinxijuesecontroller) Patch() {
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	sz0 := zfzhi.Shuzi0zhi()
	sz1 := zfzhi.Shuzi1zhi()
	mh := zfzhi.Maohaozhi()
	xinxijuese := appmodels.Xinxijuese{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &xinxijuese)
	serviceret := zdxinxijueseservices.Xiugaixinxijuese(&xinxijuese)
	tishi := baseinits.Tishis[serviceret].Zhi
	if tishi == kzf {
		splitret := strings.Split(serviceret, xhx)
		c.Data[zf.Zfs.Json(true)] = baseinits.Tishis[splitret[sz0]].Zhi + mh + splitret[sz1]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Xinxijuesecontroller) Delete() {
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	sz0 := zfzhi.Shuzi0zhi()
	sz1 := zfzhi.Shuzi1zhi()
	mh := zfzhi.Maohaozhi()
	canshu := c.GetString(mh + zf.Zfs.Id(false))
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
	if tishi == kzf {
		splitret := strings.Split(serviceret, xhx)
		c.Data[zf.Zfs.Json(true)] = baseinits.Tishis[splitret[sz0]].Zhi + mh + splitret[sz1]
		c.ServeJSON()
		return
	}
	c.Data[zf.Json(true)] = tishi
	c.ServeJSON()
	return
}
