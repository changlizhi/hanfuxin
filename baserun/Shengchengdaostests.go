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

// tests的import
func testsimports(bianma string, buffer *bytes.Buffer) {
	// package tests \n
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	kgf := zfzhi.Konggefuzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	syh := zfzhi.Shuangyinhaozhi()
	xx := zfzhi.Xiexianzhi()

	pac := zf.Package(true) + kgf + zf.Tests(true) + hhf
	buffer.WriteString(pac)

	// import(\n
	impstr := zf.Import(true) + xkhz + hhf
	buffer.WriteString(impstr)

	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		if Huoquyigeleixing(lk) == zf.Time(true) {
			// "time" \n
			timebao := syh + zf.Time(true) + syh + hhf
			buffer.WriteString(timebao)
			break
		}
	}
	// "hanfuxin/appmodels" \n
	apm := syh + zf.Hanfuxin(true) + xx + zf.Appmodels(true) + syh + hhf
	buffer.WriteString(apm)

	// "hanfuxin/testing" \n
	tbao := syh + zf.Testing(true) + syh + hhf
	buffer.WriteString(tbao)

	// "log" \n
	logbao := syh + zf.Log(true) + syh + hhf
	buffer.WriteString(logbao)

	// "hanfuxin/zdjuesedaos"
	daobao := syh + zf.Hanfuxin(true) + xx + zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true) + syh + hhf
	buffer.WriteString(daobao)

	buffer.WriteString(xkhy + hhf)
}

func testchaxunyige(bianma string, buffer *bytes.Buffer) {
	// func TestChaxunyigeJuese
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	xh := zfzhi.Xinghaozhi()
	sz1 := strconv.Itoa(zfzhi.Shuzi1zhi())

	// func TestChaxunyigeXXX
	funstr := zf.Func(true) + kgf + zf.Test(false) + zf.Chaxunyige(false) + bianma
	buffer.WriteString(funstr)

	//(t*testing.T){\n
	csstr := xkhz + zf.T(true) + kgf + xh + zf.Testing(true) + dh + zf.T(false) + xkhy + dkhz + hhf
	buffer.WriteString(csstr)

	//juese:=
	//zdjuesedaos.Chaxunyige(1)
	fhstr := strings.ToLower(bianma) + mh + dyh +
		zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true) + dh + zf.Chaxunyige(false) + xkhz + sz1 + xkhy + hhf
	buffer.WriteString(fhstr)

	//log.Println(juese+
	dayin := zf.Log(true) + dh + zf.Println(false) + xkhz + strings.ToLower(bianma) + xkhy + hhf
	buffer.WriteString(dayin)

	// \n } \n
	buffer.WriteString(dkhy + hhf)
}
func testtianjiayige(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	tianjiahuoxiugai(bianma, buffer, zf.Tianjiayige(false))
}
func testxiugaiyige(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	tianjiahuoxiugai(bianma, buffer, zf.Xiugaiyige(false))
}

func tianjiahuoxiugai(bianma string, buffer *bytes.Buffer, fangfa string) {
	//func TestxxxJuese
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	mh := zfzhi.Maohaozhi()
	qh := zfzhi.Qiehaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	xh := zfzhi.Xinghaozhi()
	sz1 := strconv.Itoa(zfzhi.Shuzi1zhi())

	funstr := zf.Func(true) + kgf + zf.Test(false) + fangfa + bianma
	buffer.WriteString(funstr)

	// (t *testing.T){\n
	csstr := xkhz + zf.T(true) + kgf + xh + zf.Testing(true) + dh + zf.T(false) + xkhy + dkhz + hhf
	buffer.WriteString(csstr)

	// juese := &appmodels.Juese{\n
	dxstr := strings.ToLower(bianma) + mh + dyh + qh + zf.Appmodels(true) + dh + bianma + dkhz + hhf
	buffer.WriteString(dxstr)

	// 生成字段和值
	pinjieziduan(bianma, buffer, fangfa, sz1)
	buffer.WriteString(dkhy + hhf)

	// zdjuesedaos.TestxxxJuese(juese)
	baoming := zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true) + dh + fangfa + xkhz + strings.ToLower(bianma) + xkhy + hhf
	buffer.WriteString(baoming)

	// \n } \n
	buffer.WriteString(dkhy + hhf)
}
func pinjieziduan(bianma string, buffer *bytes.Buffer, fangfa string, houzhui string) {
	zfzhi := zfzhi.Zfzhi{}
	mh := zfzhi.Maohaozhi()
	douhao := zfzhi.Douhaozhi()
	hhf := zfzhi.Huanhangfuzhi()

	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		buffer.WriteString(lk + mh)
		zhi := shengchengzhi(lk, Huoquyigeleixing(lk), fangfa, houzhui)
		buffer.WriteString(zhi + douhao + hhf)
	}

}
func shengchengzhi(lieming string, leixing string, fangfa string, houzhui string) string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	syh := zfzhi.Shuangyinhaozhi()
	dh := zfzhi.Dianhaozhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()

	sz1 := strconv.Itoa(zfzhi.Shuzi1zhi())
	sz0 := strconv.Itoa(zfzhi.Shuzi0zhi())

	if leixing == zf.String(true) {
		// "lieZengjiaTest1"
		ret := syh + lieming + fangfa + zf.Test(false) + houzhui + syh
		return ret
	}
	// 1
	if leixing == zf.Int(true) {
		ret := sz1 // 1
		return ret
	}
	// time.Now()
	if leixing == zf.Time(true) {
		// time.Now()
		ret := zf.Time(true) + dh + zf.Now(false) + xkhz + xkhy
		return ret
	}
	// 1.0
	if leixing == zf.Float32(true) || leixing == zf.Float64(true) {
		// 1.0
		ret := sz1 + dh + sz0
		return ret
	}
	// Null
	return zf.Null(false)
}

