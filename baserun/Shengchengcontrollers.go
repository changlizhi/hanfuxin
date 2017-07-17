package baserun

import (
	"bytes"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strings"
)

func controllerimports(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	hhf := zfzhi.Huanhangfuzhi()
	syh := zfzhi.Shuangyinhaozhi()
	xx := zfzhi.Xiexianzhi()
	dh := zfzhi.Dianhaozhi()

	bmx := strings.ToLower(bianma)

	// import
	buffer.WriteString(zf.Import(true))
	buffer.WriteString(xkhz + hhf)

	// "encoding/json"
	encstr := syh + zf.Encoding(true) + xx + zf.Json(true) + syh + hhf
	buffer.WriteString(encstr)
	// "github.com/astaxie/beego"
	asstr := syh + zf.Github(true) + dh + zf.Com(true) + xx + zf.Astaxie(true) + xx + zf.Beego(true) + syh + hhf
	buffer.WriteString(asstr)
	// 'hanfuxin/appmodels" \n
	apm := syh + zf.Hanfuxin(true) + xx + zf.Appmodels(true) + syh + hhf
	buffer.WriteString(apm)
	//"hanfuxin/baseinits"
	bistr := syh + zf.Hanfuxin(true) + xx + zf.Baseinits(true) + syh + hhf
	buffer.WriteString(bistr)
	//"hanfuxin/zdxxxservices"
	serstr := syh + zf.Hanfuxin(true) + xx + zf.Zd(true) + bmx + zf.Services(true) + syh + hhf
	buffer.WriteString(serstr)
	//"hanfuxin/zf"
	zfstr := syh + zf.Hanfuxin(true) + xx + zf.Zf(true) + syh + hhf
	buffer.WriteString(zfstr)
	//"hanfuxin/zfzhi"
	zfzhistr := syh + zf.Hanfuxin(true) + xx + zf.Zfzhi(true) + syh + hhf
	buffer.WriteString(zfzhistr)
	//"log"
	logstr := syh + zf.Log(true) + syh + hhf
	buffer.WriteString(logstr)
	//"strconv"
	convstr := syh + zf.Strconv(true) + syh + hhf
	buffer.WriteString(convstr)
	//"string"
	strstr := syh + zf.Strings(true) + syh + hhf
	buffer.WriteString(strstr)

	buffer.WriteString(hhf + xkhy + hhf)
}
func controllertype(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()

	typestr := zf.Type(true) + kgf + bianma + zf.Controller(true) + kgf + zf.Struct(true)
	buffer.WriteString(typestr)
	buffer.WriteString(dkhz + hhf)

	bcstr := zf.Beego(true) + dh + zf.Controller(false)
	buffer.WriteString(bcstr)

	buffer.WriteString(hhf + dkhy + hhf)
}
func controllerget(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	kgf := zfzhi.Konggefuzhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	dh := zfzhi.Dianhaozhi()
	douhao := zfzhi.Douhaozhi()
	gth := zfzhi.Gantanhaozhi()
	xh := zfzhi.Xinghaozhi()
	hhf := zfzhi.Huanhangfuzhi()

	bmx := strings.ToLower(bianma)

	//func (c *Xxxcontroller) Get()
	funcstr := zf.Func(true) + kgf + xkhz + zf.C(true) + kgf + xh + bianma + zf.Controller(true) + xkhy + zf.Get(false) + xkhz + xkhy
	buffer.WriteString(funcstr)
	buffer.WriteString(dkhz + hhf)

	//zf := zf.Zf{}
	zfstr := zf.Zf(true) + mh + dyh + zf.Zf(true) + dh + zf.Zf(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfstr)

	// canshu:=c.GetString(zf.Id(false))
	csstr := zf.Canshu(true) + mh + dyh + zf.C(true) + dh + zf.GetString(false) + xkhz + zf.Zf(true) + dh + zf.Id(false) + xkhz + zf.False(true) + xkhy + xkhy + hhf
	buffer.WriteString(csstr)

	//id,err := strconv.Atoi(canshu)
	idintstr := zf.Id(true) + douhao + zf.Err(true) + mh + dyh + zf.Strconv(true) + dh + zf.Atoi(false) + xkhz + zf.Canshu(true) + xkhy + hhf
	buffer.WriteString(idintstr)

	//if err != nil
	ifstr := zf.If(true) + kgf + zf.Err(true) + gth + dyh + zf.Nil(true)
	buffer.WriteString(ifstr)

	buffer.WriteString(dkhz + hhf)

	//log.Println(err)
	logerr := zf.Log(true) + dh + zf.Println(false) + xkhz + zf.Err(true) + xkhy + hhf
	buffer.WriteString(logerr)

	// c.Data[zf.Json(true)] = map[string]string
	dstr := zf.C(true) + dh + zf.Data(false) + zkhz + zf.Zf(true) + dh + zf.Json(false) + xkhz + zf.True(true) + xkhy + zkhy + dyh + zf.Map(true) + zkhz + zf.String(true) + zkhy + zf.String(true)
	buffer.WriteString(dstr)
	buffer.WriteString(dkhz + hhf)
	//zf.Error05(false):baseinits.Cuowus[zf.Error05(false)].Zhi,
	errretstr := zf.Zf(true) + dh + zf.Error05(false) + xkhz + zf.False(true) + xkhy + mh + zf.Baseinits(true) + dh + zf.Cuowus(false) + zkhz + zf.Zf(true) + dh + zf.Error05(false) + xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Zhi(false) + douhao + hhf
	buffer.WriteString(errretstr)
	buffer.WriteString(hhf + dkhy + hhf)

	//c.ServeJSON()
	servstr := zf.C(true) + dh + zf.ServeJSON(false) + xkhz + xkhy + hhf
	buffer.WriteString(servstr)
	//return
	buffer.WriteString(zf.Return(true))

	buffer.WriteString(hhf + dkhy + hhf)

	//xxx := zdxxxservices.Chaxunxxx(id)
	objret := bmx + mh + dyh + zf.Zd(true) + bmx + zf.Services(true) + dh + zf.Chaxun(false) + bmx + xkhz + zf.Id(true) + xkhy + hhf
	buffer.WriteString(objret)

	//c.Data[zf.Json(true)]=xxx
	serobj := zf.C(true) + dh + zf.Data(false) + zkhz + zf.Zf(true) + dh + zf.Json(false) + xkhz + zf.True(true) + xkhy + zkhy + dyh + bmx + hhf
	buffer.WriteString(serobj)

	buffer.WriteString(servstr)
	//return
	buffer.WriteString(zf.Return(true))

	buffer.WriteString(hhf + dkhy + hhf)
}

