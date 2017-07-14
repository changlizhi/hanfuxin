package tests

import (
	"github.com/astaxie/beego/context"
	"hanfuxin/controllers"
	_ "hanfuxin/routers"
	"hanfuxin/tests"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strconv"
	"testing"
)

func controller() *controllers.Juesecontroller {
	c := controllers.Juesecontroller{}
	c.Data = make(map[interface{}]interface{})
	c.Ctx = context.NewContext()
	c.Ctx.Input = context.NewInput()
	c.Ctx.Output = context.NewOutput()
	c.Ctx.Output.Context = context.NewContext()
	c.Ctx.Output.Context.ResponseWriter = &context.Response{new(tests.Mywriter), true, 200}
	return &c
}
func TestPostjuese(t *testing.T) {
	c := controller()
	reqstr := `{"Id":1,"Bianma":"ROLE","Mingcheng":"角色1","Biaoji":"Youxiao"}`
	c.Ctx.Input.RequestBody = []byte(reqstr)
	c.Post()
	log.Println("cress-----------", c.Data)
}
func TestGetjuese(t *testing.T) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	idstr := strconv.Itoa(zfzhi.Shuzi1zhi())
	c := controller()
	c.Ctx.Input.SetParam(zf.Id(false), idstr)
	c.Get()
	log.Println("cressget-----------", c.Data[zf.Json(true)])
}
func TestDeletejuese(t *testing.T) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	idstr := strconv.Itoa(zfzhi.Shuzi2zhi())
	c := controller()
	c.Ctx.Input.SetParam(zf.Id(false), idstr)
	c.Delete()
	log.Println("cressget-----------", c.Data[zf.Json(true)])
}
func TestPatchjuese(t *testing.T) {
	c := controller()
	reqstr := `{"Id":1,"Bianma":"ROLE222","Mingcheng":"角色222","Biaoji":"Youxiao"}`
	c.Ctx.Input.RequestBody = []byte(reqstr)
	c.Patch()
	log.Println("cress-----------", c.Data)
}
