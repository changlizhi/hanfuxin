package baserun

// 在应用启动之前生成一些必要的model
import (
	"bytes"
	"hanfuxin/baseinits"
	"io/ioutil"
	"os"
)
func Buffertoumodel(buffer *bytes.Buffer, bianma string) {
	gen := baseinits.Gen
	fuhao := baseinits.Fuhaos
	mulu := baseinits.Mulus

	lujing := mulu[gen.Appmodels].Zhi
	kongge := fuhao[gen.Kongge].Zhi
	huanhang := fuhao[gen.Huanhang].Zhi

	buffer.WriteString(gen.Package)
	buffer.WriteString(kongge)
	buffer.WriteString(lujing)
	buffer.WriteString(huanhang)
	buffer.WriteString(gen.Type)
	buffer.WriteString(kongge)
	buffer.WriteString(bianma)
	buffer.WriteString(kongge)
	buffer.WriteString(gen.Struct)
	buffer.WriteString(kongge)
	buffer.WriteString(fuhao[gen.Dakuohaozuo].Zhi)
	buffer.WriteString(huanhang)
}

func Shengcheng_yingyongzi_model() {
	gen := baseinits.Gen
	buffer := bytes.Buffer{}
	Buffertou(&buffer, gen.Yingyongzi)
	for k := range baseinits.Zifus {
		buffer.WriteString(k)
		buffer.WriteString(baseinits.Fuhaos[gen.Kongge].Zhi)
		buffer.WriteString(gen.String)
		buffer.WriteString(baseinits.Fuhaos[gen.Huanhang].Zhi)
	}
	buffer.WriteString(baseinits.Fuhaos[gen.Dakuohaoyou].Zhi)
	path := baseinits.Mulus[gen.Appmodels].Zhi + baseinits.Fuhaos[gen.Xiexian].Zhi + gen.Yingyongzi + baseinits.Fuhaos[gen.Dianhao].Zhi + gen.Go
	ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
}
