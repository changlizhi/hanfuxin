package zdyanzhengleixingservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/baserun"
	"hanfuxin/zdyanzhengleixingdaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"time"
)

func yanzhengziduanchangdu(yanzhengleixing *appmodels.Yanzhengleixing) error {
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(yanzhengleixing.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Mingcheng(false))
	lenmingchengshiti := len(yanzhengleixing.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}

	lenbianma := baserun.Huoquyigechangdu(zf.Bianma(false))
	lenbianmashiti := len(yanzhengleixing.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiayanzhengleixing(yanzhengleixing *appmodels.Yanzhengleixing) string {
	err := yanzhengziduanchangdu(yanzhengleixing)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	return zdyanzhengleixingdaos.Tianjiayige(yanzhengleixing)

}
func Xiugaiyanzhengleixing(yanzhengleixing *appmodels.Yanzhengleixing) string {
	err := yanzhengziduanchangdu(yanzhengleixing)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	yanzhengleixingfind := Chaxunyanzhengleixing(yanzhengleixing.Id)
	if yanzhengleixingfind != nil {

		if yanzhengleixing.Biaoji != kzf {
			yanzhengleixingfind.Biaoji = yanzhengleixing.Biaoji
		}
		if yanzhengleixing.Mingcheng != kzf {
			yanzhengleixingfind.Mingcheng = yanzhengleixing.Mingcheng
		}
		if yanzhengleixing.Bianma != kzf {
			yanzhengleixingfind.Bianma = yanzhengleixing.Bianma
		}
		return zdyanzhengleixingdaos.Xiugaiyige(yanzhengleixingfind)

	}
	return baseinits.Cuowus[zf.Error04(false)].Zhi
}
func Shanchuyanzhengleixing(id int) string {
	return zdyanzhengleixingdaos.Shanchuyige(id)
}
func Chaxunyanzhengleixing(id int) *appmodels.Yanzhengleixing {
	return zdyanzhengleixingdaos.Chaxunyige(id)
}
