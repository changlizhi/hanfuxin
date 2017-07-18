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
	zf := zf.Zf{}
	canshu := c.GetString(zf.Id(false))
	id, err := strconv.Atoi(canshu)
	if err != nil {
		log.Println(err)
		c.Data[zf.Json(true)] = map[string]string{
			zf.Error05(false): baseinits.Cuowus[zf.Error05(false)].Zhi,
		}
		c.ServeJSON()
		return
	}
	yanzhengleixing := zdyanzhengleixingservices.Chaxunyanzhengleixing(id)
	c.Data[zf.Json(true)] = yanzhengleixing
	c.ServeJSON()
	return
}
func (c *Yanzhengleixingcontroller) Post() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	sz0 := zfzhi.Shuzi0zhi()
	sz1 := zfzhi.Shuzi1zhi()
	mh := zfzhi.Maohaozhi()
	yanzhengleixing := appmodels.Yanzhengleixing{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &yanzhengleixing)
	serviceret := zdyanzhengleixingservices.Tianjiayanzhengleixing(&yanzhengleixing)
	tishi := baseinits.Tishis[serviceret].Zhi
	if tishi == kzf {
		splitret := strings.Split(serviceret, xhx)
		c.Data[zf.Json(true)] = baseinits.Tishis[splitret[sz0]].Zhi + mh + splitret[sz1]
		c.ServeJSON()
		return
	}
	c.Data[zf.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Yanzhengleixingcontroller) Patch() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	sz0 := zfzhi.Shuzi0zhi()
	sz1 := zfzhi.Shuzi1zhi()
	mh := zfzhi.Maohaozhi()
	yanzhengleixing := appmodels.Yanzhengleixing{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &yanzhengleixing)
	serviceret := zdyanzhengleixingservices.Xiugaiyanzhengleixing(&yanzhengleixing)
	tishi := baseinits.Tishis[serviceret].Zhi
	if tishi == kzf {
		splitret := strings.Split(serviceret, xhx)
		c.Data[zf.Json(true)] = baseinits.Tishis[splitret[sz0]].Zhi + mh + splitret[sz1]
		c.ServeJSON()
		return
	}
	c.Data[zf.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Yanzhengleixingcontroller) Delete() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	sz0 := zfzhi.Shuzi0zhi()
	sz1 := zfzhi.Shuzi1zhi()
	mh := zfzhi.Maohaozhi()
	canshu := c.GetString(zf.Id(false))
	id, err := strconv.Atoi(canshu)
	if err != nil {
		log.Println(err)
		c.Data[zf.Json(true)] = map[string]string{
			zf.Error05(false): baseinits.Cuowus[zf.Error05(false)].Zhi,
		}
		c.ServeJSON()
		return
	}
	serviceret := zdyanzhengleixingservices.Shanchuyanzhengleixing(id)
	tishi := baseinits.Tishis[serviceret].Zhi
	if tishi == kzf {
		splitret := strings.Split(serviceret, xhx)
		c.Data[zf.Json(true)] = baseinits.Tishis[splitret[sz0]].Zhi + mh + splitret[sz1]
		c.ServeJSON()
		return
	}
	c.Data[zf.Json(true)] = tishi
	c.ServeJSON()
	return
}
