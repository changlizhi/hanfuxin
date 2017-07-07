package baserun

import (
	"bytes"
	"hanfuxin/appinits"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/basemodels"
	"io/ioutil"
	"os"
	"strings"
)

//判断是否有time格式的字段
func havetime(bianma string) bool {
	zf := zf.Zf{}
	for _, lie := range baseinits.Lies {
		pipei, _ := apputils.Pipei3lei(bianma, lie.Biaoming)
		if pipei && lie.Leixing == zf.Time(true) {
			return true
		}
	}
	return false
}

// tests的import
func testsimports(bianma string, buffer *bytes.Buffer) {
	// package tests \n
	zf := zfz.Zf{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	qh := zfzhi.Qiehaozhi()
	xh := zfzhi.Xinghaozhi()
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

	if havetime(bianma) {
		// "time" \n
		buffer.WriteString(syh)
		buffer.WriteString(zf.Time(true))
		buffer.WriteString(syh)
		buffer.WriteString(hhf)
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
	zf := zfz.Zf{}
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
	syh := zfzhi.Shuangyinhaozhi()
	xx := zfzhi.Xiexianzhi()
	xh := zfzhi.Xinghaozhi()
	sz1 := zfzhi.Shuzhi1zhi()

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
	zf := zfz.Zf{}
	tianjiahuoxiugai(bianma, buffer, zf.Tianjiayige(false))
}
func testxiugaiyige(bianma string, buffer *bytes.Buffer) {
	zf := zfz.Zf{}
	tianjiahuoxiugai(bianma, buffer, zf.Xiugaiyige(false))
}

func tianjiahuoxiugai(bianma string, buffer *bytes.Buffer, fangfa string) {
	//func TestxxxJuese
	zf := zfz.Zf{}
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
	syh := zfzhi.Shuangyinhaozhi()
	xx := zfzhi.Xiexianzhi()
	xh := zfzhi.Xinghaozhi()
	sz1 := zfzhi.Shuzhi1zhi()

	buffer.WriteString(zf.Func(true))
	buffer.WriteString(kgf)
	fangfaming := zf.Test(false) + fangfa + bianma
	buffer.WriteString(fangfaming)

	// (t *testing.T){\n
	buffer.WriteString(xkgz)
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
	mh := zfzhi.Maohaozhi()
	douhao := zfzhi.Douhaozhi()
	hhf := zfzhi.Huanghangfuzhi()
	for _, lie := range baseinits.Lies {
		pipei, _ := apputils.Pipei3lei(bianma, lie.Biaoming)
		if pipei {
			buffer.WriteString(lie.Bianma)                                 //Juese
			buffer.WriteString(mh)                                         // :
			zhi := shengchengzhi(lie.Bianma, lie.Leixing, fangfa, houzhui) //zhi
			buffer.WriteString(zhi)
			buffer.WriteString(douhao) //,
			buffer.WriteString(hhf)    //\n
		}
	}
}
func shengchengzhi(lieming string, leixing string, fangfa string, houzhui string) string {
	zf := zfz.Zf{}
	syh := zfzhi.Shuangyinhaozhi()
	dh := zfzhi.Dianhao()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()

	sz1 := zfzhi.Shuzi1zhi()
	sz0 := zfzhi.Shuzi0zhi()

	if leixing == baseinits.Gen.String {
		// "lieZengjiaTest1"
		ret := syh + lieming + fangfa + zf.Test(false) + houzhui + syh
		return ret
	}
	// 1
	if leixing == baseinits.Gen.Int {
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

	zf := zfz.Zf{}
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
	syh := zfzhi.Shuangyinhaozhi()
	xx := zfzhi.Xiexianzhi()
	xh := zfzhi.Xinghaozhi()

	// func TestShanchuyige
	buffer.WriteString(baseinits.Gen.Func) //func
	buffer.WriteString(kgf)                // konggefu
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
	zf := zfz.Zf{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	mh := zfzhi.Maohaozhi()
	qh := zfzhi.Qiehaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	xh := zfzhi.Xinghaozhi()
	syh := zfzhi.Shuangyinhaozhi()
	xx := zfzhi.Xiexianzhi()
	xh := zfzhi.Xinghaozhi()
	sz2 := zfzhi.Shuzi2()
	sz3 := zfzhi.Shuzi3()
	//组装方法名
	// func TestTianjiaduoge
	buffer.WriteString(baseinits.Gen.Func)
	buffer.WriteString(kgf)
	buffer.WriteString(zf.Test(false))
	buffer.WriteString(zf.Tianjiaduoge(false))

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
	zf := zfz.Zf{}
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
	syh := zfzhi.Shuangyinhaozhi()
	xx := zfzhi.Xiexianzhi()
	xh := zfzhi.Xinghaozhi()
	xhx := zfzhi.Xiahuaxianzhi()
	for _, biao := range baseinits.Biaos {
		buffer := bytes.Buffer{}
		testsimports(biao.Bianma, &buffer)
		testchaxunyige(biao.Bianma, &buffer)
		testshanchuyige(biao.Bianma, &buffer)
		testtianjiayige(biao.Bianma, &buffer)
		testxiugaiyige(biao.Bianma, &buffer)
		testtianjiaduoge(biao.Bianma, &buffer)

		path := basemodels.Getapppath() + xx + zf.Tests(true) + xx + biao.Bianma + zf.Daos(true) + xhx + zf.Test(true) + dh + zf.Go(true)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
