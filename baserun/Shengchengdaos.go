package baserun

import (
	"bytes"
	"hanfuxin/appinits"
	"hanfuxin/baseinits"
	"hanfuxin/basemodels"
	"io/ioutil"
	"os"
	"strings"
)

func imports(bianma string, buffer *bytes.Buffer) {

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Gen.Import)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Hanfuxin].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiexian].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appinits].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Hanfuxin].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiexian].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appmodels].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
}
func chaxunyige(bianma string, buffer *bytes.Buffer) {
	buffer.WriteString(baseinits.Gen.Func)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Chaxunyige].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Lies[appinits.Yingyongzi.Id].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Gen.Int)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xinghao].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appmodels].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Maohao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dengyuhao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Qiehao].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appmodels].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)
	buffer.WriteString(baseinits.Lies[appinits.Yingyongzi.Id].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Maohao].Zhi)
	buffer.WriteString(baseinits.Lies[appinits.Yingyongzi.Id].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appinits].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Hanfuxinormer].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Read].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Gen.Return)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(strings.ToLower(bianma))

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
}
func tianjiayige(bianma string, buffer *bytes.Buffer) {
	buffer.WriteString(baseinits.Gen.Func)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Tianjiayige].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xinghao].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appmodels].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appinits].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Hanfuxinormer].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Insert].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
}
func tianjiaduoge(bianma string, buffer *bytes.Buffer) {
	buffer.WriteString(baseinits.Gen.Func)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Tianjiaduoge].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	canshu := strings.ToLower(bianma) + baseinits.Wenzis[appinits.Yingyongzi.Shuzu].Zhi
	buffer.WriteString(canshu)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Zhongkuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Zhongkuohaoyou].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appmodels].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appinits].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Hanfuxinormer].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.InsertMulti].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Gen.Len)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(canshu)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Douhao].Zhi)
	buffer.WriteString(canshu)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
}
func shanchuyige(bianma string, buffer *bytes.Buffer) {
	buffer.WriteString(baseinits.Gen.Func)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Shanchuyige].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Lies[appinits.Yingyongzi.Id].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Gen.Int)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appinits].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Hanfuxinormer].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Delete].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Chaxunyige].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Lies[appinits.Yingyongzi.Id].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
}
func xiugaiyige(bianma string, buffer *bytes.Buffer) {
	buffer.WriteString(baseinits.Gen.Func)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Xiugaiyige].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xinghao].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appmodels].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appinits].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Hanfuxinormer].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Update].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)

}
func Shengchengdaos() {
	for _, biao := range baseinits.Biaos {
		buffer := bytes.Buffer{}
		buffer.WriteString(baseinits.Gen.Package)
		buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
		lujing := strings.ToLower(biao.Bianma) + baseinits.Mulus[appinits.Yingyongzi.Daos].Zhi
		buffer.WriteString(lujing)

		imports(biao.Bianma, &buffer)
		chaxunyige(biao.Bianma, &buffer)
		tianjiayige(biao.Bianma, &buffer)
		tianjiaduoge(biao.Bianma, &buffer)
		shanchuyige(biao.Bianma, &buffer)
		xiugaiyige(biao.Bianma, &buffer)

		dir := basemodels.Getapppath() +
			baseinits.Fuhaos[appinits.Yingyongzi.Xiexian].Zhi +
			lujing +
			baseinits.Fuhaos[appinits.Yingyongzi.Xiexian].Zhi
		path := dir +
			biao.Bianma +
			baseinits.Mulus[appinits.Yingyongzi.Daos].Zhi +
			baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi +
			baseinits.Gen.Go
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
