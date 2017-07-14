package zdjueseservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/baserun"
	"hanfuxin/zdjuesedaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"time"
)

func yanzhengziduanchangdu(juese *appmodels.Juese) error {
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Mingcheng(false))
	lenmingchengshiti := len(juese.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(juese.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenbianma := baserun.Huoquyigechangdu(zf.Bianma(false))
	lenbianmashiti := len(juese.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiajuese(juese *appmodels.Juese) string {
	err := yanzhengziduanchangdu(juese)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	return zdjuesedaos.Tianjiayige(juese)

}
func Xiugaijuese(juese *appmodels.Juese) string {
	err := yanzhengziduanchangdu(juese)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	juesefind := Chaxunjuese(juese.Id)
	if juesefind != nil {

		if juese.Biaoji != kzf {
			juesefind.Biaoji = juese.Biaoji
		}
		if juese.Mingcheng != kzf {
			juesefind.Mingcheng = juese.Mingcheng
		}
		if juese.Bianma != kzf {
			juesefind.Bianma = juese.Bianma
		}
		return zdjuesedaos.Xiugaiyige(juesefind)

	}
	return baseinits.Cuowus[zf.Error04(false)].Zhi
}
func Shanchujuese(id int) string {
	return zdjuesedaos.Shanchuyige(id)
}
func Chaxunjuese(id int) *appmodels.Juese {
	return zdjuesedaos.Chaxunyige(id)
}
