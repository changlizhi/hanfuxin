package tests

import (
	"github.com/astaxie/beego/context"
	"hanfuxin/controllers"
	"hanfuxin/fortests"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strconv"
	"testing"
)

func juesecontroller() *controllers.Juesecontroller {
	c := controllers.Juesecontroller{}
	c.Data = make(map[interface{}]interface{})
	c.Ctx = context.NewContext()
	c.Ctx.Input = context.NewInput()
	c.Ctx.Output = context.NewOutput()
	c.Ctx.Output.Context = context.NewContext()
	c.Ctx.Output.Context.ResponseWriter = &context.Response{new(fortests.Mywriter), true, 200}
	return &c
}
func TestPostjuese(t *testing.T) {
	c := juesecontroller()
	reqjson := zfzhi.Zhi.Postjuesezhi()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Post()
	log.Println(c.Data)

}
func TestPatchjuese(t *testing.T) {
	c := juesecontroller()
	reqjson := zfzhi.Zhi.Patchjuesezhi()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Patch()
	log.Println(c.Data)

}
func TestDeletejuese(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())
	c := juesecontroller()
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false), paramid)
	c.Delete()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
func TestGetjuese(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())
	c := juesecontroller()
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false), paramid)
	c.Get()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
