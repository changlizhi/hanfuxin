package baserun

import (
	"bytes"
	"hanfuxin/basemodels"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"io/ioutil"
	"os"
	"strings"
)

func controllerimports(bianma string, buffer *bytes.Buffer) {

	bmx := strings.ToLower(bianma)

	// import
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	// "encoding/json"
	encstr := zfzhi.Zhi.Syh() + zf.Zfs.Encoding(true) + zfzhi.Zhi.Xx() + zf.Zfs.Json(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(encstr)

	// "github.com/astaxie/beego"
	asstr := zfzhi.Zhi.Syh() + zf.Zfs.Github(true) + zfzhi.Zhi.Dh() + zf.Zfs.Com(true) + zfzhi.Zhi.Xx() + zf.Zfs.Astaxie(true) + zfzhi.Zhi.Xx() + zf.Zfs.Beego(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(asstr)

	// 'hanfuxin/appmodels" \n
	apm := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(apm)

	//"hanfuxin/baseinits"
	bistr := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bistr)

	//"hanfuxin/zdzfzhi.Zhi.Xx()xservices"
	serstr := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + bmx + zf.Zfs.Services(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(serstr)

	//"hanfuxin/zf"
	zfstr := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfstr)

	//"hanfuxin/zfzhi"
	zfzhistr := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfzhistr)

	//"log"
	logstr := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logstr)

	//"strconv"
	convstr := zfzhi.Zhi.Syh() + zf.Zfs.Strconv(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(convstr)

	//"string"
	strstr := zfzhi.Zhi.Syh() + zf.Zfs.Strings(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(strstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

}
func controllertype(bianma string, buffer *bytes.Buffer) {

	typestr := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + bianma + zf.Zfs.Controller(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Struct(true)
	buffer.WriteString(typestr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	bcstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.Controller(false)
	buffer.WriteString(bcstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func controllerget(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	//func (c *Xzfzhi.Zhi.Xx()controller) Get()
	funcstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() + zf.Zfs.C(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + bianma + zf.Zfs.Controller(true) + zfzhi.Zhi.Xkhy() + zf.Zfs.Get(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funcstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	// canshu:=c.GetString(zfzhi.Zhi.Mh() + zf.Zfs.Id(false))
	csstr := zf.Zfs.Canshu(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.GetString(false) +
		zfzhi.Zhi.Xkhz() +

		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Mh(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +

		zfzhi.Zhi.Jia() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Id(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	//id,err := strconv.Atoi(canshu)
	idintstr := zf.Zfs.Id(true) + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Strconv(true) + zfzhi.Zhi.Dh() + zf.Zfs.Atoi(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.Canshu(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(idintstr)

	//if err != nil
	ifstr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(ifstr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//log.Println(err)
	logerr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.Err(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logerr)

	// c.Data[zf.Zfs.Json(true)] = map[string]string
	dstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Data(false) +
		zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.True(true) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Map(true) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.String(true) + zfzhi.Zhi.Zkhy() + zf.Zfs.String(true)
	buffer.WriteString(dstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//zf.Zfs.Error05(false):baseinits.Cuowus[zf.Zfs.Error05(false)].Zhi,
	errretstr := zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Error05(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Mh() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Cuowus(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Error05(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(errretstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//c.ServeJSON()
	servstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.ServeJSON(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(servstr)
	//return
	buffer.WriteString(zf.Zfs.Return(true))

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//zfzhi.Zhi.Xx()x := zdzfzhi.Zhi.Xx()xservices.Chaxunzfzhi.Zhi.Xx()x(id)
	objret := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Zd(true) + bmx + zf.Zfs.Services(true) + zfzhi.Zhi.Dh() + zf.Zfs.Chaxun(false) + bmx + zfzhi.Zhi.Xkhz() + zf.Zfs.Id(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(objret)

	//c.Data[zf.Zfs.Json(true)]=zfzhi.Zhi.Xx()x
	serobj := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Data(false) +
		zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() +
		zf.Zfs.Json(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.True(true) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Dyh() + bmx + zfzhi.Zhi.Hhf()
	buffer.WriteString(serobj)

	buffer.WriteString(servstr)
	//return
	buffer.WriteString(zf.Zfs.Return(true))

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func patchpost(fangfaming string, bianma string, buffer *bytes.Buffer) {

	bmx := strings.ToLower(bianma)
	// func (c *Xzfzhi.Zhi.Xx()controller) fangfaming()
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() + zf.Zfs.C(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + bianma + zf.Zfs.Controller(true) + zfzhi.Zhi.Xkhy() + fangfaming + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//xxx := appmodels.Xxx{}
	objstr := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(objstr)

	//json.Unmarshal(c.Ctx.Input.RequestBody,&zfzhi.Zhi.Xx()x)
	jstr := zf.Zfs.Json(true) + zfzhi.Zhi.Dh() + zf.Zfs.Unmarshal(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ctx(false) + zfzhi.Zhi.Dh() + zf.Zfs.Input(false) + zfzhi.Zhi.Dh() + zf.Zfs.RequestBody(false) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Qh() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(jstr)

	postpatch := ""
	if fangfaming == zf.Zfs.Post(false) {
		postpatch = zf.Zfs.Tianjia(false)
	}
	if fangfaming == zf.Zfs.Patch(false) {
		postpatch = zf.Zfs.Xiugai(false)
	}

	//serviceret := zdzfzhi.Zhi.Xx()xservices.Tianjiazfzhi.Zhi.Xx()x(&zfzhi.Zhi.Xx()x)
	sretstr := zf.Zfs.Service(true) + zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Zd(true) + bmx + zf.Zfs.Services(true) + zfzhi.Zhi.Dh() + postpatch + bmx + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Qh() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(sretstr)

	//tishi:=baseinits.Tishis[serviceret].Zhi
	tsstr := zf.Zfs.Tishi(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Service(true) + zf.Zfs.Ret(true) + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(tsstr)

	//if tishi==kzf
	ifstr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tishi(true) + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyh() +

		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Kzf(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()

	buffer.WriteString(ifstr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//splitret:=strings.Split(serviceret,xhx)
	spstr := zf.Zfs.Split(true) + zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Dyh() + zf.Zfs.Strings(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Split(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.Service(true) +
		zf.Zfs.Ret(true) + zfzhi.Zhi.Dou() +

		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Xhx(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() +

		zfzhi.Zhi.Hhf()
	buffer.WriteString(spstr)

	//c.Data[zf.Zfs.Json(true)]
	dstrs := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Data(false) +
		zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.True(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy()
	buffer.WriteString(dstrs)

	//=baseinits.Tishis[splitret[sz0]].Zhi+zfzhi.Zhi.Mh()+splitret[sz1]
	davalstr := zfzhi.Zhi.Dyh() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Split(true) + zf.Zfs.Ret(true) +
		zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Shuzi0zhi(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) +
		zfzhi.Zhi.Jia() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Mh(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Jia() +
		zf.Zfs.Split(true) + zf.Zfs.Ret(true) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Shuzi1zhi(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Hhf()
	buffer.WriteString(davalstr)

	//c.ServeJSON()
	serstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.ServeJSON(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(serstr)
	buffer.WriteString(zf.Zfs.Return(true))
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//c.Data[zf.Zfs.Json(true)]=tishi
	tisret := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Data(false) +
		zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.True(true) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Tishi(true) + zfzhi.Zhi.Hhf()
	buffer.WriteString(tisret)

	//c.ServeJSON()
	buffer.WriteString(serstr)
	buffer.WriteString(zf.Zfs.Return(true))

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func controllerpost(bianma string, buffer *bytes.Buffer) {
	patchpost(zf.Zfs.Post(false), bianma, buffer)
}
func controllerpatch(bianma string, buffer *bytes.Buffer) {
	patchpost(zf.Zfs.Patch(false), bianma, buffer)
}
func controllerdelete(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	// func (c *Xxxcontroller) Delete()
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() + zf.Zfs.C(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + bianma + zf.Zfs.Controller(true) + zfzhi.Zhi.Xkhy() + zf.Zfs.Delete(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	// canshu:=c.GetString(zfzhi.Zhi.Mh()+zf.Zfs.Id(false))
	csstr := zf.Zfs.Canshu(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.GetString(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Mh(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Jia() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Id(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	//id,err := strconv.Atoi(canshu)
	idintstr := zf.Zfs.Id(true) + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Strconv(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Atoi(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Canshu(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(idintstr)
	// if err !=nil
	ifstr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(ifstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	// log.Println(err)
	logstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Err(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logstr)
	// c.Data[zf.Zfs.Json(true)] = map[string]string
	dstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Data(false) +
		zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.True(true) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Map(true) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.String(true) + zfzhi.Zhi.Zkhy() + zf.Zfs.String(true)
	buffer.WriteString(dstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//zf.Zfs.Error05(false):baseinits.Cuowus[zf.Zfs.Error05(false)].Zhi,
	errretstr := zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Error05(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Mh() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Cuowus(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Error05(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(errretstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//c.ServeJSON()
	servstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.ServeJSON(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(servstr)
	//return
	buffer.WriteString(zf.Zfs.Return(true))

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//serviceret := zdzfzhi.Zhi.Xx()xservices.Shanchuzfzhi.Zhi.Xx()x(id)
	sretstr := zf.Zfs.Service(true) + zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Dyh() + zf.Zfs.Zd(true) + bmx + zf.Zfs.Services(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Shanchu(false) + bmx + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Id(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(sretstr)
	//tishi:=baseinits.Tishis[serviceret].Zhi
	tsstr := zf.Zfs.Tishi(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Service(true) + zf.Zfs.Ret(true) + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(tsstr)
	//if tishi==kzf
	iftsstr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tishi(true) + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Kzf(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(iftsstr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//splitret:=strings.Split(serviceret,xhx)
	spstr := zf.Zfs.Split(true) + zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Dyh() + zf.Zfs.Strings(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Split(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.Service(true) + zf.Zfs.Ret(true) + zfzhi.Zhi.Dou() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Xhx(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(spstr)
	//c.Data[zf.Zfs.Json(true)]
	dstrs := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Data(false) +
		zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.True(true) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy()
	buffer.WriteString(dstrs)
	//=baseinits.Tishis[splitret[sz0]].Zhi+zfzhi.Zhi.Mh()+splitret[sz1]
	davalstr := zfzhi.Zhi.Dyh() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Split(true) + zf.Zfs.Ret(true) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Shuzi0zhi(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Jia() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Mh(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Jia() + zf.Zfs.Split(true) + zf.Zfs.Ret(true) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Shuzi1zhi(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(davalstr)
	//c.ServeJSON()
	serstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.ServeJSON(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(serstr)
	buffer.WriteString(zf.Zfs.Return(true))
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//c.Data[zf.Zfs.Json(true)]=tishi
	tisret := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Data(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.True(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dyh() + zf.Zfs.Tishi(true) + zfzhi.Zhi.Hhf()
	buffer.WriteString(tisret)

	//c.ServeJSON()
	buffer.WriteString(serstr)
	buffer.WriteString(zf.Zfs.Return(true))

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func Shengchengcontrollers() {
	for bk, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		buffer.WriteString(zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Controllers(true) + zfzhi.Zhi.Hhf())
		controllerimports(bk, &buffer)
		controllertype(bk, &buffer)
		controllerget(bk, &buffer)
		controllerpost(bk, &buffer)
		controllerpatch(bk, &buffer)
		controllerdelete(bk, &buffer)
		dir := basemodels.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Controllers(true)
		path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Controller(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
