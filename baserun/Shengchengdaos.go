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

func imports(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	xx := zfzhi.Xiexianzhi()
	hfx := zf.Hanfuxin(true)
	syh := zfzhi.Shuangyinhaozhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	kgf := zfzhi.Konggefuzhi()
	// \n import (\n
	importstr := hhf + zf.Import(true) + kgf + xkhz + hhf
	buffer.WriteString(importstr)

	// "hanfuxin/appinits" \n
	api := syh + hfx + xx + zf.Appinits(true) + syh + hhf
	buffer.WriteString(api)

	// "hanfuxin/appmodels" \n
	apm := syh + hfx + xx + zf.Appmodels(true) + syh + hhf
	buffer.WriteString(apm)

	// "hanfuxin/baseinits"
	bai := syh + hfx + xx + zf.Baseinits(true) + syh + hhf
	buffer.WriteString(bai)

	//"hanfuxin/zf"
	zfbao := syh + hfx + xx + zf.Zf(true) + syh
	buffer.WriteString(zfbao)

	// ) \n
	buffer.WriteString(xkhy + hhf)
}
func chaxunyige(bianma string, buffer *bytes.Buffer) {
	// func Chaxunyige (id int)
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	idx := zf.Id(true)
	idd := zf.Id(false)
	appmodels := zf.Appmodels(true)

	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kongge := zfzhi.Konggefuzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	qh := zfzhi.Qiehaozhi()
	xh := zfzhi.Xinghaozhi()
	gth := zfzhi.Gantanhaozhi()
	kgf := zfzhi.Konggefuzhi()

	bmx := strings.ToLower(bianma)

	//func Chaxunyige(id int)
	funstr := zf.Func(true) + kongge + zf.Chaxunyige(false) + xkhz + idx + kongge + zf.Int(true) + xkhy
	buffer.WriteString(funstr)

	// *appmodels.Juese{ \n
	fanhui := kongge + xh + appmodels + dh + bianma + kongge + dkhz + hhf
	buffer.WriteString(fanhui)

	// juese := &appmodels.Juese
	modelstr := bmx + kongge + mh + dyh + kongge + qh + appmodels + dh + bianma
	buffer.WriteString(modelstr)

	// {Id:id}
	idstr := dkhz + idd + mh + idx + dkhy + hhf
	buffer.WriteString(idstr)

	// \n err := appinits.Hanfuxinormer.Read(juese)
	read := zf.Err(true) + mh + dyh + zf.Appinits(true) + dh + zf.Hanfuxinormer(false) + dh + zf.Read(false) + xkhz + bmx + xkhy + hhf
	buffer.WriteString(read)

	// if err := nil
	errstrss := zf.If(true) + kgf + zf.Err(true) + kgf + gth + dyh + zf.Nil(true)
	buffer.WriteString(errstrss)

	buffer.WriteString(dkhz + hhf)

	//return nil
	retstrsss := zf.Return(true) + kgf + zf.Nil(true)
	buffer.WriteString(retstrsss)

	buffer.WriteString(hhf + dkhy + hhf)

	// \n return juese
	retstr := zf.Return(true) + kongge + bmx + hhf
	buffer.WriteString(retstr)

	//} \n
	buffer.WriteString(dkhy + hhf)
}
func tianjiayige(bianma string, buffer *bytes.Buffer) {
	// func Tianjiayige
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	douhao := zfzhi.Douhaozhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	gth := zfzhi.Gantanhaozhi()

	bmx := strings.ToLower(bianma)

	funcstr := zf.Func(true) + kgf + zf.Tianjiayige(false)
	buffer.WriteString(funcstr)

	//(juese*appmodels.Juese){\n
	canshu := xkhz + bmx + kgf + zfzhi.Xinghaozhi() + zf.Appmodels(true) + dh + bianma + xkhy + zf.String(true) + dkhz + hhf
	buffer.WriteString(canshu)

	// zf := zf.Zf{} \n
	zfstr := zf.Zf(true) + mh + dyh + zf.Zf(true) + dh + zf.Zf(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfstr)

	// _, err := appinits.Hanfuxinormer.Insert(juese)
	insert := xhx + douhao + zf.Err(true) + mh + dyh + zf.Appinits(true) + dh + zf.Hanfuxinormer(false) + dh + zf.Insert(false) + xkhz + bmx + xkhy + hhf
	buffer.WriteString(insert)

	// if err != nil
	iferr := zf.If(true) + kgf + zf.Err(true) + gth + dyh + zf.Nil(true)
	buffer.WriteString(iferr)

	buffer.WriteString(dkhz + hhf)

	// return baseinits.Tishis[zf.Tishi04].Bianma
	errret := zf.Return(true) + kgf + zf.Baseinits(true) + dh + zf.Tishis(false) + zkhz + zf.Zf(true) + dh + zf.Tishi04(false) + xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Bianma(false) + hhf
	buffer.WriteString(errret)
	buffer.WriteString(dkhy + hhf)
	// return baseinits.Tishis[zf.Tishi03].Bianma
	retstr := zf.Return(true) + kgf + zf.Baseinits(true) + dh + zf.Tishis(false) + zkhz + zf.Zf(true) + dh + zf.Tishi03(false) + xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Bianma(false) + hhf
	buffer.WriteString(retstr)
	//\n}\n
	buffer.WriteString(dkhy + hhf)
}
func tianjiaduoge(bianma string, buffer *bytes.Buffer) {
	//func Tianjiaduoge(
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	douhao := zfzhi.Douhaozhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	kgf := zfzhi.Konggefuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	gth := zfzhi.Gantanhaozhi()

	//func Tianjiaduoge(
	funstr := zf.Func(true) + kgf + zf.Tianjiaduoge(false) + xkhz
	buffer.WriteString(funstr)

	//jueseshuzu
	canshu := strings.ToLower(bianma) + zf.Shuzu(true)
	buffer.WriteString(canshu)

	//[]appmodels.Juese{\n
	cslx := kgf + zkhz + zkhy + zf.Appmodels(true) + dh + bianma + xkhy + zf.String(true) + dkhz + hhf
	buffer.WriteString(cslx)
	// zf := zf.Zf{} \n
	zfstr := zf.Zf(true) + mh + dyh + zf.Zf(true) + dh + zf.Zf(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfstr)

	//_err:=appinits.Hanfuxinormer.InsertMulti
	insert := xhx + douhao + zf.Err(true) + mh + dyh + zf.Appinits(true) + dh + zf.Hanfuxinormer(false) + dh + zf.InsertMulti(false)
	buffer.WriteString(insert)

	//(len(jueseshuzu+,jueseshuzu+
	lenstr := xkhz + zf.Len(true) + xkhz + canshu + xkhy + douhao + canshu + xkhy + hhf
	buffer.WriteString(lenstr)
	// if err != nil
	iferr := zf.If(true) + kgf + zf.Err(true) + gth + dyh + zf.Nil(true)
	buffer.WriteString(iferr)

	buffer.WriteString(dkhz + hhf)

	// return baseinits.Tishis[zf.Tishi04].Bianma
	errret := zf.Return(true) + kgf + zf.Baseinits(true) + dh + zf.Tishis(false) + zkhz + zf.Zf(true) + dh + zf.Tishi04(false) + xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Bianma(false) + hhf
	buffer.WriteString(errret)
	buffer.WriteString(dkhy + hhf)
	// return baseinits.Tishis[zf.Tishi03].Bianma
	retstr := zf.Return(true) + kgf + zf.Baseinits(true) + dh + zf.Tishis(false) + zkhz + zf.Zf(true) + dh + zf.Tishi03(false) + xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Bianma(false) + hhf
	buffer.WriteString(retstr)

	//}\n
	buffer.WriteString(dkhy + hhf)
}
func shanchuyige(bianma string, buffer *bytes.Buffer) {
	// func Shanchuyige
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	douhao := zfzhi.Douhaozhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	id := zf.Id(true)
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	gth := zfzhi.Gantanhaozhi()

	//func Shanchuyige(id int+{ \n
	funcstr := zf.Func(true) + kgf + zf.Shanchuyige(false) + xkhz + id + kgf + zf.Int(true) + xkhy + zf.String(true) + dkhz + hhf
	buffer.WriteString(funcstr)
	// zf := zf.Zf{} \n
	zfstr := zf.Zf(true) + mh + dyh + zf.Zf(true) + dh + zf.Zf(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfstr)

	//_,errappinits.Hanfuxinormer.Delete
	delete := xhx + douhao + zf.Err(true) + mh + dyh + zf.Appinits(true) + dh + zf.Hanfuxinormer(false) + dh + zf.Delete(false)
	buffer.WriteString(delete)

	//(Chaxunyige(id++
	cxstr := xkhz + zf.Chaxunyige(false) + xkhz + id + xkhy + xkhy + hhf
	buffer.WriteString(cxstr)
	// if err != nil
	iferr := zf.If(true) + kgf + zf.Err(true) + gth + dyh + zf.Nil(true)
	buffer.WriteString(iferr)

	buffer.WriteString(dkhz + hhf)

	// return baseinits.Tishis[zf.Tishi04].Bianma
	errret := zf.Return(true) + kgf + zf.Baseinits(true) + dh + zf.Tishis(false) + zkhz + zf.Zf(true) + dh + zf.Tishi08(false) + xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Bianma(false) + hhf
	buffer.WriteString(errret)
	buffer.WriteString(dkhy + hhf)
	// return baseinits.Tishis[zf.Tishi03].Bianma
	retstr := zf.Return(true) + kgf + zf.Baseinits(true) + dh + zf.Tishis(false) + zkhz + zf.Zf(true) + dh + zf.Tishi07(false) + xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Bianma(false) + hhf
	buffer.WriteString(retstr)

	//\n}\n
	buffer.WriteString(dkhy + hhf)
}
func xiugaiyige(bianma string, buffer *bytes.Buffer) {
	// func Xiugaiyige
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xh := zfzhi.Xinghaozhi()
	xhx := zfzhi.Xiahuaxianzhi()
	douhao := zfzhi.Douhaozhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	gth := zfzhi.Gantanhaozhi()

	bmx := strings.ToLower(bianma)

	funstr := zf.Func(true) + kgf + zf.Xiugaiyige(false)
	buffer.WriteString(funstr)

	//(juese*appmodels.Juese){\n
	csstr := xkhz + bmx + kgf + xh + zf.Appmodels(true) + dh + bianma + xkhy + zf.String(true) + dkhz + hhf
	buffer.WriteString(csstr)
	// zf := zf.Zf{} \n
	zfstr := zf.Zf(true) + mh + dyh + zf.Zf(true) + dh + zf.Zf(false) + dkhz + dkhy + hhf
	buffer.WriteString(zfstr)

	//_,err := appinits.Hanfuxinormer.Update(juese)
	update := xhx + douhao + zf.Err(true) + mh + dyh + zf.Appinits(true) + dh + zf.Hanfuxinormer(false) + dh + zf.Update(false) + xkhz + bmx + xkhy + hhf
	buffer.WriteString(update)
	// if err != nil
	iferr := zf.If(true) + kgf + zf.Err(true) + gth + dyh + zf.Nil(true)

	buffer.WriteString(iferr)

	buffer.WriteString(dkhz + hhf)

	// return baseinits.Tishis[zf.Tishi04].Bianma
	errret := zf.Return(true) + kgf + zf.Baseinits(true) + dh + zf.Tishis(false) + zkhz + zf.Zf(true) + dh + zf.Tishi06(false) + xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Bianma(false) + hhf
	buffer.WriteString(errret)
	buffer.WriteString(dkhy + hhf)
	// return baseinits.Tishis[zf.Tishi03].Bianma
	retstr := zf.Return(true) + kgf + zf.Baseinits(true) + dh + zf.Tishis(false) + zkhz + zf.Zf(true) + dh + zf.Tishi05(false) + xkhz + zf.False(true) + xkhy + zkhy + dh + zf.Bianma(false) + hhf
	buffer.WriteString(retstr)

	//\n}\n
	buffer.WriteString(dkhy + hhf)

}
func Shengchengdaos() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	dh := zfzhi.Dianhaozhi()
	xx := zfzhi.Xiexianzhi()

	daos := zf.Daos(true)
	for biao, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		buffer.WriteString(zf.Package(true))    //package
		buffer.WriteString(zfzhi.Konggefuzhi()) // kongge
		lujing := zf.Zd(true) + strings.ToLower(biao) + daos
		buffer.WriteString(lujing) // zdjuesedaos

		imports(biao, &buffer)      //import all
		chaxunyige(biao, &buffer)   // Chaxunyige
		tianjiayige(biao, &buffer)  // Tianjiayige
		tianjiaduoge(biao, &buffer) // Tianjiaduoge
		shanchuyige(biao, &buffer)  // Shanchuyige
		xiugaiyige(biao, &buffer)   // Xiugaiyige

		dir := basemodels.Getapppath() + xx + lujing + xx
		path := dir + biao + daos + dh + zf.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
