package baserun

// 在应用启动之前生成一些必要的model
import (
	"bytes"
	"go/token"
	"hanfuxin/baseinits"
	"hanfuxin/basemodels"
	"hanfuxin/zfz"
	"hanfuxin/zfzhi"
	"io/ioutil"
	"os"
)

func Buffertoumodel(buffer *bytes.Buffer, bianma string, baoming string) {
	zf := zfz.Zf{}
	// package xxx \n
	kgf := zfzhi.Konggefuzhi()
	hhf = zfzhi.Huanhangfuzhi()
	buffer.WriteString(zf.Package(true))
	buffer.WriteString(kgf)
	buffer.WriteString(baoming)
	buffer.WriteString(hhf)

	// type Juese struct{\n
	buffer.WriteString(zf.Type(true))
	buffer.WriteString(kgf)
	buffer.WriteString(bianma)
	buffer.WriteString(kgf)
	buffer.WriteString(zf.Struct(true))
	buffer.WriteString(zfzhi.Dakuohaozuozhi())
	buffer.WriteString(hhf)
}

func Shengchengyingyongzimodel() {
	buffer := bytes.Buffer{}
	zf := zfz.Zf{}
	wjm := zf.Suoyouzf(false) //文件名 Suoyouzf，不包含格式
	bm := zf.Zfmodel(true)    //包名 zfmodel
	xx := zfzhi.Xiexianzhi()  //斜线 /
	dh := zfzhi.Dianhaozhi()  //点号 .

	Buffertoumodel(&buffer, wjm, bm)
	for k := range baseinits.Zifus {
		buffer.WriteString(k)
		buffer.WriteString(zfzhi.Konggefuzhi())
		buffer.WriteString(zf.String(true))
		buffer.WriteString(zfzhi.Huanhangfuzhi())
	}
	buffer.WriteString(zfzhi.Dakuohaoyouzhi())
	path := basemodels.Getapppath() + xx + bm + xx + wjm + dh + zf.Go(true)
	ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
}
