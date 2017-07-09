package baserun

// 在应用启动之前生成一些必要的model
import (
	"bytes"
	"hanfuxin/baseinits"
	"hanfuxin/basemodels"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"io/ioutil"
	"log"
	"os"
)

func Shengchengyingyongzimodel() {
	buffer := bytes.Buffer{}
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
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
	log.Println("pathdaos---------", path)
	log.Println("bufferdaos---------", buffer.String())
	ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
}
