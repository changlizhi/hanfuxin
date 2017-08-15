package baserun

import (
	"bytes"
	"hanfuxin/appinits"
	"changliang/zf"
	"changliang/zfzhi"
	"io/ioutil"
	"os"
	"strings"
)

func serviceimports(bianma string, buffer *bytes.Buffer) {
	importstr := zf.Zfs.Import(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(importstr)
	//"hanfuxin/allerrors" \n
	errorbao := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Allerrors(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(errorbao)
	// "bytes"
	bytesbao := zfzhi.Zhi.Syh() + zf.Zfs.Bytes(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bytesbao)
	// "time"
	timebao := zfzhi.Zhi.Syh() + zf.Zfs.Time(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(timebao)
	// "hanfuxin/appmodels"
	apmbao := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(apmbao)
	// "hanfuxin/apputils"
	apubao := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Apputils(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(apubao)
	// "hanfuxin/baserun"
	brbao := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Baserun(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(brbao)
	// "changliang/zf"
	zfbao := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfbao)
	// "changliang/zfzhi"
	zfzhibao := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfzhibao)
	// "hanfuxin/Xxxdaos"
	daobao := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Daos(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(daobao)
	//"hanfuxin/appinits"
	binitbao := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Syh()
	buffer.WriteString(binitbao)
	buffer.WriteString(zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
}
func yanzhengchangdu(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	funcone := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Yanzhengziduanchangdu(true)
	buffer.WriteString(funcone)
	//(bmx *appmodels.bm) error { \n
	canshu := zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Xkhy() + zf.Zfs.Error(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()

	buffer.WriteString(canshu)
	cuowu := zf.Zfs.Cuowu(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.False(true) + zfzhi.Zhi.Hhf()
	buffer.WriteString(cuowu)

	cuowutrue := zf.Zfs.Cuowu(true) + zfzhi.Zhi.Dyh() + zf.Zfs.True(true) + zfzhi.Zhi.Hhf()
	//buffer := bytes.Buffer{}\n
	bf := zf.Zfs.Buffer(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Bytes(true) + zfzhi.Zhi.Dh() + zf.Zfs.Buffer(false) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bf)
	for lk, lv := range Huoquyigebiaojiegou(bianma) {
		if Huoquyigeleixing(lk) == zf.Zfs.String(true) {
			buffer.WriteString(zfzhi.Zhi.Hhf())
			lenlvstr := zf.Zfs.Len(true) + lv
			lenlvshitistr := lenlvstr + zf.Zfs.Shiti(true)
			//lenbianma :=
			//baserun.Huoquyigechangdu
			//(zf.Zfs.Bianma(false))

			lenlv := lenlvstr + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
				zf.Zfs.Baserun(true) + zfzhi.Zhi.Dh() + zf.Zfs.Huoquyigechangdu(false) +
				zfzhi.Zhi.Xkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() +
				zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + lk + zfzhi.Zhi.Xkhz() +
				zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
			buffer.WriteString(lenlv)

			//	lenbianmashiti := len(juese.Bianma)
			lenshiti := lenlvshitistr + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Len(true) + zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Dh() + lk + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
			buffer.WriteString(lenshiti)

			// if leibianmashiti > int(lenbianma) {\n
			oneif := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + lenlvshitistr +
				zfzhi.Zhi.Dy() + zf.Zfs.Int(true) + zfzhi.Zhi.Xkhz() +
				lenlvstr + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
			buffer.WriteString(oneif)
			buffer.WriteString(cuowutrue)
			//buffer.WriteString(....)
			writestr := zf.Zfs.Buffer(true) + zfzhi.Zhi.Dh() + zf.Zfs.WriteString(false)
			buffer.WriteString(writestr)

			// (apputils.Shengchengerrorchangdu
			scstr := zfzhi.Zhi.Xkhz() + zf.Zfs.Apputils(true) + zfzhi.Zhi.Dh() + zf.Zfs.Shengchengerrorchangdu(false)
			buffer.WriteString(scstr)

			// (zf.Zfs.Mingcheng(false)
			csstrs := zfzhi.Zhi.Xkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Mingcheng(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy()
			buffer.WriteString(csstrs)

			//, int64(xxx)
			int64str := zfzhi.Zhi.Dou() + zf.Zfs.Int64(true) + zfzhi.Zhi.Xkhz() + lenlvstr + zfzhi.Zhi.Xkhy()
			buffer.WriteString(int64str)

			shitis := zfzhi.Zhi.Dou() + lenlvshitistr + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
			buffer.WriteString(shitis)
			buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
		}
	}
	// if cuowu {\n
	cwstr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Cuowu(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cwstr)
	// return allerrors.Ziduanerror
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Allerrors(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ziduanerror(false)
	buffer.WriteString(retstr)
	//{Shijian:time.Now(),Wenti:buffer.String()}
	wtsjstr := zfzhi.Zhi.Dkhz() + zf.Zfs.Shijian(false) + zfzhi.Zhi.Mh() + zf.Zfs.Time(true) + zfzhi.Zhi.Dh() + zf.Zfs.Now(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(wtsjstr)
	wtwtstr := zfzhi.Zhi.Dou() + zf.Zfs.Wenti(false) + zfzhi.Zhi.Mh() + zf.Zfs.Buffer(true) + zfzhi.Zhi.Dh() + zf.Zfs.String(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(wtwtstr)
	buffer.WriteString(zfzhi.Zhi.Dkhy())
	buffer.WriteString(zfzhi.Zhi.Hhf())

	nilstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Nil(true) + zfzhi.Zhi.Hhf()
	buffer.WriteString(nilstr)

	buffer.WriteString(zfzhi.Zhi.Dkhy())
	buffer.WriteString(zfzhi.Zhi.Hhf())
}
func servicexiugai(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	funcstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Xiugai(false) + bmx + zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true)
	buffer.WriteString(funcstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	errstr := zf.Zfs.Err(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Yanzhengziduanchangdu(true) + zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(errstr)

	// if err != nil
	ifstr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(ifstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	// return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xh()x + err.Error()
	reterr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Baseinits(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi09(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) +
		zfzhi.Zhi.Jia() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Xhx(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Jia() +
		zf.Zfs.Err(true) + zfzhi.Zhi.Dh() + zf.Zfs.Error(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(reterr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy())

	// xxxfind := Chaxunxxx(xxx.Id)
	findstr := zfzhi.Zhi.Hhf() + bmx + zf.Zfs.Find(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Dyh() + zf.Zfs.Chaxun(false) + bmx +
		zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Dh() + zf.Zfs.Id(false) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(findstr)

	// if xxxfind != nil
	iffind := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + bmx + zf.Zfs.Find(true) + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(iffind)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	for lk, _ := range Huoquyigebiaojiegou(bianma) {
		if Huoquyigeleixing(lk) == zf.Zfs.String(true) {
			buffer.WriteString(zfzhi.Zhi.Hhf())
			iflie := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + bmx + zfzhi.Zhi.Dh() + lk + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() +

				zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Kzf(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()

			buffer.WriteString(iflie)
			buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
			findlie := bmx + zf.Zfs.Find(true) + zfzhi.Zhi.Dh() + lk + zfzhi.Zhi.Dyh() + bmx + zfzhi.Zhi.Dh() + lk
			buffer.WriteString(findlie)
			buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy())
		}
	}
	xiugaistr := zfzhi.Zhi.Hhf() + zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.Zd(true) + bmx + zf.Zfs.Daos(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Xiugaiyige(false) + zfzhi.Zhi.Xkhz() + bmx +
		zf.Zfs.Find(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()

	buffer.WriteString(xiugaistr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	finalstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Baseinits(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Cuowus(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Error04(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false)
	buffer.WriteString(finalstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func servicetianjia(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	// func Tianjiazfzhi.Zhi.Xx()x(zfzhi.Zhi.Xx()x *appmodels.Xzfzhi.Zhi.Xx())
	funcstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tianjia(false) + bmx + zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true)
	buffer.WriteString(funcstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	// err := yanzhengyigechangdu(zfzhi.Zhi.Xx()x)
	errstr := zf.Zfs.Err(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Yanzhengziduanchangdu(true) + zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(errstr)

	// if err != nil{\n
	ifstr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Nil(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ifstr)

	// return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xh()x + err.Error()
	reterr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Baseinits(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi09(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) +
		zfzhi.Zhi.Jia() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Xhx(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +

		zfzhi.Zhi.Jia() +
		zf.Zfs.Err(true) + zfzhi.Zhi.Dh() + zf.Zfs.Error(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(reterr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//return zdjuesedaos.Tianjiayige(juese)
	tjstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zd(true) + bmx + zf.Zfs.Daos(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tianjiayige(false) + zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tjstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func servicechaxun(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chaxun(false) + bmx +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Id(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.Int(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xh() +
		zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zd(true) + bmx +
		zf.Zfs.Daos(true) + zfzhi.Zhi.Dh() + zf.Zfs.Chaxunyige(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Id(true) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(retstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy())
	buffer.WriteString(zfzhi.Zhi.Hhf())
}
func serviceshanchu(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	// func ShanchuXxx(id int) string
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Shanchu(false) + bmx +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Id(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.Int(true) + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true)
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zd(true) + bmx +
		zf.Zfs.Daos(true) + zfzhi.Zhi.Dh() + zf.Zfs.Shanchuyige(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Id(true) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(retstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy())
	buffer.WriteString(zfzhi.Zhi.Hhf())
}

func Shengchengservice() {
	for bk, bv := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		bm := zf.Zfs.Zd(true) + bv + zf.Zfs.Services(true)
		//package zfzhi.Zhi.Xx()xservices \n
		pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + bm + zfzhi.Zhi.Hhf()
		buffer.WriteString(pac)

		serviceimports(bk, &buffer)
		yanzhengchangdu(bk, &buffer)
		servicetianjia(bk, &buffer)
		servicexiugai(bk, &buffer)
		serviceshanchu(bk, &buffer)
		servicechaxun(bk, &buffer)
		dir := appinits.Getapppath() + zfzhi.Zhi.Xx() + bm
		os.MkdirAll(dir, os.ModePerm)
		path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Services(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
