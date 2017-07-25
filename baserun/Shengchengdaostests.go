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

// tests的import
func testsimports(bianma string, buffer *bytes.Buffer) {
	pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf()
	buffer.WriteString(pac)

	// import(\n
	impstr := zf.Zfs.Import(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(impstr)

	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		if Huoquyigeleixing(lk) == zf.Zfs.Time(true) {
			// "time" \n
			timebao := zfzhi.Zhi.Syh() + zf.Zfs.Time(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
			buffer.WriteString(timebao)
			break
		}
	}
	// "hanfuxin/appmodels" \n
	apm := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(apm)

	// "hanfuxin/testing" \n
	tbao := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tbao)

	// "log" \n
	logbao := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logbao)

	// "hanfuxin/zdjuesedaos"
	daobao := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Daos(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(daobao)

	buffer.WriteString(zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
}

func testchaxunyige(bianma string, buffer *bytes.Buffer) {
	sz1 := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())

	// func TestChaxunyigeXXX
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + zf.Zfs.Chaxunyige(false) + bianma
	buffer.WriteString(funstr)

	//(t*testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	//juese:=
	//zdjuesedaos.Chaxunyige(1)
	fhstr := strings.ToLower(bianma) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Daos(true) + zfzhi.Zhi.Dh() + zf.Zfs.Chaxunyige(false) + zfzhi.Zhi.Xkhz() + sz1 + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fhstr)

	//log.Println(juese+
	dayin := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) + zfzhi.Zhi.Xkhz() + strings.ToLower(bianma) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(dayin)

	// \n } \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testtianjiayige(bianma string, buffer *bytes.Buffer) {
	tianjiahuoxiugai(bianma, buffer, zf.Zfs.Tianjiayige(false))
}
func testxiugaiyige(bianma string, buffer *bytes.Buffer) {
	tianjiahuoxiugai(bianma, buffer, zf.Zfs.Xiugaiyige(false))
}

func tianjiahuoxiugai(bianma string, buffer *bytes.Buffer, fangfa string) {
	sz1 := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())

	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + fangfa + bianma
	buffer.WriteString(funstr)

	// (t *testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	// juese := &appmodels.Juese{\n
	dxstr := strings.ToLower(bianma) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Qh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(dxstr)

	// 生成字段和值
	pinjieziduan(bianma, buffer, fangfa, sz1)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	// zdjuesedaos.Testzfzhi.Zhi.Xx()xJuese(juese)
	baoming := zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Daos(true) + zfzhi.Zhi.Dh() + fangfa + zfzhi.Zhi.Xkhz() + strings.ToLower(bianma) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(baoming)

	// \n } \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func pinjieziduan(bianma string, buffer *bytes.Buffer, fangfa string, houzhui string) {

	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		buffer.WriteString(lk + zfzhi.Zhi.Mh())
		zhi := shengchengzhi(lk, Huoquyigeleixing(lk), fangfa, houzhui)
		buffer.WriteString(zhi + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())
	}

}
func shengchengzhi(lieming string, leixing string, fangfa string, houzhui string) string {
	sz1 := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())
	sz0 := strconv.Itoa(zfzhi.Zhi.Shuzi0zhi())

	if leixing == zf.Zfs.String(true) {
		// "lieZengjiaTest1"
		ret := zfzhi.Zhi.Syh() + lieming + fangfa + zf.Zfs.Test(false) + houzhui + zfzhi.Zhi.Syh()
		return ret
	}
	// 1
	if leixing == zf.Zfs.Int(true) {
		ret := sz1 // 1
		return ret
	}
	// time.Now()
	if leixing == zf.Zfs.Time(true) {
		// time.Now()
		ret := zf.Zfs.Time(true) + zfzhi.Zhi.Dh() + zf.Zfs.Now(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
		return ret
	}
	// 1.0
	if leixing == zf.Zfs.Float32(true) || leixing == zf.Zfs.Float64(true) {
		// 1.0
		ret := sz1 + zfzhi.Zhi.Dh() + sz0
		return ret
	}
	// Null
	return zf.Zfs.Null(false)
}

func testshanchuyige(bianma string, buffer *bytes.Buffer) {
	sz1 := strconv.Itoa(zfzhi.Zhi.Shuzi1zhi())

	// func TestShanchuyige
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + zf.Zfs.Shanchuyige(false) + bianma
	buffer.WriteString(funstr)

	// (t *testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	// zdjuesedaos.Shanchuyige(1)
	baoming := zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Daos(true) + zfzhi.Zhi.Dh() + zf.Zfs.Shanchuyige(false) + zfzhi.Zhi.Xkhz() + sz1 + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(baoming)
	// \n } \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testtianjiaduoge(bianma string, buffer *bytes.Buffer) {
	//zfzhi.Zhi.Syh() := zfzhi.Shuangyinhaozhi()
	sz2 := strconv.Itoa(zfzhi.Zhi.Shuzi2zhi())
	sz3 := strconv.Itoa(zfzhi.Zhi.Shuzi3zhi())
	//组装方法名
	// func TestTianjiaduogeJuese
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + zf.Zfs.Tianjiaduoge(false) + bianma
	buffer.WriteString(funstr)

	//(t *testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	// 组装第一个实体
	//juese2 := appmodels.Juese
	shiti2 := strings.ToLower(bianma) + sz2 + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(shiti2)

	//{ \n zhi... \n }
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	pinjieziduan(bianma, buffer, zf.Zfs.Tianjiaduoge(false), sz2)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	// juese3 := appmodels.Juese{\n
	shiti3 := strings.ToLower(bianma) + sz3 + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(shiti3)

	// \n zhi \n
	pinjieziduan(bianma, buffer, zf.Zfs.Tianjiaduoge(false), sz3)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	// jueses :=[]appmodes.Juese
	duoge := strings.ToLower(bianma) + zf.Zfs.S(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(duoge)

	//{juese1,juese2}
	szs := zfzhi.Zhi.Dkhz() + strings.ToLower(bianma) + sz2 + zfzhi.Zhi.Dou() + strings.ToLower(bianma) + sz3 + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(szs)

	//zdjuesedaos.Tianjiaduoge(jueses)
	baoming := zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Daos(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tianjiaduoge(false) + zfzhi.Zhi.Xkhz() + strings.ToLower(bianma) + zf.Zfs.S(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(baoming)

	//} \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func Shengchengdaostests() {
	for biao, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		testsimports(biao, &buffer)
		testtianjiaduoge(biao, &buffer)
		testtianjiayige(biao, &buffer)
		testxiugaiyige(biao, &buffer)
		testchaxunyige(biao, &buffer)
		testshanchuyige(biao, &buffer)

		path := basemodels.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true) + zfzhi.Zhi.Xx() + biao + zf.Zfs.Daos(true) + zfzhi.Zhi.Xhx() + zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
