package baserun

import (
	"bytes"
	"hanfuxin/appinits"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"log"
	"strings"
)

func havetime(bianma string) bool {
	for _, lie := range baseinits.Lies {
		pipei, _ := apputils.Pipei3lei(bianma, lie.Biaoming)
		if pipei && lie.Leixing == baseinits.Wenzis[appinits.Yingyongzi.Time].Zhi {
			return true
		}
	}
	return false
}

func testsimports(bianma string, buffer *bytes.Buffer) {
	buffer.WriteString(baseinits.Gen.Package)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Tests].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Gen.Import)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	if havetime(bianma) {
		buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
		buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Time].Zhi)
		buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
		buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	}

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Hanfuxin].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiexian].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appmodels].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Testing].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Log].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Hanfuxin].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiexian].Zhi)
	baoming := strings.ToLower(bianma) + baseinits.Mulus[appinits.Yingyongzi.Daos].Zhi
	buffer.WriteString(baoming)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
}

func testchaxunyige(bianma string, buffer *bytes.Buffer) {
	buffer.WriteString(baseinits.Gen.Func)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	fangfaming := baseinits.Wenzis[appinits.Yingyongzi.Test].Bianma +
		baseinits.Wenzis[appinits.Yingyongzi.Chaxunyige].Bianma +
		bianma
	buffer.WriteString(fangfaming)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.T].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xinghao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Testing].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.T].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Maohao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dengyuhao].Zhi)
	baoming := strings.ToLower(bianma) + baseinits.Mulus[appinits.Yingyongzi.Daos].Zhi
	buffer.WriteString(baoming)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Chaxunyige].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Jichus[appinits.Yingyongzi.Shuzi1].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Log].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Println].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
}

func testtianjiayige(bianma string, buffer *bytes.Buffer) {
	buffer.WriteString(baseinits.Gen.Func)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	fangfaming := baseinits.Wenzis[appinits.Yingyongzi.Test].Bianma +
		baseinits.Wenzis[appinits.Yingyongzi.Tianjiayige].Bianma +
		bianma
	buffer.WriteString(fangfaming)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.T].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xinghao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Testing].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.T].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Maohao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dengyuhao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Qiehao].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appmodels].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	pingjieziduan(bianma, buffer)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
}
func pingjieziduan(bianma string, buffer *bytes.Buffer) {
	for _, lie := range baseinits.Lies {
		pipei, _ := apputils.Pipei3lei(bianma, lie.Biaoming)
		if pipei {
			buffer.WriteString(lie.Bianma)
			buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Maohao].Zhi)
			zhi := shengchengzhi(lie.Bianma, lie.Leixing)
			buffer.WriteString(zhi)
			buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Douhao].Zhi)
			buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
		}
	}
}
func shengchengzhi(lieming string, leixing string) string {
	if leixing == baseinits.Gen.String {
		ret := baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi +
			lieming +
			baseinits.Wenzis[appinits.Yingyongzi.Test].Bianma +
			baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi
		return ret
	}
	if leixing == baseinits.Gen.Int {
		ret := baseinits.Jichus[appinits.Yingyongzi.Shuzi1].Zhi
		return ret
	}
	if leixing == baseinits.Wenzis[appinits.Yingyongzi.Time].Zhi {
		ret := baseinits.Wenzis[appinits.Yingyongzi.Time].Zhi +
			baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi +
			baseinits.Wenzis[appinits.Yingyongzi.Now].Bianma +
			baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi +
			baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi
		return ret
	}
	if leixing == baseinits.Gen.Float32 || leixing == baseinits.Gen.Float64 {
		ret := baseinits.Jichus[appinits.Yingyongzi.Shuzi1].Zhi +
			baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi +
			baseinits.Jichus[appinits.Yingyongzi.Shuzi0].Zhi
		return ret
	}
	return baseinits.Wenzis[appinits.Yingyongzi.Null].Bianma
}

func testshanchuyige(bianma string, buffer *bytes.Buffer) {
	buffer.WriteString(baseinits.Gen.Func)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	fangfaming := baseinits.Wenzis[appinits.Yingyongzi.Test].Bianma +
		baseinits.Wenzis[appinits.Yingyongzi.Shanchuyige].Bianma +
		bianma
	buffer.WriteString(fangfaming)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.T].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xinghao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Testing].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.T].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	baoming := strings.ToLower(bianma) + baseinits.Mulus[appinits.Yingyongzi.Daos].Zhi
	buffer.WriteString(baoming)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Shanchuyige].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Jichus[appinits.Yingyongzi.Shuzi1].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
}
func Shengchengdaostests() {
	for _, biao := range baseinits.Biaos {
		buffer := bytes.Buffer{}
		testsimports(biao.Bianma, &buffer)
		testchaxunyige(biao.Bianma, &buffer)
		testtianjiayige(biao.Bianma, &buffer)
		testshanchuyige(biao.Bianma, &buffer)

		log.Println("buffer-------", buffer.String())
	}
}
