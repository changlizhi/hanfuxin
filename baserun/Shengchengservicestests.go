package baserun

import (
	"bytes"
	"hanfuxin/basemodels"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func servicetestimport(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	syh := zfzhi.Shuangyinhaozhi()
	xx := zfzhi.Xiexianzhi()

	bmx := strings.ToLower(bianma)
	importstr := zf.Import(true) + kgf
	buffer.WriteString(importstr)
	buffer.WriteString(xkhz + hhf)

	// "hanfuxin/appmodels"
	apmstr := syh + zf.Hanfuxin(true) + xx + zf.Appmodels(true) + syh + hhf
	buffer.WriteString(apmstr)

	// "hanfuxin/zdjueseservices"
	serstr := syh + zf.Hanfuxin(true) + xx + zf.Zd(true) + bmx + zf.Services(true) + syh + hhf
	buffer.WriteString(serstr)

	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		if Huoquyigeleixing(lk) == zf.Time(true) {
			// "time" \n
			timebao := syh + zf.Time(true) + syh + hhf
			buffer.WriteString(timebao)
			break
		}
	}

	//"log"
	logstr := syh + zf.Log(true) + syh + hhf
	buffer.WriteString(logstr)
	// "testing"
	tstr := syh + zf.Testing(true) + syh + hhf
	buffer.WriteString(tstr)
	buffer.WriteString(hhf + xkhy)
}
func tianjiaxiugai(fangfa string, bianma string, buffer *bytes.Buffer) {
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	hhf := zfzhi.Huanhangfuzhi()
	kgf := zfzhi.Konggefuzhi()
	dh := zfzhi.Dianhaozhi()
	douhao := zfzhi.Douhaozhi()
	xh := zfzhi.Xinghaozhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	sz1 := zfzhi.Shuzi1zhi()
	sz1str := strconv.Itoa(sz1)
	sz0str := strconv.Itoa(zfzhi.Shuzi0zhi())
	syh := zfzhi.Shuangyinhaozhi()
	qh := zfzhi.Qiehaozhi()

	bmx := strings.ToLower(bianma)
	//func TestJueseserviceTianjia(t *testing.T)
	funcstr := hhf + zf.Func(true) + kgf + zf.Test(false) + bianma + zf.Services(true) + fangfa + xkhz + zf.T(true) + kgf + xh + zf.Testing(true) + dh + zf.T(false) + xkhy
	buffer.WriteString(funcstr)
	buffer.WriteString(dkhz + hhf)
	//juese := appmodels.Juese
	objstr := bmx + mh + dyh + zf.Appmodels(true) + dh + bianma
	buffer.WriteString(objstr)

	buffer.WriteString(dkhz + hhf)
	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		// Id : 1,
		if zf.Int(true) == Huoquyigeleixing(lk) {
			intstr := lk + mh + sz1str + douhao + hhf
			buffer.WriteString(intstr)
			continue
		}
		// Bianma :"",
		if zf.String(true) == Huoquyigeleixing(lk) {
			strstr := lk + mh + syh + bianma + fangfa + syh + douhao + hhf
			buffer.WriteString(strstr)
			continue
		}
		// time : timeNow()
		if zf.Time(true) == Huoquyigeleixing(lk) {
			timestr := lk + mh + zf.Time(true) + dh + zf.Now(false) + xkhz + xkhy + douhao + hhf
			buffer.WriteString(timestr)
			continue
		}
		// Qian : 1.0
		if zf.Float(true) == Huoquyigeleixing(lk) {
			flostr := lk + mh + sz1str + dh + sz0str + douhao + hhf
			buffer.WriteString(flostr)
			continue
		}
	}
	buffer.WriteString(hhf + dkhy)
	//zdxxxservices.Tianjiaxxx(&xxx)
	tjstr := hhf + zf.Zd(true) + bmx + zf.Services(true) + dh + fangfa + bmx + xkhz + qh + bmx + xkhy
	buffer.WriteString(tjstr)
	buffer.WriteString(hhf + dkhy + hhf)
}
func testservicetianjia(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	tianjiaxiugai(zf.Tianjia(false), bianma, buffer)
}
func testservicexiugai(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	tianjiaxiugai(zf.Xiugai(false), bianma, buffer)
}
func testservicechaxun(bianma string, buffer *bytes.Buffer) {
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	hhf := zfzhi.Huanhangfuzhi()
	kgf := zfzhi.Konggefuzhi()
	dh := zfzhi.Dianhaozhi()
	sz1str := strconv.Itoa(zfzhi.Shuzi1zhi())
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xh := zfzhi.Xinghaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	mh := zfzhi.Maohaozhi()

	bmx := strings.ToLower(bianma)
	// func TestXxxservicesChaxun
	funstr := zf.Func(true) + kgf + zf.Test(false) + bianma + zf.Services(true) + zf.Chaxun(false)
	buffer.WriteString(funstr)
	//(t *testing.T)
	tstr := xkhz + zf.T(true) + kgf + xh + zf.Testing(true) + dh + zf.T(false) + xkhy
	buffer.WriteString(tstr)
	buffer.WriteString(dkhz + hhf)
	// xxx := zdxxxservices.Chaxunxxx(1)
	objstr := bmx + mh + dyh + zf.Zd(true) + bmx + zf.Services(true) + dh + zf.Chaxun(false) + bmx + xkhz + sz1str + xkhy + hhf
	buffer.WriteString(objstr)
	logstr := zf.Log(true) + dh + zf.Println(false) + xkhz + bmx + xkhy + hhf
	buffer.WriteString(logstr)
	buffer.WriteString(hhf + dkhy + hhf)
}
func testserviceshanchu(bianma string, buffer *bytes.Buffer) {
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	hhf := zfzhi.Huanhangfuzhi()
	kgf := zfzhi.Konggefuzhi()
	dh := zfzhi.Dianhaozhi()
	sz1str := strconv.Itoa(zfzhi.Shuzi1zhi())
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xh := zfzhi.Xinghaozhi()

	bmx := strings.ToLower(bianma)
	// func TestXxxserviceShanchu
	funstr := zf.Func(true) + kgf + zf.Test(false) + bianma + zf.Services(true) + zf.Shanchu(false)
	buffer.WriteString(funstr)
	//(t *testing.T)
	tstr := xkhz + zf.T(true) + kgf + xh + zf.Testing(true) + dh + zf.T(false) + xkhy
	buffer.WriteString(tstr)
	buffer.WriteString(dkhz + hhf)
	//zdxxxservices.Shanchuxxx(1)
	scstr := zf.Zd(true) + bmx + zf.Services(true) + dh + zf.Shanchu(false) + bmx + xkhz + sz1str + xkhy
	buffer.WriteString(scstr)
	buffer.WriteString(hhf + dkhy + hhf)
}
func Shengchengservicetest() {
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	hhf := zfzhi.Huanhangfuzhi()
	kgf := zfzhi.Konggefuzhi()
	xx := zfzhi.Xiexianzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	dh := zfzhi.Dianhaozhi()

	for bk, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		pacstr := zf.Package(true) + kgf + zf.Tests(true) + hhf
		buffer.WriteString(pacstr)
		servicetestimport(bk, &buffer)
		testservicetianjia(bk, &buffer)
		testservicexiugai(bk, &buffer)
		testservicechaxun(bk, &buffer)
		testserviceshanchu(bk, &buffer)
		dir := basemodels.Getapppath() + xx + zf.Tests(true)
		path := dir + xx + bk + zf.Services(true) + xhx + zf.Test(true) + dh + zf.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}

}
