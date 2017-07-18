package tests
import(
"github.com/astaxie/beego/context"
"hanfuxin/controllers"
"hanfuxin/tests"
"hanfuxin/zf"
"hanfuxin/zfzhi"
"log"
"strconv"
"testing"

)
func juesequanxiancontroller()*controllers.Juesequanxiancontroller{
c:=controllers.Juesequanxiancontroller{}
c.Data=make(map[interface{}]interface{})
c.Ctx=context.NewContext()
c.Ctx.Input=context.NewInput()
c.Ctx.Output=context.NewOutput()
c.Ctx.Output.Context=context.NewContext()
c.Ctx.Output.Context.ResponseWriter=&context.Response{new(tests.Mywriter),true,200}
return &c
}
func TestPostjuesequanxian(t*testing.T){
zfzhi:=zfzhi.Zfzhi{}
c:=juesequanxiancontroller()
reqjson:=zfzhi.Postjuesequanxianzhi()
c.Ctx.Input.RequestBody=[]byte(reqjson)
c.Post()
log.Println(c.Data)

}
func TestPatchjuesequanxian(t*testing.T){
zfzhi:=zfzhi.Zfzhi{}
c:=juesequanxiancontroller()
reqjson:=zfzhi.Patchjuesequanxianzhi()
c.Ctx.Input.RequestBody=[]byte(reqjson)
c.Patch()
log.Println(c.Data)

}
func TestDeletejuesequanxian(t *testing.T){
zf:=zf.Zf{}
zfzhi:=zfzhi.Zfzhi{}
paramid:=strconv.Itoa(zfzhi.Shuzi1zhi())
c:=juesequanxiancontroller()
c.Ctx.Input.SetParam(zf.Id(false),paramid)
c.Delete()
log.Println(c.Data[zf.Json(true)])

}
func TestGetjuesequanxian(t *testing.T){
zf:=zf.Zf{}
zfzhi:=zfzhi.Zfzhi{}
paramid:=strconv.Itoa(zfzhi.Shuzi1zhi())
c:=juesequanxiancontroller()
c.Ctx.Input.SetParam(zf.Id(false),paramid)
c.Get()
log.Println(c.Data[zf.Json(true)])

}