func patchpost(fangfaming string, bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xh := zfzhi.Xinghaozhi()
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	douhao := zfzhi.Douhaozhi()
	qiehao := zfzhi.Qiehaozhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	jiahao := zfzhi.Jiahaozhi()

	bmx := strings.ToLower(bianma)
	// func (c *Xxxcontroller) fangfaming()
	funstr := zf.Func(true) + kgf + xkhz + zf.C(true) + kgf + xh + bianma + zf.Controller(true) + xkhy + fangfaming + xkhz + xkhy
	buffer.WriteString(funstr)
	buffer.WriteString(dkhz + hhf)

	//zf:=zf.Zf{}
	zfstr := zf.Zf(true) + mh + dyh + zf.Zf(true) + dh + zf.Zf(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfstr)

	//zfzhi:=zfzhi.Zfzhi{}
	zfzstr := zf.Zfzhi(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Zfzhi(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfzstr)

	//kzf := zfzhi.Kongzifuzhi()
	kzfstr := zf.Kzf(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Kongzifuzhi(false) + xkhz + xkhy + hhf
	buffer.WriteString(kzfstr)

	//xhx := zfzhi.Xiahuaxianzhi()
	xhxstr := zf.Xhx(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Xiahuaxianzhi(false) + xkhz + xkhy + hhf
	buffer.WriteString(xhxstr)
	//sz0:=zfzhi.Shuzi0zhi()
	sz0str := zf.Sz0(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Shuzi0zhi(false) + xkhz + xkhy + hhf
	buffer.WriteString(sz0str)
	//sz1:=zfzhi.Shuzi1zhi()
	sz1str := zf.Sz1(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Shuzi1zhi(false) + xkhz + xkhy + hhf
	buffer.WriteString(sz1str)
	//mh := zfzhi.Maohaozhi()
	mhstr := zf.Mh(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Maohaozhi(false) + xkhz + xkhy + hhf
	buffer.WriteString(mhstr)

	//xxx:=appmodels.Xxx{}
	objstr := bmx + mh + dyh + zf.Appmodels(true) + dh + bianma + dkhz + dkhy + hhf
	buffer.WriteString(objstr)
	//json.Unmarshal(c.Ctx.Input.RequestBody,&xxx)
	jstr := zf.Json(true) + dh + zf.Unmarshal(false) + xkhz + zf.C(true) + dh + zf.Ctx(false) + dh + zf.Input(false) + dh + zf.RequestBody(false) + douhao + qiehao + bmx + xkhy + hhf
	buffer.WriteString(jstr)
	//serviceret := zdxxxservices.Tianjiaxxx(&xxx)

	postpatch := ""
	if fangfaming == zf.Post(false) {
		postpatch = zf.Tianjia(false)
	}
	if fangfaming == zf.Patch(false) {
		postpatch = zf.Xiugai(false)
	}

	sretstr := zf.Service(true) + zf.Ret(true) + mh + dyh + zf.Zd(true) + zf.Services(true) + dh + postpatch + bmx + xkhz + qiehao + bmx + xkhy + hhf
	buffer.WriteString(sretstr)
	//tishi:=baseinits.Tishis[serviceret].Zhi
	tsstr := zf.Tishi(true) + mh + dyh + zf.Baseinits(true) + dh + zf.Tishis(false) + zkhz + zf.Service(true) + zf.Ret(true) + zkhy + dh + zf.Zhi(false) + hhf
	buffer.WriteString(tsstr)
	//if tishi==kzf
	ifstr := zf.If(true) + kgf + zf.Tishi(true) + dyh + dyh + zf.Kzf(true)
	buffer.WriteString(ifstr)

	buffer.WriteString(dkhz + hhf)

	//splitret:=strings.Split(serviceret,xhx)
	spstr := zf.Split(true) + zf.Ret(true) + mh + dyh + zf.Strings(true) + dh + zf.Split(false) + xkhz + zf.Service(true) + zf.Ret(true) + douhao + zf.Xhx(true) + xkhy + hhf
	buffer.WriteString(spstr)
	//c.Data[zf.Json(true)]
	dstrs := zf.C(true) + dh + zf.Data(false) + zkhz + zf.Zf(true) + dh + zf.Json(false) + xkhz + zf.True(true) + xkhy + zkhy
	buffer.WriteString(dstrs)
	//=baseinits.Tishis[splitret[sz0]].Zhi+mh+splitret[sz1]
	davalstr := dyh + zf.Baseinits(true) + dh +
		zf.Tishis(false) + zkhz + zf.Split(true) + zf.Ret(true) +
		zkhz + zf.Sz0(true) + zkhy + zkhy + dh + zf.Zhi(false) +
		jiahao + zf.Mh(true) + jiahao +
		zf.Split(true) + zf.Ret(true) + zkhz + zf.Sz1(true) + zkhy + hhf
	buffer.WriteString(davalstr)
	//c.ServeJSON()
	serstr := zf.C(true) + dh + zf.ServeJSON(false) + xkhz + xkhy + hhf
	buffer.WriteString(serstr)
	buffer.WriteString(zf.Return(true))
	buffer.WriteString(hhf + dkhy + hhf)

	//c.Data[zf.Json(true)]=tishi
	tisret := zf.C(true) + dh + zf.Data(false) + zkhz + zf.Zf(true) + dh + zf.Json(false) + xkhz + zf.True(true) + xkhy + zkhy + dyh + zf.Tishi(true) + hhf
	buffer.WriteString(tisret)

	//c.ServeJSON()
	buffer.WriteString(serstr)
	buffer.WriteString(zf.Return(true))

	buffer.WriteString(hhf + dkhy + hhf)
}
func controllerpost(bianma string,buffer *bytes.Buffer){
	zf:=zf.Zf{}
	patchpost(zf.Post(false),bianma,buffer)
}
func controllerpatch(bianma string,buffer *bytes.Buffer){
	zf:=zf.Zf{}
	patchpost(zf.Patch(false),bianma,buffer)
}
func Shengchengcontrollers() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	buffer := bytes.Buffer{}
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()
	for bk, _ := range Huoqubiaos() {
		buffer.WriteString(zf.Package(true) + kgf + zf.Controllers(true) + hhf)
		controllerimports(bk, &buffer)
		controllertype(bk, &buffer)
		controllerget(bk, &buffer)
		controllerpost(bk, &buffer)
		controllerpatch(bk, &buffer)
	}
	log.Println("buffer-----------------------\n", buffer.String())
}
