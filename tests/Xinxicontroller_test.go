package tests

import (
	"github.com/astaxie/beego/context"
	"hanfuxin/controllers"
	"hanfuxin/tests"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strconv"
	"testing"
)

func xinxicontroller() *controllers.Xinxicontroller {
	c := controllers.Xinxicontroller{}
	c.Data = make(map[interface{}]interface{})
	c.Ctx = context.NewContext()
	c.Ctx.Input = context.NewInput()
	c.Ctx.Output = context.NewOutput()
	c.Ctx.Output.Context = context.NewContext()
	c.Ctx.Output.Context.ResponseWriter = &context.Response{new(tests.Mywriter), true, 200}
	return &c
}
func TestPostxinxi(t *testing.T) {
	zfzhi := zfzhi.Zfzhi{}
	c := xinxicontroller()
	reqjson := zfzhi.Postxinxizhi()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Post()
	log.Println(c.Data)

}
func TestPatchxinxi(t *testing.T) {
	zfzhi := zfzhi.Zfzhi{}
	c := xinxicontroller()
	reqjson := zfzhi.Patchxinxizhi()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Patch()
	log.Println(c.Data)

}
func TestDeletexinxi(t *testing.T) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	mh := zfzhi.Maohaozhi()
	paramid := strconv.Itoa(zfzhi.Shuzi1zhi())
	c := xinxicontroller()
	c.Ctx.Input.SetParam(mh+zf.Id(false), paramid)
	c.Delete()
	log.Println(c.Data[zf.Json(true)])

}
func TestGetxinxi(t *testing.T) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	mh := zfzhi.Maohaozhi()
	paramid := strconv.Itoa(zfzhi.Shuzi1zhi())
	c := xinxicontroller()
	c.Ctx.Input.SetParam(mh+zf.Id(false), paramid)
	c.Get()
	log.Println(c.Data[zf.Json(true)])

}
