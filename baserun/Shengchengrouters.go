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

func routersimports(bianma string, buffer *bytes.Buffer) {
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	//"github.com/astaxie/beego"
	beestr := zfzhi.Zhi.Syh() + zf.Zfs.Github(true) + zfzhi.Zhi.Dh() + zf.Zfs.Com(true) + zfzhi.Zhi.Xx() + zf.Zfs.Astaxie(true) + zfzhi.Zhi.Xx() + zf.Zfs.Beego(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(beestr)

	//"hanfuxin/controllers"
	constr := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Controllers(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(constr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

}
func routersinit(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	//func init()
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Init(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//beego.Router("/zfzhi.Zhi.Xx()x",&controllers.Xzfzhi.Zhi.Xx()controller{})
	rstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.Router(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Syh() + zfzhi.Zhi.Xx() + bmx + zfzhi.Zhi.Syh() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Qh() + zf.Zfs.Controllers(true) + zfzhi.Zhi.Dh() + bianma + zf.Zfs.Controller(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(rstr)

	//beego.Router("/zfzhi.Zhi.Xx()x/:Id",&controllers.Xzfzhi.Zhi.Xx()controller{})
	ridstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.Router(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Syh() + zfzhi.Zhi.Xx() + bmx + zfzhi.Zhi.Xx() + zfzhi.Zhi.Mh() + zf.Zfs.Id(false) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Qh() + zf.Zfs.Controllers(true) + zfzhi.Zhi.Dh() + bianma + zf.Zfs.Controller(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ridstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func Shengchengrouters() {
	for bk, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		pacstr := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Routers(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pacstr)

		routersimports(bk, &buffer)
		routersinit(bk, &buffer)
		dir := basemodels.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Routers(true)
		path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Router(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
