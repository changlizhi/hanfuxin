package zdjueseservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/appinits"
	"hanfuxin/baserun"
	"hanfuxin/zdjuesedaos"
	"changliang/zf"
	"changliang/zfzhi"
	"time"
)

func yanzhengziduanchangdu(juese *appmodels.Juese) error {
	cuowu := false
	buffer := bytes.Buffer{}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Zfs.Mingcheng(false))
	lenmingchengshiti := len(juese.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Zfs.Biaoji(false))
	lenbiaojishiti := len(juese.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenbianma := baserun.Huoquyigechangdu(zf.Zfs.Bianma(false))
	lenbianmashiti := len(juese.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiajuese(juese *appmodels.Juese) string {
	err := yanzhengziduanchangdu(juese)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	return zdjuesedaos.Tianjiayige(juese)

}
func Xiugaijuese(juese *appmodels.Juese) string {
	err := yanzhengziduanchangdu(juese)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	juesefind := Chaxunjuese(juese.Id)
	if juesefind != nil {

		if juese.Biaoji != zfzhi.Zhi.Kzf() {
			juesefind.Biaoji = juese.Biaoji
		}
		if juese.Bianma != zfzhi.Zhi.Kzf() {
			juesefind.Bianma = juese.Bianma
		}
		if juese.Mingcheng != zfzhi.Zhi.Kzf() {
			juesefind.Mingcheng = juese.Mingcheng
		}
		return zdjuesedaos.Xiugaiyige(juesefind)

	}
	return appinits.Cuowus[zf.Zfs.Error04(false)].Zhi
}
func Shanchujuese(id int) string {
	return zdjuesedaos.Shanchuyige(id)
}
func Chaxunjuese(id int) *appmodels.Juese {
	return zdjuesedaos.Chaxunyige(id)
}
