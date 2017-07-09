package baserun

import (
	"bytes"
	"hanfuxin/basemodels"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"io/ioutil"
	"os"
)

func Shengchenghanfuxinmodels() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	xx := zfzhi.Xiexianzhi() // 斜线 /
	bm := zf.Appmodels(true) //包名 appmodels
	dh := zfzhi.Dianhaozhi() //点号 .
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()
	dkhy := zfzhi.Dakuohaoyouzhi()

	for bk, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		// model生成时大多数都是一样的，所以提供一个公用的package,import之类的
		Buffertoumodel(&buffer, bk, bm)

		for lk, _ := range Huoqulies() {
			lie := Huoqubiaolieguanlian(bk, lk)
			for lk, _ := range lie {
				buffer.WriteString(lk)                   //字段字符
				buffer.WriteString(kgf)                  //空格
				buffer.WriteString(Huoquyigeleixing(lk)) // 类型(int,string,float32,time)
				buffer.WriteString(hhf)                  //换行
			}
		}
		//左大括号在头里有了
		buffer.WriteString(dkhy) // }
		// hanfuxin/appmodels/Juese.go
		path := basemodels.Getapppath() + xx + bm + xx + bk + dh + zf.Go(true)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
