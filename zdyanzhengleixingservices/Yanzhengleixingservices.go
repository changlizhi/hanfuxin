package zdyanzhengleixingservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/appinits"
	"hanfuxin/baserun"
	"hanfuxin/zdyanzhengleixingdaos"
	"changliang/zf"
	"changliang/zfzhi"
	"time"
)

func yanzhengziduanchangdu(yanzhengleixing *appmodels.Yanzhengleixing) error {
	cuowu := false
	buffer := bytes.Buffer{}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Zfs.Mingcheng(false))
	lenmingchengshiti := len(yanzhengleixing.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Zfs.Biaoji(false))
	lenbiaojishiti := len(yanzhengleixing.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenbianma := baserun.Huoquyigechangdu(zf.Zfs.Bianma(false))
	lenbianmashiti := len(yanzhengleixing.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiayanzhengleixing(yanzhengleixing *appmodels.Yanzhengleixing) string {
	err := yanzhengziduanchangdu(yanzhengleixing)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	return zdyanzhengleixingdaos.Tianjiayige(yanzhengleixing)

}
func Xiugaiyanzhengleixing(yanzhengleixing *appmodels.Yanzhengleixing) string {
	err := yanzhengziduanchangdu(yanzhengleixing)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	yanzhengleixingfind := Chaxunyanzhengleixing(yanzhengleixing.Id)
	if yanzhengleixingfind != nil {

		if yanzhengleixing.Mingcheng != zfzhi.Zhi.Kzf() {
			yanzhengleixingfind.Mingcheng = yanzhengleixing.Mingcheng
		}
		if yanzhengleixing.Bianma != zfzhi.Zhi.Kzf() {
			yanzhengleixingfind.Bianma = yanzhengleixing.Bianma
		}
		if yanzhengleixing.Biaoji != zfzhi.Zhi.Kzf() {
			yanzhengleixingfind.Biaoji = yanzhengleixing.Biaoji
		}
		return zdyanzhengleixingdaos.Xiugaiyige(yanzhengleixingfind)

	}
	return appinits.Cuowus[zf.Zfs.Error04(false)].Zhi
}
func Shanchuyanzhengleixing(id int) string {
	return zdyanzhengleixingdaos.Shanchuyige(id)
}
func Chaxunyanzhengleixing(id int) *appmodels.Yanzhengleixing {
	return zdyanzhengleixingdaos.Chaxunyige(id)
}
