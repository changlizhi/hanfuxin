package baserun

import (
	"bytes"
	"hanfuxin/appinits"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/basemodels"
  //	"io/ioutil"
	//  "os"
	"log"
)

func Shengchenghanfuxinmodels() {
	biaos := baseinits.Biaos
	lies := baseinits.Lies
	for bk, bv := range biaos {
		buffer := bytes.Buffer{}
		Buffertoumodel(&buffer, bv.Bianma, baseinits.Mulus[appinits.Yingyongzi.Appmodels].Zhi)
		for _, lv := range lies {
			pipei, _ := apputils.Pipei3lei(bk, lv.Biaoming)
			if pipei {
				buffer.WriteString(lv.Bianma)
				buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Konggefu].Zhi)
				buffer.WriteString(lv.Leixing)
				buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Huanhangfu].Zhi)
			}
		}
		buffer.WriteString(baseinits.Fuhaos[baseinits.Gen.Dakuohaoyou].Zhi)
		path := basemodels.Getapppath() +
			baseinits.Fuhaos[baseinits.Gen.Xiexian].Zhi +
			baseinits.Mulus[baseinits.Gen.Appmodels].Zhi +
			baseinits.Fuhaos[baseinits.Gen.Xiexian].Zhi +
			bv.Bianma +
			baseinits.Fuhaos[baseinits.Gen.Dianhao].Zhi +
			baseinits.Gen.Go
		// ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		log.Println("appmodels-path---------",path)
	}
}
