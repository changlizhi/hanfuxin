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

func yanzhengleixingcontroller() *controllers.Yanzhengleixingcontroller {
	c := controllers.Yanzhengleixingcontroller{}
	c.Data = make(map[interface{}]interface{})
	c.Ctx = context.NewContext()
	c.Ctx.Input = context.NewInput()
	c.Ctx.Output = context.NewOutput()
	c.Ctx.Output.Context = context.NewContext()
	c.Ctx.Output.Context.ResponseWriter = &context.Response{new(fortests.Mywriter), true, 200}
	return &c
}
func TestPostyanzhengleixing(t *testing.T) {
	c := yanzhengleixingcontroller()
	reqjson := zfzhi.Zhi.Postyanzhengleixingzhi()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Post()
	log.Println(c.Data)

}
func TestPatchyanzhengleixing(t *testing.T) {
	c := yanzhengleixingcontroller()
	reqjson := zfzhi.Zhi.Patchyanzhengleixingzhi()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Patch()
	log.Println(c.Data)

}
func TestDeleteyanzhengleixing(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())
	c := yanzhengleixingcontroller()
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false), paramid)
	c.Delete()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
func TestGetyanzhengleixing(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())
	c := yanzhengleixingcontroller()
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false), paramid)
	c.Get()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
