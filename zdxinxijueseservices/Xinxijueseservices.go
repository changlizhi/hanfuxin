package zdxinxijueseservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/appinits"
	"hanfuxin/baserun"
	"hanfuxin/zdxinxijuesedaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"time"
)

func yanzhengziduanchangdu(xinxijuese *appmodels.Xinxijuese) error {
	cuowu := false
	buffer := bytes.Buffer{}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Zfs.Biaoji(false))
	lenbiaojishiti := len(xinxijuese.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenjuesebianma := baserun.Huoquyigechangdu(zf.Zfs.Juesebianma(false))
	lenjuesebianmashiti := len(xinxijuese.Juesebianma)
	if lenjuesebianmashiti > int(lenjuesebianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenjuesebianma), lenjuesebianmashiti))
	}

	lenxinxibianma := baserun.Huoquyigechangdu(zf.Zfs.Xinxibianma(false))
	lenxinxibianmashiti := len(xinxijuese.Xinxibianma)
	if lenxinxibianmashiti > int(lenxinxibianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenxinxibianma), lenxinxibianmashiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiaxinxijuese(xinxijuese *appmodels.Xinxijuese) string {
	err := yanzhengziduanchangdu(xinxijuese)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	return zdxinxijuesedaos.Tianjiayige(xinxijuese)

}
func Xiugaixinxijuese(xinxijuese *appmodels.Xinxijuese) string {
	err := yanzhengziduanchangdu(xinxijuese)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	xinxijuesefind := Chaxunxinxijuese(xinxijuese.Id)
	if xinxijuesefind != nil {

		if xinxijuese.Juesebianma != zfzhi.Zhi.Kzf() {
			xinxijuesefind.Juesebianma = xinxijuese.Juesebianma
		}
		if xinxijuese.Biaoji != zfzhi.Zhi.Kzf() {
			xinxijuesefind.Biaoji = xinxijuese.Biaoji
		}
		if xinxijuese.Xinxibianma != zfzhi.Zhi.Kzf() {
			xinxijuesefind.Xinxibianma = xinxijuese.Xinxibianma
		}
		return zdxinxijuesedaos.Xiugaiyige(xinxijuesefind)

	}
	return appinits.Cuowus[zf.Zfs.Error04(false)].Zhi
}
func Shanchuxinxijuese(id int) string {
	return zdxinxijuesedaos.Shanchuyige(id)
}
func Chaxunxinxijuese(id int) *appmodels.Xinxijuese {
	return zdxinxijuesedaos.Chaxunyige(id)
}