func testshanchuyige(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	xh := zfzhi.Xinghaozhi()
	sz1 := strconv.Itoa(zfzhi.Shuzi1zhi())

	// func TestShanchuyige
	funstr := zf.Func(true) + kgf + zf.Test(false) + zf.Shanchuyige(false) + bianma
	buffer.WriteString(funstr)

	// (t *testing.T){\n
	csstr := xkhz + zf.T(true) + kgf + xh + zf.Testing(true) + dh + zf.T(false) + xkhy + dkhz + hhf
	buffer.WriteString(csstr)

	// zdjuesedaos.Shanchuyige(1)
	baoming := zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true) + dh + zf.Shanchuyige(false) + xkhz + sz1 + xkhy + hhf
	buffer.WriteString(baoming)
	// \n } \n
	buffer.WriteString(dkhy + hhf)
}
func testtianjiaduoge(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	dyh := zfzhi.Dengyuhaozhi()
	mh := zfzhi.Maohaozhi()
	xh := zfzhi.Xinghaozhi()
	//syh := zfzhi.Shuangyinhaozhi()
	sz2 := strconv.Itoa(zfzhi.Shuzi2zhi())
	sz3 := strconv.Itoa(zfzhi.Shuzi3zhi())
	douhao := zfzhi.Douhaozhi()
	//组装方法名
	// func TestTianjiaduogeJuese
	funstr := zf.Func(true) + kgf + zf.Test(false) + zf.Tianjiaduoge(false) + bianma
	buffer.WriteString(funstr)

	//(t *testing.T){\n
	csstr := xkhz + zf.T(true) + kgf + xh + zf.Testing(true) + dh + zf.T(false) + xkhy + dkhz + hhf
	buffer.WriteString(csstr)

	// 组装第一个实体
	//juese2 := appmodels.Juese
	shiti2 := strings.ToLower(bianma) + sz2 + mh + dyh + zf.Appmodels(true) + dh + bianma
	buffer.WriteString(shiti2)

	//{ \n zhi... \n }
	buffer.WriteString(dkhz + hhf)
	pinjieziduan(bianma, buffer, zf.Tianjiaduoge(false), sz2)
	buffer.WriteString(dkhy + hhf)

	// juese3 := appmodels.Juese{\n
	shiti3 := strings.ToLower(bianma) + sz3 + mh + dyh + zf.Appmodels(true) + dh + bianma + dkhz + hhf
	buffer.WriteString(shiti3)

	// \n zhi \n
	pinjieziduan(bianma, buffer, zf.Tianjiaduoge(false), sz3)
	buffer.WriteString(dkhy + hhf)

	// jueses :=[]appmodes.Juese
	duoge := strings.ToLower(bianma) + zf.S(true) + mh + dyh + zkhz + zkhy + zf.Appmodels(true) + dh + bianma
	buffer.WriteString(duoge)

	//{juese1,juese2}
	szs := dkhz + strings.ToLower(bianma) + sz2 + douhao + strings.ToLower(bianma) + sz3 + dkhy + hhf
	buffer.WriteString(szs)

	//zdjuesedaos.Tianjiaduoge(jueses)
	baoming := zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true) + dh + zf.Tianjiaduoge(false) + xkhz + strings.ToLower(bianma) + zf.S(true) + xkhy + hhf
	buffer.WriteString(baoming)

	//} \n
	buffer.WriteString(dkhy + hhf)
}
func Shengchengdaostests() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	dh := zfzhi.Dianhaozhi()
	xx := zfzhi.Xiexianzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	for biao, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		testsimports(biao, &buffer)
		testchaxunyige(biao, &buffer)
		testshanchuyige(biao, &buffer)
		testtianjiayige(biao, &buffer)
		testxiugaiyige(biao, &buffer)
		testtianjiaduoge(biao, &buffer)
		log.Println("buffer-------", buffer.String())

		path := basemodels.Getapppath() + xx + zf.Tests(true) + xx + biao + zf.Daos(true) + xhx + zf.Test(true) + dh + zf.Go(true)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
