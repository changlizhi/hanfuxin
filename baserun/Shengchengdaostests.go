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

	buffer.WriteString(zf.Package(true))
	buffer.WriteString(kgf)
	buffer.WriteString(zf.Tests(true))
	buffer.WriteString(hhf)

	// import(\n
	buffer.WriteString(zf.Import(true))
	buffer.WriteString(xkhz)
	buffer.WriteString(hhf)

	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		if Huoquyigeleixing(lk) == zf.Time(true) {
			// "time" \n
			buffer.WriteString(syh)
			buffer.WriteString(zf.Time(true))
			buffer.WriteString(syh)
			buffer.WriteString(hhf)
			break
		}
	}
	// "hanfuxin/appmodels" \n
	buffer.WriteString(syh)
	buffer.WriteString(zf.Hanfuxin(true))
	buffer.WriteString(xx)
	buffer.WriteString(zf.Appmodels(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)

	// "hanfuxin/testing" \n
	buffer.WriteString(syh)
	buffer.WriteString(zf.Testing(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)

	// "log" \n
	buffer.WriteString(syh)
	buffer.WriteString(zf.Log(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)

	// "hanfuxin/zdjuesedaos"
	buffer.WriteString(syh)
	buffer.WriteString(zf.Hanfuxin(true))
	buffer.WriteString(xx)
	baoming := zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true) // daos
	buffer.WriteString(baoming)                                      // zdjuesedaos
	buffer.WriteString(syh)                                          // "
	buffer.WriteString(hhf)                                          // \n

	buffer.WriteString(xkhy) // )
	buffer.WriteString(hhf)  // \n
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

	buffer.WriteString(zf.Func(true)) // func
	buffer.WriteString(kgf)
	fangfaming := zf.Test(false) + zf.Chaxunyige(false) + bianma
	buffer.WriteString(fangfaming) //TestChaxunyigeJuese

	// (t *testing.T){\n
	buffer.WriteString(xkhz)             // (
	buffer.WriteString(zf.T(true))       //t
	buffer.WriteString(kgf)              // konggefu
	buffer.WriteString(xh)               // *
	buffer.WriteString(zf.Testing(true)) // testing
	buffer.WriteString(dh)               // .
	buffer.WriteString(zf.T(false))      // T
	buffer.WriteString(xkhy)             // (
	buffer.WriteString(dkhz)             // {
	buffer.WriteString(hhf)              // \n

	// juese := zdjuesedaos.Chaxunyige
	buffer.WriteString(strings.ToLower(bianma)) //juese
	buffer.WriteString(mh)                      //:
	buffer.WriteString(dyh)                     //=
	baoming := zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true)
	buffer.WriteString(baoming)              //zidongjuesedaos
	buffer.WriteString(dh)                   //.
	buffer.WriteString(zf.Chaxunyige(false)) //Juese

	// (1)
	buffer.WriteString(xkhz)
	buffer.WriteString(sz1)
	buffer.WriteString(xkhy)
	buffer.WriteString(hhf)

	//log.Println(juese)
	buffer.WriteString(zf.Log(true))
	buffer.WriteString(dh)
	buffer.WriteString(zf.Println(false))
	buffer.WriteString(xkhz)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(xkhy)

	// \n } \n
	buffer.WriteString(hhf)
	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)
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

	buffer.WriteString(zf.Func(true))
	buffer.WriteString(kgf)
	fangfaming := zf.Test(false) + fangfa + bianma
	buffer.WriteString(fangfaming)

	// (t *testing.T){\n
	buffer.WriteString(xkhz)
	buffer.WriteString(zf.T(true))
	buffer.WriteString(kgf)
	buffer.WriteString(xh)
	buffer.WriteString(zf.Testing(true))
	buffer.WriteString(dh)
	buffer.WriteString(zf.T(false))
	buffer.WriteString(xkhy)
	buffer.WriteString(dkhz)
	buffer.WriteString(hhf)

	// juese := &appmodels.Juese{\n
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(mh)
	buffer.WriteString(dyh)
	buffer.WriteString(qh)
	buffer.WriteString(zf.Appmodels(true))
	buffer.WriteString(dh)
	buffer.WriteString(bianma)
	buffer.WriteString(dkhz)
	buffer.WriteString(hhf)

	// 生成字段和值
	pinjieziduan(bianma, buffer, fangfa, sz1)

	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)

	// zdjuesedaos.TestxxxJuese(juese)
	baoming := zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true)
	buffer.WriteString(baoming) // zidongjuesedaos
	buffer.WriteString(dh)      // .
	buffer.WriteString(fangfa)  // TestxxxJuese

	//(juese)
	buffer.WriteString(xkhz)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(xkhy)

	// \n } \n
	buffer.WriteString(hhf)
	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)
}
func pinjieziduan(bianma string, buffer *bytes.Buffer, fangfa string, houzhui string) {
	zfzhi := zfzhi.Zfzhi{}
	mh := zfzhi.Maohaozhi()
	douhao := zfzhi.Douhaozhi()
	hhf := zfzhi.Huanhangfuzhi()

	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		buffer.WriteString(lk)                                          //Juese
		buffer.WriteString(mh)                                          // :
		zhi := shengchengzhi(lk, Huoquyigeleixing(lk), fangfa, houzhui) //zhi
		buffer.WriteString(zhi)
		buffer.WriteString(douhao) //,
		buffer.WriteString(hhf)    //\n
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
	buffer.WriteString(zf.Func(true)) //func
	buffer.WriteString(kgf)           // konggefu
	fangfaming := zf.Test(false) + zf.Shanchuyige(false) + bianma
	buffer.WriteString(fangfaming) // TestShanchuyige

	// (t *testing.T){\n
	buffer.WriteString(xkhz)
	buffer.WriteString(zf.T(true))
	buffer.WriteString(kgf)
	buffer.WriteString(xh)
	buffer.WriteString(zf.Testing(true))
	buffer.WriteString(dh)
	buffer.WriteString(zf.T(false))
	buffer.WriteString(xkhy)
	buffer.WriteString(dkhz)
	buffer.WriteString(hhf)

	// zdjuesedaos.Shanchuyige(1)
	baoming := zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true)
	buffer.WriteString(baoming)               //zdjuesedaos
	buffer.WriteString(dh)                    //.
	buffer.WriteString(zf.Shanchuyige(false)) //Shanchuyige
	buffer.WriteString(xkhz)                  //(
	buffer.WriteString(sz1)                   //1
	buffer.WriteString(xkhy)                  //)

	// \n } \n
	buffer.WriteString(hhf)
	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)
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
	buffer.WriteString(zf.Func(true))
	buffer.WriteString(kgf)
	buffer.WriteString(zf.Test(false))
	buffer.WriteString(zf.Tianjiaduoge(false))
	buffer.WriteString(bianma)

	//(t *testing.T){\n
	buffer.WriteString(xkhz)
	buffer.WriteString(zf.T(true))
	buffer.WriteString(kgf)
	buffer.WriteString(xh)
	buffer.WriteString(zf.Testing(true))
	buffer.WriteString(dh)
	buffer.WriteString(zf.T(false))
	buffer.WriteString(xkhy)
	buffer.WriteString(dkhz)
	buffer.WriteString(hhf)

	// 组装第一个实体
	//juese2 := appmodels.Juese

	shiti2 := strings.ToLower(bianma) + sz2
	buffer.WriteString(shiti2)             //juese2
	buffer.WriteString(mh)                 // :
	buffer.WriteString(dyh)                // =
	buffer.WriteString(zf.Appmodels(true)) //appmodels
	buffer.WriteString(dh)                 // .
	buffer.WriteString(bianma)             //Juese

	//{ \n zhi... \n }
	buffer.WriteString(dkhz)
	buffer.WriteString(hhf)
	pinjieziduan(bianma, buffer, zf.Tianjiaduoge(false), sz2)
	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)

	// juese3 := appmodels.Juese{\n
	shiti3 := strings.ToLower(bianma) + sz3
	buffer.WriteString(shiti3)
	buffer.WriteString(mh)
	buffer.WriteString(dyh)
	buffer.WriteString(zf.Appmodels(true))
	buffer.WriteString(dh)
	buffer.WriteString(bianma)
	buffer.WriteString(dkhz)
	buffer.WriteString(hhf)

	// \n zhi \n
	pinjieziduan(bianma, buffer, zf.Tianjiaduoge(false), sz3)
	buffer.WriteString(dkhy) //}
	buffer.WriteString(hhf)

	// zdjuesedaos
	baoming := zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true)
	// jueses :=
	duoge := strings.ToLower(bianma) + zf.S(true)
	buffer.WriteString(duoge)
	buffer.WriteString(mh)
	buffer.WriteString(dyh)

	//[]appmodes.Juese
	buffer.WriteString(zkhz)
	buffer.WriteString(zkhy)
	buffer.WriteString(zf.Appmodels(true))
	buffer.WriteString(dh)
	buffer.WriteString(bianma)

	//{juese1,juese2}
	buffer.WriteString(dkhz)
	buffer.WriteString(shiti2)
	buffer.WriteString(douhao)
	buffer.WriteString(shiti3)
	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)

	//zdjuesedaso.Tianjiaduoge
	buffer.WriteString(baoming)
	buffer.WriteString(dh)
	buffer.WriteString(zf.Tianjiaduoge(false))

	//(jueses)
	buffer.WriteString(xkhz)
	buffer.WriteString(duoge)
	buffer.WriteString(xkhy)

	//\n } \n
	buffer.WriteString(hhf)
	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)
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
