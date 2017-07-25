package zdyanzhengservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/baserun"
	"hanfuxin/zdyanzhengdaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"time"
)

func yanzhengziduanchangdu(yanzheng *appmodels.Yanzheng) error {
	cuowu := false
	buffer := bytes.Buffer{}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Zfs.Biaoji(false))
	lenbiaojishiti := len(yanzheng.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenleixing := baserun.Huoquyigechangdu(zf.Zfs.Leixing(false))
	lenleixingshiti := len(yanzheng.Leixing)
	if lenleixingshiti > int(lenleixing) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenleixing), lenleixingshiti))
	}

	lenzhi := baserun.Huoquyigechangdu(zf.Zfs.Zhi(false))
	lenzhishiti := len(yanzheng.Zhi)
	if lenzhishiti > int(lenzhi) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenzhi), lenzhishiti))
	}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Zfs.Mingcheng(false))
	lenmingchengshiti := len(yanzheng.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}

	lenbianma := baserun.Huoquyigechangdu(zf.Zfs.Bianma(false))
	lenbianmashiti := len(yanzheng.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiayanzheng(yanzheng *appmodels.Yanzheng) string {
	err := yanzhengziduanchangdu(yanzheng)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	return zdyanzhengdaos.Tianjiayige(yanzheng)

}
func Xiugaiyanzheng(yanzheng *appmodels.Yanzheng) string {
	err := yanzhengziduanchangdu(yanzheng)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	yanzhengfind := Chaxunyanzheng(yanzheng.Id)
	if yanzhengfind != nil {

		if yanzheng.Bianma != zfzhi.Zhi.Kzf() {
			yanzhengfind.Bianma = yanzheng.Bianma
		}
		if yanzheng.Zhi != zfzhi.Zhi.Kzf() {
			yanzhengfind.Zhi = yanzheng.Zhi
		}
		if yanzheng.Mingcheng != zfzhi.Zhi.Kzf() {
			yanzhengfind.Mingcheng = yanzheng.Mingcheng
		}
		if yanzheng.Leixing != zfzhi.Zhi.Kzf() {
			yanzhengfind.Leixing = yanzheng.Leixing
		}
		if yanzheng.Biaoji != zfzhi.Zhi.Kzf() {
			yanzhengfind.Biaoji = yanzheng.Biaoji
		}
		return zdyanzhengdaos.Xiugaiyige(yanzhengfind)

	}
	return baseinits.Cuowus[zf.Zfs.Error04(false)].Zhi
}
func Shanchuyanzheng(id int) string {
	return zdyanzhengdaos.Shanchuyige(id)
}
func Chaxunyanzheng(id int) *appmodels.Yanzheng {
	return zdyanzhengdaos.Chaxunyige(id)
}
