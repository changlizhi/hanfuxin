package baserun

import (
	"bytes"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"io/ioutil"
	"os"
)

func Shengcheng_yingyong_model() {
	biaos := baseinits.Biaos
	lies := baseinits.Lies
	for bk, bv := range biaos {
		buffer := bytes.Buffer{}
		Buffertoumodel(&buffer, bv.Bianma)
		for _, lv := range lies {
			pipei, _ := apputils.Pipei3lei(bk, lv.Biaoming)
			if pipei {
				buffer.WriteString(lv.Bianma)
				buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Kongge].Zhi)
				buffer.WriteString(lv.Leixing)
				buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Huanhang].Zhi)
			}
		}
		buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Dakuohaoyou].Zhi)
		path := baseinits.Mulus[baseinits.Gen.Appmodels].Zhi + baseinits.Fuhaos[baseinits.Gen.Xiexian].Zhi + bv.Bianma + baseinits.Fuhaos[baseinits.Gen.Dianhao].Zhi + baseinits.Gen.Go
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
