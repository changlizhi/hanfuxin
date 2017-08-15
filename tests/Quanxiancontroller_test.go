package tests

import (
	"github.com/astaxie/beego/context"
	"hanfuxin/controllers"
	"hanfuxin/fortests"
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"strconv"
	"testing"
)

func quanxiancontroller() *controllers.Quanxiancontroller {
	c := controllers.Quanxiancontroller{}
	c.Data = make(map[interface{}]interface{})
	c.Ctx = context.NewContext()
	c.Ctx.Input = context.NewInput()
	c.Ctx.Output = context.NewOutput()
	c.Ctx.Output.Context = context.NewContext()
	c.Ctx.Output.Context.ResponseWriter = &context.Response{new(fortests.Mywriter), true, 200}
	return &c
}
func TestPostquanxian(t *testing.T) {
	c := quanxiancontroller()
	reqjson := zfzhi.Zhi.Postquanxianzhi()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Post()
	log.Println(c.Data)

}
func TestPatchquanxian(t *testing.T) {
	c := quanxiancontroller()
	reqjson := zfzhi.Zhi.Patchquanxianzhi()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Patch()
	log.Println(c.Data)

}
func TestDeletequanxian(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())
	c := quanxiancontroller()
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false), paramid)
	c.Delete()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
func TestGetquanxian(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())
	c := quanxiancontroller()
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false), paramid)
	c.Get()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
