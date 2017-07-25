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

func juesequanxiancontroller() *controllers.Juesequanxiancontroller {
	c := controllers.Juesequanxiancontroller{}
	c.Data = make(map[interface{}]interface{})
	c.Ctx = context.NewContext()
	c.Ctx.Input = context.NewInput()
	c.Ctx.Output = context.NewOutput()
	c.Ctx.Output.Context = context.NewContext()
	c.Ctx.Output.Context.ResponseWriter = &context.Response{new(fortests.Mywriter), true, 200}
	return &c
}
func TestPostjuesequanxian(t *testing.T) {
	c := juesequanxiancontroller()
	reqjson := zfzhi.Zhi.Postjuesequanxianzhi()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Post()
	log.Println(c.Data)

}
func TestPatchjuesequanxian(t *testing.T) {
	c := juesequanxiancontroller()
	reqjson := zfzhi.Zhi.Patchjuesequanxianzhi()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Patch()
	log.Println(c.Data)

}
func TestDeletejuesequanxian(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())
	c := juesequanxiancontroller()
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false), paramid)
	c.Delete()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
func TestGetjuesequanxian(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())
	c := juesequanxiancontroller()
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false), paramid)
	c.Get()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
