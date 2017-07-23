package controllers
import(
"encoding/json"
"github.com/astaxie/beego"
"hanfuxin/appmodels"
"hanfuxin/baseinits"
"hanfuxin/zdjueseservices"
"hanfuxin/zf"
"hanfuxin/zfzhi"
"log"
"strconv"
"strings"

)
type Juesecontroller struct{
beego.Controller
}
func (c *Juesecontroller)Get(){
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
juese:=zdjueseservices.Chaxunjuese(id)
c.Data[zf.Json(true)]=juese
c.ServeJSON()
return
}
func (c *Juesecontroller)Post(){
zf:=zf.Zf{}
zfzhi:=zfzhi.Zfzhi{}
kzf:=zfzhi.Kongzifuzhi()
xhx:=zfzhi.Xiahuaxianzhi()
sz0:=zfzhi.Shuzi0zhi()
sz1:=zfzhi.Shuzi1zhi()
mh:=zfzhi.Maohaozhi()
juese:=appmodels.Juese{}
json.Unmarshal(c.Ctx.Input.RequestBody,&juese)
serviceret:=zdjueseservices.Tianjiajuese(&juese)
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
func (c *Juesecontroller)Patch(){
zf:=zf.Zf{}
zfzhi:=zfzhi.Zfzhi{}
kzf:=zfzhi.Kongzifuzhi()
xhx:=zfzhi.Xiahuaxianzhi()
sz0:=zfzhi.Shuzi0zhi()
sz1:=zfzhi.Shuzi1zhi()
mh:=zfzhi.Maohaozhi()
juese:=appmodels.Juese{}
json.Unmarshal(c.Ctx.Input.RequestBody,&juese)
serviceret:=zdjueseservices.Xiugaijuese(&juese)
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
func (c *Juesecontroller)Delete(){
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
serviceret:=zdjueseservices.Shanchujuese(id)
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
