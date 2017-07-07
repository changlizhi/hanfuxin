package baserun

import (
	"bytes"
	"go/token"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/basemodels"
	"hanfuxin/zfz"
	"hanfuxin/zfzhi"
	"io/ioutil"
	"os"
)

func Shengchenghanfuxinmodels() {
	biaos := baseinits.Biaos
	lies := baseinits.Lies
	zf := zfz.Zf{}
	xx := zfzhi.Xiexianzhi() // 斜线 /
	bm := zf.Appmodels(true) //包名 appmodels
	dh := zfzhi.Dianhaozhi() //点号 .
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	for bk, bv := range biaos {
		buffer := bytes.Buffer{}
		// model生成时大多数都是一样的，所以提供一个公用的package,import之类的
		Buffertoumodel(&buffer, bv.Bianma, bm)
		for _, lv := range lies {
			pipei, _ := apputils.Pipei3lei(bk, lv.Biaoming) //判断表名是否在lie的Biaoming里
			if pipei {
				// 如: Bianma string \n
				buffer.WriteString(lv.Bianma)  //字段字符
				buffer.WriteString(kgf)        //空格
				buffer.WriteString(lv.Leixing) // 类型(int,string,float32,time)
				buffer.WriteString(hhf)        //换行
			}
		}
		//左大括号在头里有了
		buffer.WriteString(dkhy) // }
		// hanfuxin/appmodels/Juese.go
		path := basemodels.Getapppath() + xx + bm + xx + bv.Bianma + dh + zf.Go(true)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
