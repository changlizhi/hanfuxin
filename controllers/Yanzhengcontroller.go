package controllers
import(
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
type Yanzhengcontroller struct{
beego.Controller
}
func (c *Yanzhengcontroller)Get(){
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
yanzheng:=zdyanzhengservices.Chaxunyanzheng(id)
c.Data[zf.Json(true)]=yanzheng
c.ServeJSON()
return
}
func (c *Yanzhengcontroller)Post(){
zf:=zf.Zf{}
zfzhi:=zfzhi.Zfzhi{}
kzf:=zfzhi.Kongzifuzhi()
xhx:=zfzhi.Xiahuaxianzhi()
sz0:=zfzhi.Shuzi0zhi()
sz1:=zfzhi.Shuzi1zhi()
mh:=zfzhi.Maohaozhi()
yanzheng:=appmodels.Yanzheng{}
json.Unmarshal(c.Ctx.Input.RequestBody,&yanzheng)
serviceret:=zdyanzhengservices.Tianjiayanzheng(&yanzheng)
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
func (c *Yanzhengcontroller)Patch(){
zf:=zf.Zf{}
zfzhi:=zfzhi.Zfzhi{}
kzf:=zfzhi.Kongzifuzhi()
xhx:=zfzhi.Xiahuaxianzhi()
sz0:=zfzhi.Shuzi0zhi()
sz1:=zfzhi.Shuzi1zhi()
mh:=zfzhi.Maohaozhi()
yanzheng:=appmodels.Yanzheng{}
json.Unmarshal(c.Ctx.Input.RequestBody,&yanzheng)
serviceret:=zdyanzhengservices.Xiugaiyanzheng(&yanzheng)
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
func (c *Yanzhengcontroller)Delete(){
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
serviceret:=zdyanzhengservices.Shanchuyanzheng(id)
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
