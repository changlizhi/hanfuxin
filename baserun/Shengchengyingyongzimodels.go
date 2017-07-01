package baserun

// 在应用启动之前生成一些必要的model
import (
	"bytes"
	"hanfuxin/baseinits"
	"io/ioutil"
	"os"
)

func Buffertoumodel(buffer *bytes.Buffer, bianma string) {
	buffer.WriteString(baseinits.Gen.Package)
	buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Konggefu].Zhi)
	buffer.WriteString(baseinits.Mulus[baseinits.Gen.Yingyongzimodel].Zhi)
	buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Gen.Type)
	buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Konggefu].Zhi)
	buffer.WriteString(bianma)
	buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Konggefu].Zhi)
	buffer.WriteString(baseinits.Gen.Struct)
	buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Dakuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Huanhangfu].Zhi)
}

func Shengcheng_yingyongzi_model() {
	buffer := bytes.Buffer{}
	Buffertoumodel(&buffer, baseinits.Gen.Yingyongzi)
	for k := range baseinits.Zifus {
		buffer.WriteString(k)
		buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Konggefu].Zhi)
		buffer.WriteString(baseinits.Gen.String)
		buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Huanhangfu].Zhi)
	}
	buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Dakuohaoyou].Zhi)
	path := baseinits.Mulus[baseinits.Gen.Yingyongzimodel].Zhi +
		baseinits.Fuhaos[baseinits.Gen.Xiexian].Zhi +
		baseinits.Gen.Yingyongzi +
		baseinits.Fuhaos[baseinits.Gen.Dianhao].Zhi +
		baseinits.Gen.Go
	ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
}
