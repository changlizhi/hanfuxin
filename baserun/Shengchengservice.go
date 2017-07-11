package baserun

import (
	"bytes"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strings"
)

func serviceimports(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	syh := zfzhi.Shuangyinhaozhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	xx := zfzhi.Xiexianzhi()
	hhf := zfzhi.Huanhangfuzhi()

	importstr := zf.Import(true) + xkhz + hhf
	buffer.WriteString(importstr)
	//"hanfuxin/allerrors" \n
	errorbao := syh + zf.Hanfuxin(true) + xx + zf.Allerrors(true) + syh + hhf
	buffer.WriteString(errorbao)
	// "bytes"
	bytesbao := syh + zf.Bytes(true) + syh + hhf
	buffer.WriteString(bytesbao)
	// "log"
	logbao := syh + zf.Log(true) + syh + hhf
	buffer.WriteString(logbao)
	// "time"
	timebao := syh + zf.Time(true) + syh + hhf
	buffer.WriteString(timebao)
	// "hanfuxin/appmodels"
	apmbao := syh + zf.Hanfuxin(true) + xx + zf.Appmodels(true) + syh + hhf
	buffer.WriteString(apmbao)
	// "hanfuxin/apputils"
	apubao := syh + zf.Hanfuxin(true) + xx + zf.Apputils(true) + syh + hhf
	buffer.WriteString(apubao)
	// "hanfuxin/baserun"
	brbao := syh + zf.Hanfuxin(true) + xx + zf.Baserun(true) + syh + hhf
	buffer.WriteString(brbao)
	// "hanfuxin/zf"
	zfbao := syh + zf.Hanfuxin(true) + xx + zf.Zf(true) + syh + hhf
	buffer.WriteString(zfbao)
	// "hanfuxin/zfzhi"
	zfzhibao := syh + zf.Hanfuxin(true) + xx + zf.Zfzhi(true) + syh + hhf
	buffer.WriteString(zfzhibao)
	// "hanfuxin/zdxxxdaos"
	daobao := syh + zf.Hanfuxin(true) + xx + zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true) + syh + hhf + xkhy + hhf
	buffer.WriteString(daobao)
}
func yanzhengchangdu(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xh := zfzhi.Xinghaozhi()
	hhf := zfzhi.Huanhangfuzhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	dh := zfzhi.Dianhaozhi()

	bmx := strings.ToLower(bianma)

	kgf := zfzhi.Konggefuzhi()
	funcone := zf.Func(true) + kgf + zf.Yanzhengziduanchangdu(true)
	buffer.WriteString(funcone)
	//(bmx *appmodels.bm) error { \n
	canshu := xkhz + bmx + kgf + xh + zf.Appmodels(true) + bianma + xkhy + zf.Error(true) + dkhz + hhf
	buffer.WriteString(canshu)
	// zf := zf.Zf{}
	zfobj := zf.Zf(true) + mh + dyh + zf.Zf(true) + dh + zf.Zf(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfobj)
	cuowu := zf.Cuowu(true) + mh + dyh + zf.False(true) + hhf
	buffer.WriteString(cuowu)

	for lk, lv := range Huoquyigebiaojiegou(bianma) {
		//lenbianma :=
		//baserun.Huoquyigechangdu
		//(zf.Bianma(false))
		lenone := zf.Len(true) + lv + mh + dyh +
			zf.Baserun(true) + dh + zf.Huoquyigechangdu(false) +
			xkhz + zf.Zf(true) + dh + lk + xkhz + zf.False(true) + xkhy + xkhy + hhf
		buffer.WriteString(lenone)
	}

	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)

}
func Shengchengservice() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()
	for bk, bv := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		bm := zf.Zd(true) + bv + zf.Services(true)
		//package xxxservices \n
		pac := zf.Package(true) + kgf + bm + hhf
		buffer.WriteString(pac)

		serviceimports(bk, &buffer)
		yanzhengchangdu(bk, &buffer)
		log.Println("buffer--------", buffer.String())
	}
}
