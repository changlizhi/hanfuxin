package baserun

import (
	"bytes"
	"hanfuxin/basemodels"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func controllertestimports(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dh := zfzhi.Dianhaozhi()
	xx := zfzhi.Xiexianzhi()
	syh := zfzhi.Shuangyinhaozhi()

	buffer.WriteString(zf.Import(true))
	buffer.WriteString(xkhz + hhf)
	//"github.com/astaxie/beego/context"
	gitstr := syh + zf.Github(true) + dh + zf.Com(true) + xx + zf.Astaxie(true) + xx + zf.Beego(true) + xx + zf.Context(true) + syh + hhf
	buffer.WriteString(gitstr)

	//"hanfuxin/controllers"
	constr := syh + zf.Hanfuxin(true) + xx + zf.Controllers(true) + syh + hhf
	buffer.WriteString(constr)

	//"hanfuxin/tests"
	teststr := syh + zf.Hanfuxin(true) + xx + zf.Tests(true) + syh + hhf
	buffer.WriteString(teststr)
	//"hanfuxin/zf"
	zfstr := syh + zf.Hanfuxin(true) + xx + zf.Zf(true) + syh + hhf
	buffer.WriteString(zfstr)
	// "hanfuxin/zfzhi"
	zfzhistr := syh + zf.Hanfuxin(true) + xx + zf.Zfzhi(true) + syh + hhf
	buffer.WriteString(zfzhistr)

	//"log"
	logstr := syh + zf.Log(true) + syh + hhf
	buffer.WriteString(logstr)
	//"strconv"
	strcstr := syh + zf.Strconv(true) + syh + hhf
	buffer.WriteString(strcstr)
	//testing
	testingstr := syh + zf.Testing(true) + syh + hhf
	buffer.WriteString(testingstr)

	buffer.WriteString(hhf + xkhy + hhf)
}
func bkcontroller(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()
	bmx := strings.ToLower(bianma)
	xh := zfzhi.Xinghaozhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dh := zfzhi.Dianhaozhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	sz200 := zfzhi.Shuzi2zhi() * zfzhi.Shuzi10zhi() * zfzhi.Shuzi10zhi()
	sz200str := strconv.Itoa(sz200)

	qh := zfzhi.Qiehaozhi()
	douhao := zfzhi.Douhaozhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()

	// func xxxcontroller() *controllers.Xxxcontroller
	funstr := zf.Func(true) + kgf + bmx + zf.Controller(true) + xkhz + xkhy + xh + zf.Controllers(true) + dh + bianma + zf.Controller(true)
	buffer.WriteString(funstr)
	buffer.WriteString(dkhz + hhf)

	//  c := controllers.Xxxcontroller{}
	objcon := zf.C(true) + mh + dyh + zf.Controllers(true) + dh + bianma + zf.Controller(true) + dkhz + dkhy + hhf
	buffer.WriteString(objcon)
	// c.Data=make(map[interface{}]interface{})
	newd := zf.C(true) + dh + zf.Data(false) + dyh + zf.Make(true) + xkhz + zf.Map(true) + zkhz + zf.Interface(true) + dkhz + dkhy + zkhy + zf.Interface(true) + dkhz + dkhy + xkhy + hhf
	buffer.WriteString(newd)
	// c.Ctx=context.NewContext()
	newcon := zf.C(true) + dh + zf.Ctx(false) + dyh + zf.Context(true) + dh + zf.NewContext(false) + xkhz + xkhy + hhf
	buffer.WriteString(newcon)
	// c.Ctx.Input =context.NewInput()
	newin := zf.C(true) + dh + zf.Ctx(false) + dh + zf.Input(false) + dyh + zf.Context(true) + dh + zf.NewInput(false) + xkhz + xkhy + hhf
	buffer.WriteString(newin)
	//c.Ctx.Output=context.NewOutput()
	newout := zf.C(true) + dh + zf.Ctx(false) + dh + zf.Output(false) + dyh + zf.Context(true) + dh + zf.NewOutput(false) + xkhz + xkhy + hhf
	buffer.WriteString(newout)
	//c.Ctx.Output.Context = context.NewContext()
	outcon := zf.C(true) + dh + zf.Ctx(false) + dh + zf.Output(false) + dh + zf.Context(false) + dyh + zf.Context(true) + dh + zf.NewContext(false) + xkhz + xkhy + hhf
	buffer.WriteString(outcon)
	//c.Ctx.Output.Context.ResponseWriter = &context.Response
	conresp := zf.C(true) + dh + zf.Ctx(false) + dh + zf.Output(false) + dh + zf.Context(false) + dh + zf.ResponseWriter(false) + dyh + qh + zf.Context(true) + dh + zf.Response(false)
	buffer.WriteString(conresp)
	//{new(texts.Mywriter),true,200}
	respobj := dkhz + zf.New(true) + xkhz + zf.Tests(true) + dh + zf.Mywriter(false) + xkhy + douhao + zf.True(true) + douhao + sz200str + dkhy + hhf
	buffer.WriteString(respobj)
	//return &c
	retcon := zf.Return(true) + kgf + qh + zf.C(true)
	buffer.WriteString(retcon)

	buffer.WriteString(hhf + dkhy + hhf)
}
func testpostpatch(fangfa string, bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()
	bmx := strings.ToLower(bianma)
	xh := zfzhi.Xinghaozhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	dh := zfzhi.Dianhaozhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()

	//func TestFangfajuese(t *testing.T)
	funstr := zf.Func(true) + kgf + zf.Test(false) + fangfa + bmx + xkhz + zf.T(true) + xh + zf.Testing(true) + dh + zf.T(false) + xkhy
	buffer.WriteString(funstr)
	buffer.WriteString(dkhz + hhf)
	// zfzhi := zfzhi.Zfzhi{}
	zfzhistr := zf.Zfzhi(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Zfzhi(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfzhistr)

	//c:= xxxcontroller()
	contr := zf.C(true) + mh + dyh + bmx + zf.Controller(true) + xkhz + xkhy + hhf
	buffer.WriteString(contr)

	// reqjson := zfzhi.Fangfajuesezhi()
	jsonstr := zf.Req(true) + zf.Json(true) + mh + dyh + zf.Zfzhi(true) + dh + fangfa + bmx + zf.Zhi(true) + xkhz + xkhy + hhf
	buffer.WriteString(jsonstr)
	// c.Ctx.Input.RequestBody = []byte(reqjson)
	cinstr := zf.C(true) + dh + zf.Ctx(false) + dh + zf.Input(false) + dh + zf.RequestBody(false) + dyh + zkhz + zkhy + zf.Byte(true) + xkhz + zf.Req(true) + zf.Json(true) + xkhy + hhf
	buffer.WriteString(cinstr)
	//c.Fangfa()
	pstr := zf.C(true) + dh + fangfa + xkhz + xkhy + hhf
	buffer.WriteString(pstr)
	//log.Println(c.Data)
	logstr := zf.Log(true) + dh + zf.Println(false) + xkhz + zf.C(true) + dh + zf.Data(false) + xkhy + hhf
	buffer.WriteString(logstr)

	buffer.WriteString(hhf + dkhy + hhf)
}
func testpostcontroller(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	testpostpatch(zf.Post(false), bianma, buffer)
}
func testpatchcontroller(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	testpostpatch(zf.Patch(false), bianma, buffer)
}
func deletegetcontroller(fangfa string, bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()
	bmx := strings.ToLower(bianma)
	xh := zfzhi.Xinghaozhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()

	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	dh := zfzhi.Dianhaozhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()

	douhao := zfzhi.Douhaozhi()
	// func TestDeletexxx(t * testing.T)
	funstr := zf.Func(true) + kgf + zf.Test(false) + fangfa + bmx + xkhz + zf.T(true) + kgf + xh + zf.Testing(true) + dh + zf.T(false) + xkhy
	buffer.WriteString(funstr)
	buffer.WriteString(dkhz + hhf)

	// zf:= zf.Zf{}
	zfstr := zf.Zf(true) + mh + dyh + zf.Zf(true) + dh + zf.Zf(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfstr)

	// zfzhi := zfzhi.Zfzhi{}
	zfzhistr := zf.Zfzhi(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Zfzhi(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfzhistr)
	//paramid := strconv.Itoa(zfzhi.Shuzi1zhi())
	parstr := zf.Param(true) + zf.Id(true) + mh + dyh + zf.Strconv(true) + dh + zf.Itoa(false) + xkhz + zf.Zfzhi(true) + dh + zf.Shuzi1zhi(false) + xkhz + xkhy + xkhy + hhf
	buffer.WriteString(parstr)
	//c := xxxcontroller()
	cstr := zf.C(true) + mh + dyh + bmx + zf.Controller(true) + xkhz + xkhy + hhf
	buffer.WriteString(cstr)
	//c.Ctx.Input.SetParam(zf.Id(false),paramid)
	cinstr := zf.C(true) + dh + zf.Ctx(false) + dh + zf.Input(false) + dh + zf.SetParam(false)
	buffer.WriteString(cinstr)
	//(zf.Id(false),paramid)
	cinparam := xkhz + zf.Zf(true) + dh + zf.Id(false) + xkhz + zf.False(true) + xkhy + douhao + zf.Param(true) + zf.Id(true) + xkhy + hhf
	buffer.WriteString(cinparam)
	// c.Get()
	callstr := zf.C(true) + dh + fangfa + xkhz + xkhy + hhf
	buffer.WriteString(callstr)
	//log.Println(c.Data[zf.Json(true)])
	logstr := zf.Log(true) + dh + zf.Println(false) + xkhz + zf.C(true) + dh + zf.Data(false) + zkhz + zf.Zf(true) + dh + zf.Json(false) + xkhz + zf.True(true) + xkhy + zkhy + xkhy + hhf
	buffer.WriteString(logstr)

	buffer.WriteString(hhf + dkhy + hhf)

}
func testdeletecontroller(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	deletegetcontroller(zf.Delete(false), bianma, buffer)

}
func testgetcontroller(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	deletegetcontroller(zf.Get(false), bianma, buffer)

}
func Shengchengcontrollertest() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()
	xx := zfzhi.Xiexianzhi()
	dh := zfzhi.Dianhaozhi()
	xhx := zfzhi.Xiahuaxianzhi()
	for bk, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		buffer.WriteString(zf.Package(true) + kgf + zf.Tests(true) + hhf)
		controllertestimports(bk, &buffer)
		bkcontroller(bk, &buffer)
		testpostcontroller(bk, &buffer)
		testpatchcontroller(bk, &buffer)
		testdeletecontroller(bk, &buffer)
		testgetcontroller(bk, &buffer)

		log.Println("buffer-------\n", buffer.String())
		dir := basemodels.Getapppath() + xx + zf.Tests(true)
		path := dir + xx + bk + zf.Controller(true) + xhx + zf.Test(true) + dh + zf.Go(true)
		log.Println("dir---------", dir)
		log.Println("path--------", path)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
