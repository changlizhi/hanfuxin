package controllers
import(
"encoding/json"
"github.com/astaxie/beego"
"hanfuxin/appmodels"
"hanfuxin/baseinits"
"hanfuxin/zdjuesequanxianservices"
"hanfuxin/zf"
"hanfuxin/zfzhi"
"log"
"strconv"
"strings"

)
type Juesequanxiancontroller struct{
beego.Controller
}
func (c *Juesequanxiancontroller)Get(){
zf:=zf.Zf{}
zfzhi:=zfzhi.Zfzhi{}
mh:=zfzhi.Maohaozhi()
canshu:=c.GetString(mh+zf.Id(false))
id,err:=strconv.Atoi(canshu)
if err!=nil{
log.Println(err)
c.Data[zf.Json(true)]=map[string]string{
zf.Error05(false):baseinits.Cuowus[zf.Error05(false)].Zhi,

}
c.ServeJSON()
return
}
juesequanxian:=zdjuesequanxianservices.Chaxunjuesequanxian(id)
c.Data[zf.Json(true)]=juesequanxian
c.ServeJSON()
return
}
func (c *Juesequanxiancontroller)Post(){
zf:=zf.Zf{}
zfzhi:=zfzhi.Zfzhi{}
kzf:=zfzhi.Kongzifuzhi()
xhx:=zfzhi.Xiahuaxianzhi()
sz0:=zfzhi.Shuzi0zhi()
sz1:=zfzhi.Shuzi1zhi()
mh:=zfzhi.Maohaozhi()
juesequanxian:=appmodels.Juesequanxian{}
json.Unmarshal(c.Ctx.Input.RequestBody,&juesequanxian)
serviceret:=zdjuesequanxianservices.Tianjiajuesequanxian(&juesequanxian)
tishi:=baseinits.Tishis[serviceret].Zhi
if tishi==kzf{
splitret:=strings.Split(serviceret,xhx)
c.Data[zf.Json(true)]=baseinits.Tishis[splitret[sz0]].Zhi+mh+splitret[sz1]
c.ServeJSON()
return
}
c.Data[zf.Json(true)]=tishi
c.ServeJSON()
return
}
func (c *Juesequanxiancontroller)Patch(){
zf:=zf.Zf{}
zfzhi:=zfzhi.Zfzhi{}
kzf:=zfzhi.Kongzifuzhi()
xhx:=zfzhi.Xiahuaxianzhi()
sz0:=zfzhi.Shuzi0zhi()
sz1:=zfzhi.Shuzi1zhi()
mh:=zfzhi.Maohaozhi()
juesequanxian:=appmodels.Juesequanxian{}
json.Unmarshal(c.Ctx.Input.RequestBody,&juesequanxian)
serviceret:=zdjuesequanxianservices.Xiugaijuesequanxian(&juesequanxian)
tishi:=baseinits.Tishis[serviceret].Zhi
if tishi==kzf{
splitret:=strings.Split(serviceret,xhx)
c.Data[zf.Json(true)]=baseinits.Tishis[splitret[sz0]].Zhi+mh+splitret[sz1]
c.ServeJSON()
return
}
c.Data[zf.Json(true)]=tishi
c.ServeJSON()
return
}
func (c *Juesequanxiancontroller)Delete(){
zf:=zf.Zf{}
zfzhi:=zfzhi.Zfzhi{}
kzf:=zfzhi.Kongzifuzhi()
xhx:=zfzhi.Xiahuaxianzhi()
sz0:=zfzhi.Shuzi0zhi()
sz1:=zfzhi.Shuzi1zhi()
mh:=zfzhi.Maohaozhi()
canshu:=c.GetString(mh+zf.Id(false))
id,err:=strconv.Atoi(canshu)
if err!=nil{
log.Println(err)
c.Data[zf.Json(true)]=map[string]string{
zf.Error05(false):baseinits.Cuowus[zf.Error05(false)].Zhi,

}
c.ServeJSON()
return
}
serviceret:=zdjuesequanxianservices.Shanchujuesequanxian(id)
tishi:=baseinits.Tishis[serviceret].Zhi
if tishi==kzf{
splitret:=strings.Split(serviceret,xhx)
c.Data[zf.Json(true)]=baseinits.Tishis[splitret[sz0]].Zhi+mh+splitret[sz1]
c.ServeJSON()
return
}
c.Data[zf.Json(true)]=tishi
c.ServeJSON()
return
}
