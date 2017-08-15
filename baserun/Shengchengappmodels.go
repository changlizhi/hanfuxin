package baserun

import (
	"bytes"
	"hanfuxin/appinits"
	"changliang/zf"
	"changliang/zfzhi"
	"io/ioutil"
	"os"
)

func Shengchenghanfuxinmodels() {
	bm := zf.Zfs.Appmodels(true) //包名 appmodels

	for bk, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		// model生成时大多数都是一样的，所以提供一个公用的package,import之类的
		Buffertoumodel(&buffer, bk, bm)

		for lk, _ := range Huoquyigebiaojiegou(bk) {
			field := lk + zfzhi.Zhi.Kgf() + Huoquyigeleixing(lk) + zfzhi.Zhi.Hhf()
			buffer.WriteString(field)
		}
		//左大括号在头里有了
		buffer.WriteString(zfzhi.Zhi.Dkhy()) // }
		// hanfuxin/appmodels/Juese.go
		path := appinits.Getapppath() + zfzhi.Zhi.Xx() + bm + zfzhi.Zhi.Xx() + bk + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
