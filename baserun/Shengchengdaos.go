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
	// \n import (\n
	importstr := zfzhi.Zhi.Hhf() + zf.Zfs.Import(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(importstr)

	// "hanfuxin/appinits" \n
	api := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Appinits(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(api)

	// "hanfuxin/appmodels" \n
	apm := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(apm)

	// "hanfuxin/baseinits"
	bai := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bai)

	//"hanfuxin/zf"
	zfbao := zfzhi.Zhi.Syh() + zf.Zfs.Hanfuxin(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zf(true) + zfzhi.Zhi.Syh()
	buffer.WriteString(zfbao)

	// ) \n
	buffer.WriteString(zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
}
func chaxunyige(bianma string, buffer *bytes.Buffer) {
	// func Chaxunyige (id int)
	idx := zf.Zfs.Id(true)
	idd := zf.Zfs.Id(false)

	bmx := strings.ToLower(bianma)

	//func Chaxunyige(id int)
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chaxunyige(false) + zfzhi.Zhi.Xkhz() + idx + zfzhi.Zhi.Kgf() + zf.Zfs.Int(true) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)

	// *appmodels.Juese{ \n
	fanhui := zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fanhui)

	// juese := &appmodels.Juese
	modelstr := bmx + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Qh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(modelstr)

	// {Id:id}
	idstr := zfzhi.Zhi.Dkhz() + idd + zfzhi.Zhi.Mh() + idx + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(idstr)

	// \n err := appinits.Hanfuxinormer.Read(juese)
	read := zf.Zfs.Err(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Appinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Hanfuxinormer(false) + zfzhi.Zhi.Dh() + zf.Zfs.Read(false) + zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(read)

	// if err := nil
	errstrss := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(errstrss)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//return nil
	retstrsss := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Nil(true)
	buffer.WriteString(retstrsss)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	// \n return juese
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + bmx + zfzhi.Zhi.Hhf()
	buffer.WriteString(retstr)

	//} \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func tianjiayige(bianma string, buffer *bytes.Buffer) {
	// func Tianjiayige

	bmx := strings.ToLower(bianma)

	funcstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tianjiayige(false)
	buffer.WriteString(funcstr)

	//(juese*appmodels.Juese){\n
	canshu := zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(canshu)

	// zf := zf.Zfs.Zf{} \n
	zfstr := zf.Zfs.Zf(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zf(false) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfstr)

	// _, err := appinits.Hanfuxinormer.Insert(juese)
	insert := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Appinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Hanfuxinormer(false) + zfzhi.Zhi.Dh() + zf.Zfs.Insert(false) + zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(insert)

	// if err != nil
	iferr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(iferr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	// return baseinits.Tishis[zf.Zfs.Tishi04].Bianma
	errret := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi04(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(errret)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	// return baseinits.Tishis[zf.Zfs.Tishi03].Bianma
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi03(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(retstr)
	//\n}\n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func tianjiaduoge(bianma string, buffer *bytes.Buffer) {
	//func Tianjiaduoge(
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tianjiaduoge(false) + zfzhi.Zhi.Xkhz()
	buffer.WriteString(funstr)

	//jueseshuzu
	canshu := strings.ToLower(bianma) + zf.Zfs.Shuzu(true)
	buffer.WriteString(canshu)

	//[]appmodels.Juese{\n
	cslx := zfzhi.Zhi.Kgf() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cslx)
	// zf := zf.Zfs.Zf{} \n
	zfstr := zf.Zfs.Zf(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zf(false) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfstr)

	//_err:=appinits.Hanfuxinormer.InsertMulti
	insert := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Appinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Hanfuxinormer(false) + zfzhi.Zhi.Dh() + zf.Zfs.InsertMulti(false)
	buffer.WriteString(insert)

	//(len(jueseshuzu+,jueseshuzu+
	lenstr := zfzhi.Zhi.Xkhz() + zf.Zfs.Len(true) + zfzhi.Zhi.Xkhz() + canshu + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + canshu + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lenstr)
	// if err != nil
	iferr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(iferr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	// return baseinits.Tishis[zf.Zfs.Tishi04].Bianma
	errret := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi04(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(errret)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	// return baseinits.Tishis[zf.Zfs.Tishi03].Bianma
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi03(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(retstr)

	//}\n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func shanchuyige(bianma string, buffer *bytes.Buffer) {
	//func Shanchuyige(id int+{ \n
	funcstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Shanchuyige(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.Id(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Int(true) + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(funcstr)
	// zf := zf.Zfs.Zf{} \n
	zfstr := zf.Zfs.Zf(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zf(false) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfstr)

	//_,errappinits.Hanfuxinormer.Delete
	delete := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Appinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Hanfuxinormer(false) + zfzhi.Zhi.Dh() + zf.Zfs.Delete(false)
	buffer.WriteString(delete)

	//(Chaxunyige(id++
	cxstr := zfzhi.Zhi.Xkhz() + zf.Zfs.Chaxunyige(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.Id(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cxstr)
	// if err != nil
	iferr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(iferr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	// return baseinits.Tishis[zf.Zfs.Tishi04].Bianma
	errret := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi08(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(errret)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	// return baseinits.Tishis[zf.Zfs.Tishi03].Bianma
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi07(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(retstr)

	//\n}\n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func xiugaiyige(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Xiugaiyige(false)
	buffer.WriteString(funstr)

	//(juese*appmodels.Juese){\n
	csstr := zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Appmodels(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)
	// zf := zf.Zfs.Zf{} \n
	zfstr := zf.Zfs.Zf(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zf(false) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfstr)

	//_,err := appinits.Hanfuxinormer.Update(juese)
	update := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Appinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Hanfuxinormer(false) + zfzhi.Zhi.Dh() + zf.Zfs.Update(false) + zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(update)
	// if err != nil
	iferr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)

	buffer.WriteString(iferr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	// return baseinits.Tishis[zf.Zfs.Tishi04].Bianma
	errret := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi06(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(errret)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	// return baseinits.Tishis[zf.Zfs.Tishi03].Bianma
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Baseinits(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi05(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(retstr)

	//\n}\n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func Shengchengdaos() {
	for biao, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		buffer.WriteString(zf.Zfs.Package(true))    //package
		buffer.WriteString(zfzhi.Zhi.Konggefuzhi()) // zfzhi.Zhi.Kgf()
		lujing := zf.Zfs.Zd(true) + strings.ToLower(biao) + zf.Zfs.Daos(true)
		buffer.WriteString(lujing) // zdjuesedaos

		imports(biao, &buffer)      //import all
		chaxunyige(biao, &buffer)   // Chaxunyige
		tianjiayige(biao, &buffer)  // Tianjiayige
		tianjiaduoge(biao, &buffer) // Tianjiaduoge
		shanchuyige(biao, &buffer)  // Shanchuyige
		xiugaiyige(biao, &buffer)   // Xiugaiyige

		dir := basemodels.Getapppath() + zfzhi.Zhi.Xx() + lujing + zfzhi.Zhi.Xx()
		path := dir + biao + zf.Zfs.Daos(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
