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
	daobao := syh + zf.Hanfuxin(true) + xx + zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true) + syh + hhf
	buffer.WriteString(daobao)
	//"hanfuxin/baseinits"
	binitbao := syh + zf.Hanfuxin(true) + xx + zf.Baseinits(true) + syh
	buffer.WriteString(binitbao)
	buffer.WriteString(xkhy + hhf)
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
	dayuhao := zfzhi.Dayuhaozhi()
	douhao := zfzhi.Douhaozhi()

	bmx := strings.ToLower(bianma)

	kgf := zfzhi.Konggefuzhi()
	funcone := zf.Func(true) + kgf + zf.Yanzhengziduanchangdu(true)
	buffer.WriteString(funcone)
	//(bmx *appmodels.bm) error { \n
	canshu := xkhz + bmx + kgf + xh + zf.Appmodels(true) + dh + bianma + xkhy + zf.Error(true) + dkhz + hhf

	buffer.WriteString(canshu)
	// zf := zf.Zf{}
	zfobj := zf.Zf(true) + mh + dyh + zf.Zf(true) + dh + zf.Zf(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfobj)
	cuowu := zf.Cuowu(true) + mh + dyh + zf.False(true) + hhf
	buffer.WriteString(cuowu)

	cuowutrue := zf.Cuowu(true) + dyh + zf.True(true) + hhf
	//buffer := bytes.Buffer{}\n
	bf := zf.Buffer(true) + mh + dyh + zf.Bytes(true) + dh + zf.Buffer(false) + dkhz + dkhy + hhf
	buffer.WriteString(bf)
	for lk, lv := range Huoquyigebiaojiegou(bianma) {
		if Huoquyigeleixing(lk) == zf.String(true) {
			buffer.WriteString(hhf)
			lenlvstr := zf.Len(true) + lv
			lenlvshitistr := lenlvstr + zf.Shiti(true)
			//lenbianma :=
			//baserun.Huoquyigechangdu
			//(zf.Bianma(false))

			lenlv := lenlvstr + mh + dyh +
				zf.Baserun(true) + dh + zf.Huoquyigechangdu(false) +
				xkhz + zf.Zf(true) + dh + lk + xkhz + zf.False(true) + xkhy + xkhy + hhf
			buffer.WriteString(lenlv)

			//	lenbianmashiti := len(juese.Bianma)
			lenshiti := lenlvshitistr + mh + dyh + zf.Len(true) + xkhz + bmx + dh + lk + xkhy + hhf
			buffer.WriteString(lenshiti)

			// if leibianmashiti > int(lenbianma) {\n
			oneif := zf.If(true) + kgf + lenlvshitistr + dayuhao + zf.Int(true) + xkhz + lenlvstr + xkhy + dkhz + hhf
			buffer.WriteString(oneif)
			buffer.WriteString(cuowutrue)
			//buffer.WriteString(....)
			writestr := zf.Buffer(true) + dh + zf.WriteString(false)
			buffer.WriteString(writestr)
			scstr := xkhz + zf.Apputils(true) + dh + zf.Shengchengerrorchangdu(false)
			buffer.WriteString(scstr)
			csstrs := xkhz + zf.Zf(true) + dh + zf.Mingcheng(false) + xkhz + zf.False(true) + xkhy
			buffer.WriteString(csstrs)
			int64str := douhao + zf.Int64(true) + xkhz + lenlvstr + xkhy
			buffer.WriteString(int64str)
			shitis := douhao + lenlvshitistr + xkhy + xkhy + hhf
			buffer.WriteString(shitis)
			buffer.WriteString(dkhy + hhf)
		}
	}
	// if cuowu {\n
	cwstr := zf.If(true) + kgf + zf.Cuowu(true) + dkhz + hhf
	buffer.WriteString(cwstr)
	// return allerrors.Ziduanerror
	retstr := zf.Return(true) + kgf + zf.Allerrors(true) + dh + zf.Ziduanerror(false)
	buffer.WriteString(retstr)
	//{Shijian:time.Now(),Wenti:buffer.String()}
	wtsjstr := dkhz + zf.Shijian(false) + mh + zf.Time(true) + dh + zf.Now(false) + xkhz + xkhy
	buffer.WriteString(wtsjstr)
	wtwtstr := douhao + zf.Wenti(false) + mh + zf.Buffer(true) + dh + zf.String(false) + xkhz + xkhy + dkhy + hhf
	buffer.WriteString(wtwtstr)
	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)

	nilstr := zf.Return(true) + kgf + zf.Nil(true) + hhf
	buffer.WriteString(nilstr)

	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)
}
func servicexiugai(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	xh := zfzhi.Xinghaozhi()
	dh := zfzhi.Dianhaozhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	hhf := zfzhi.Huanhangfuzhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	gth := zfzhi.Gantanhaozhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()

	bmx := strings.ToLower(bianma)

	funcstr := zf.Func(true) + kgf + zf.Xiugai(false) + bmx + xkhz + bmx + kgf + xh + zf.Appmodels(true) + dh + bianma + xkhy + zf.String(true)
	buffer.WriteString(funcstr)
	buffer.WriteString(dkhz + hhf)

	errstr := zf.Err(true) + mh + dyh + zf.Yanzhengziduanchangdu(true) + xkhz + bmx + xkhy + hhf
	buffer.WriteString(errstr)

	// zfzhi := zfzhi.Zfzhi{}
	zfzhistr := zf.Zfzhi(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Zfzhi(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfzhistr)
	//zf := zf.Zf{}
	zfstr := zf.Zf(true) + mh + dyh + zf.Zf(true) + dh + zf.Zf(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfstr)

	// kzf := zfzhi.Kongzifuzhi()
	kzfstr := zf.Kzf(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Kongzifuzhi(false) + xkhz + xkhy + hhf
	buffer.WriteString(kzfstr)

	jiahao := zfzhi.Jiahaozhi()
	// xhx := zfzhi.Xiahuaxianzhi()
	xhxstr := zf.Xhx(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Xiahuaxianzhi(false) + xkhz + xkhy + hhf
	buffer.WriteString(xhxstr)
	// if err != nil
	ifstr := zf.If(true) + kgf + zf.Err(true) + gth + dyh + zf.Nil(true)
	buffer.WriteString(ifstr)
	buffer.WriteString(dkhz + hhf)

	// return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()
	reterr := zf.Return(true) + kgf + zf.Baseinits(true) + dh + zf.Tishis(false) + zkhz + zf.Zf(true) + dh + zf.Tishi09(false) +
		xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Bianma(false) + jiahao + zf.Xhx(true) + jiahao +
		zf.Err(true) + dh + zf.Error(false) + xkhz + xkhy + hhf
	buffer.WriteString(reterr)
	buffer.WriteString(hhf + dkhy)

	findstr := hhf + bmx + zf.Find(true) + mh + dyh + zf.Chaxun(false) + bmx + xkhz + bmx + dh + zf.Id(false) + xkhy + hhf
	buffer.WriteString(findstr)
	iffind := zf.If(true) + kgf + bmx + zf.Find(true) + gth + dyh + zf.Nil(true)
	buffer.WriteString(iffind)
	buffer.WriteString(dkhz + hhf)

	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		if Huoquyigeleixing(lk) == zf.String(true) {
			buffer.WriteString(hhf)
			iflie := zf.If(true) + kgf + bmx + dh + lk + gth + dyh + zf.Kzf(true)
			buffer.WriteString(iflie)
			buffer.WriteString(dkhz + hhf)
			findlie := bmx + zf.Find(true) + dh + lk + dyh + bmx + dh + lk
			buffer.WriteString(findlie)
			buffer.WriteString(hhf + dkhy)
		}
	}
	xiugaistr := hhf + zf.Return(true) + kgf + zf.Zd(true) + bmx + zf.Daos(true) + dh + zf.Xiugaiyige(false) + xkhz + bmx + zf.Find(true) + xkhy + hhf

	buffer.WriteString(xiugaistr)
	buffer.WriteString(hhf + dkhy + hhf)
	finalstr := zf.Return(true) + kgf + zf.Baseinits(true) + dh + zf.Cuowus(false) + zkhz + zf.Zf(true) + dh + zf.Error04(false) + xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Zhi(false)
	buffer.WriteString(finalstr)

	buffer.WriteString(hhf + dkhy + hhf)

}
func servicetianjia(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	xh := zfzhi.Xinghaozhi()
	dh := zfzhi.Dianhaozhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	hhf := zfzhi.Huanhangfuzhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	gth := zfzhi.Gantanhaozhi()
	jiahao := zfzhi.Jiahaozhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()

	bmx := strings.ToLower(bianma)
	// func Tianjiaxxx(xxx *appmodels.Xxx)
	funcstr := zf.Func(true) + kgf + zf.Tianjia(false) + bmx + xkhz + bmx + kgf + xh + zf.Appmodels(true) + dh + bianma + xkhy + zf.String(true)
	buffer.WriteString(funcstr)
	buffer.WriteString(dkhz + hhf)
	// err := yanzhengyigechangdu(xxx)
	errstr := zf.Err(true) + mh + dyh + zf.Yanzhengziduanchangdu(true) + xkhz + bmx + xkhy + hhf
	buffer.WriteString(errstr)

	// zfzhi := zfzhi.Zfzhi{}
	zfzhistr := zf.Zfzhi(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Zfzhi(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfzhistr)
	//zf := zf.Zf{}
	zfstr := zf.Zf(true) + mh + dyh + zf.Zf(true) + dh + zf.Zf(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfstr)
	// xhx := zfzhi.Xiahuaxianzhi()
	xhxstr := zf.Xhx(true) + mh + dyh + zf.Zfzhi(true) + dh + zf.Xiahuaxianzhi(false) + xkhz + xkhy + hhf
	buffer.WriteString(xhxstr)

	// if err != nil{\n
	ifstr := zf.If(true) + kgf + zf.Err(true) + gth + dyh + zf.Nil(true) + dkhz + hhf
	buffer.WriteString(ifstr)
	// return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()
	reterr := zf.Return(true) + kgf + zf.Baseinits(true) + dh + zf.Tishis(false) + zkhz + zf.Zf(true) + dh + zf.Tishi09(false) + xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Bianma(false) + jiahao + zf.Xhx(true) + jiahao + zf.Err(true) + dh + zf.Error(false) + xkhz + xkhy + hhf
	buffer.WriteString(reterr)
	buffer.WriteString(hhf + dkhy + hhf)

	//return zdjuesedaos.Tianjiayige(juese)
	tjstr := zf.Return(true) + kgf + zf.Zd(true) + bmx + zf.Daos(true) + dh + zf.Tianjiayige(false) + xkhz + bmx + xkhy + hhf
	buffer.WriteString(tjstr)

	buffer.WriteString(hhf + dkhy + hhf)
}
func servicechaxun(bianma string, buffer *bytes.Buffer) {
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	dh := zfzhi.Dianhaozhi()
	xh := zfzhi.Xinghaozhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()

	bmx := strings.ToLower(bianma)

	funstr := zf.Func(true) + kgf + zf.Chaxun(false) + bmx + xkhz + zf.Id(true) + kgf + zf.Int(true) + xkhy + xh + zf.Appmodels(true) + dh + bianma
	buffer.WriteString(funstr)
	buffer.WriteString(dkhz + hhf)
	retstr := zf.Return(true) + kgf + zf.Zd(true) + bmx + zf.Daos(true) + dh + zf.Chaxunyige(false) + xkhz + zf.Id(true) + xkhy
	buffer.WriteString(retstr)
	buffer.WriteString(hhf + dkhy)
	buffer.WriteString(hhf)
}
func serviceshanchu(bianma string, buffer *bytes.Buffer) {
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	dh := zfzhi.Dianhaozhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()

	bmx := strings.ToLower(bianma)

	// func Shanchuxxx(id int) string
	funstr := zf.Func(true) + kgf + zf.Shanchu(false) + bmx + xkhz + zf.Id(true) + kgf + zf.Int(true) + xkhy + zf.String(true)
	buffer.WriteString(funstr)
	buffer.WriteString(dkhz + hhf)
	retstr := zf.Return(true) + kgf + zf.Zd(true) + bmx + zf.Daos(true) + dh + zf.Shanchuyige(false) + xkhz + zf.Id(true) + xkhy
	buffer.WriteString(retstr)
	buffer.WriteString(hhf + dkhy)
	buffer.WriteString(hhf)
}

func Shengchengservice() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()
	xx := zfzhi.Xiexianzhi()
	dh := zfzhi.Dianhaozhi()
	for bk, bv := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		bm := zf.Zd(true) + bv + zf.Services(true)
		//package xxxservices \n
		pac := zf.Package(true) + kgf + bm + hhf
		buffer.WriteString(pac)

		serviceimports(bk, &buffer)
		yanzhengchangdu(bk, &buffer)
		servicetianjia(bk, &buffer)
		servicexiugai(bk, &buffer)
		serviceshanchu(bk, &buffer)
		servicechaxun(bk, &buffer)
		dir := basemodels.Getapppath() + xx + bm
		os.MkdirAll(dir, os.ModePerm)
		path := dir + xx + bk + zf.Services(true) + dh + zf.Go(true)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
