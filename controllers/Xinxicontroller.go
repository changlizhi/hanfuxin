package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
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
	xinxi := zdxinxiservices.Chaxunxinxi(id)
	c.Data[zf.Zfs.Json(true)] = xinxi
	c.ServeJSON()
	return
}
func (c *Xinxicontroller) Post() {
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	sz0 := zfzhi.Shuzi0zhi()
	sz1 := zfzhi.Shuzi1zhi()
	mh := zfzhi.Maohaozhi()
	xinxi := appmodels.Xinxi{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &xinxi)
	serviceret := zdxinxiservices.Tianjiaxinxi(&xinxi)
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
func (c *Xinxicontroller) Patch() {
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	sz0 := zfzhi.Shuzi0zhi()
	sz1 := zfzhi.Shuzi1zhi()
	mh := zfzhi.Maohaozhi()
	xinxi := appmodels.Xinxi{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &xinxi)
	serviceret := zdxinxiservices.Xiugaixinxi(&xinxi)
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
func (c *Xinxicontroller) Delete() {
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
	serviceret := zdxinxiservices.Shanchuxinxi(id)
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
