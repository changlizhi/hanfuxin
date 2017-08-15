package baserun

import (
	"bytes"
	"hanfuxin/appinits"
	"changliang/zf"
	"changliang/zfzhi"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func servicetestimport(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	importstr := zf.Zfs.Import(true) + zfzhi.Zhi.Kgf()
	buffer.WriteString(importstr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	// "hanfuxin/appmodels"
	apmstr := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(apmstr)

	// "hanfuxin/zdjueseservices"
	serstr := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + bmx + zf.Zfs.Services(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(serstr)

	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		if Huoquyigeleixing(lk) == zf.Zfs.Time(true) {
			// "time" \n
			timebao := zfzhi.Zhi.Syh() + zf.Zfs.Time(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
			buffer.WriteString(timebao)
			break
		}
	}

	//"log"
	logstr := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logstr)
	// "testing"
	tstr := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy())
}
func tianjiaxiugai(fangfa string, bianma string, buffer *bytes.Buffer) {
	sz1str := strconv.Itoa(zfzhi.Zhi.Shuzi1())
	sz0str := strconv.Itoa(zfzhi.Zhi.Shuzi0())

	bmx := strings.ToLower(bianma)
	//func TestJueseserviceTianjia(t *testing.T)
	funcstr := zfzhi.Zhi.Hhf() + zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + bianma + zf.Zfs.Services(true) + fangfa + zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funcstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//juese := appmodels.Juese
	objstr := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(objstr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		// Id : 1,
		if zf.Zfs.Int(true) == Huoquyigeleixing(lk) {
			intstr := lk + zfzhi.Zhi.Mh() + sz1str + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(intstr)
			continue
		}
		// Bianma :"",
		if zf.Zfs.String(true) == Huoquyigeleixing(lk) {
			strstr := lk + zfzhi.Zhi.Mh() + zfzhi.Zhi.Syh() + bianma + fangfa + zfzhi.Zhi.Syh() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(strstr)
			continue
		}
		// time : timeNow()
		if zf.Zfs.Time(true) == Huoquyigeleixing(lk) {
			timestr := lk + zfzhi.Zhi.Mh() + zf.Zfs.Time(true) + zfzhi.Zhi.Dh() + zf.Zfs.Now(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(timestr)
			continue
		}
		// Qian : 1.0
		if zf.Zfs.Float(true) == Huoquyigeleixing(lk) {
			flostr := lk + zfzhi.Zhi.Mh() + sz1str + zfzhi.Zhi.Dh() + sz0str + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(flostr)
			continue
		}
	}
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy())
	//zdzfzhi.Zhi.Xx()xservices.Tianjiazfzhi.Zhi.Xx()x(&zfzhi.Zhi.Xx()x)
	tjstr := zfzhi.Zhi.Hhf() + zf.Zfs.Zd(true) + bmx + zf.Zfs.Services(true) + zfzhi.Zhi.Dh() + fangfa + bmx + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Qh() + bmx + zfzhi.Zhi.Xkhy()
	buffer.WriteString(tjstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testservicetianjia(bianma string, buffer *bytes.Buffer) {
	tianjiaxiugai(zf.Zfs.Tianjia(false), bianma, buffer)
}
func testservicexiugai(bianma string, buffer *bytes.Buffer) {
	tianjiaxiugai(zf.Zfs.Xiugai(false), bianma, buffer)
}
func testservicechaxun(bianma string, buffer *bytes.Buffer) {
	sz1str := strconv.Itoa(zfzhi.Zhi.Shuzi1())
	bmx := strings.ToLower(bianma)
	// func TestXzfzhi.Zhi.Xx()servicesChaxun
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + bianma + zf.Zfs.Services(true) + zf.Zfs.Chaxun(false)
	buffer.WriteString(funstr)
	//(t *testing.T)
	tstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	// zfzhi.Zhi.Xx()x := zdzfzhi.Zhi.Xx()xservices.Chaxunzfzhi.Zhi.Xx()x(1)
	objstr := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Zd(true) + bmx + zf.Zfs.Services(true) + zfzhi.Zhi.Dh() + zf.Zfs.Chaxun(false) + bmx + zfzhi.Zhi.Xkhz() + sz1str + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(objstr)
	logstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) + zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testserviceshanchu(bianma string, buffer *bytes.Buffer) {
	sz1str := strconv.Itoa(zfzhi.Zhi.Shuzi1())
	bmx := strings.ToLower(bianma)
	// func TestXzfzhi.Zhi.Xx()serviceShanchu
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + bianma + zf.Zfs.Services(true) + zf.Zfs.Shanchu(false)
	buffer.WriteString(funstr)
	//(t *testing.T)
	tstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//zdzfzhi.Zhi.Xx()xservices.Shanchuzfzhi.Zhi.Xx()x(1)
	scstr := zf.Zfs.Zd(true) + bmx + zf.Zfs.Services(true) + zfzhi.Zhi.Dh() + zf.Zfs.Shanchu(false) + bmx + zfzhi.Zhi.Xkhz() + sz1str + zfzhi.Zhi.Xkhy()
	buffer.WriteString(scstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func Shengchengservicetest() {
	for bk, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		pacstr := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pacstr)
		servicetestimport(bk, &buffer)
		testservicetianjia(bk, &buffer)
		testservicexiugai(bk, &buffer)
		testservicechaxun(bk, &buffer)
		testserviceshanchu(bk, &buffer)
		dir := appinits.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
		path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Services(true) + zfzhi.Zhi.Xhx() + zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}

}
