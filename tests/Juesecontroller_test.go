package tests

import (
	"testing"
	"github.com/astaxie/beego/context"
	_ "hanfuxin/routers"
	"log"
	"hanfuxin/controllers"
	"hanfuxin/tests"
)
func TestPostjuese(t *testing.T) {
	c := controllers.Juesecontroller{}
	c.Data = make(map[interface{}]interface{})
	c.Ctx = context.NewContext()
	c.Ctx.Input = context.NewInput()
	reqstr := `{"Id":5,"Bianma":"ROLE","Mingcheng":"角色1","Biaoji":"Youxiao"}`
	c.Ctx.Input.RequestBody = []byte(reqstr)
	c.Ctx.Output = context.NewOutput()
	c.Ctx.Output.Context = context.NewContext()
	c.Ctx.Output.Context.ResponseWriter = &context.Response{new(tests.Mywriter), true, 200}
	log.Println("creq-----------", string(c.Ctx.Input.RequestBody))
	c.Post()
	log.Println("cress-----------", c.Data)
}
func TestGetjuese(t *testing.T){
	c := controllers.Juesecontroller{}
	c.Data = make(map[interface{}]interface{})
	c.Ctx = context.NewContext()
	c.Ctx.Input = context.NewInput()
	reqstr := `{"Id":5,"Bianma":"ROLE","Mingcheng":"角色1","Biaoji":"Youxiao"}`
	c.Ctx.Input.RequestBody = []byte(reqstr)
	c.Ctx.Output = context.NewOutput()
	c.Ctx.Output.Context = context.NewContext()
	c.Ctx.Output.Context.ResponseWriter = &context.Response{new(tests.Mywriter), true, 200}
	log.Println("creq-----------", string(c.Ctx.Input.RequestBody))
	c.Get()
	log.Println("cress-----------", c.TplName)
}
