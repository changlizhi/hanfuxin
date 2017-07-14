package zdxinxijueseservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/baserun"
	"hanfuxin/zdxinxijuesedaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"time"
)

func yanzhengziduanchangdu(xinxijuese *appmodels.Xinxijuese) error {
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenjuesebianma := baserun.Huoquyigechangdu(zf.Juesebianma(false))
	lenjuesebianmashiti := len(xinxijuese.Juesebianma)
	if lenjuesebianmashiti > int(lenjuesebianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenjuesebianma), lenjuesebianmashiti))
	}

	lenxinxibianma := baserun.Huoquyigechangdu(zf.Xinxibianma(false))
	lenxinxibianmashiti := len(xinxijuese.Xinxibianma)
	if lenxinxibianmashiti > int(lenxinxibianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenxinxibianma), lenxinxibianmashiti))
	}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(xinxijuese.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiaxinxijuese(xinxijuese *appmodels.Xinxijuese) string {
	err := yanzhengziduanchangdu(xinxijuese)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	return zdxinxijuesedaos.Tianjiayige(xinxijuese)

}
func Xiugaixinxijuese(xinxijuese *appmodels.Xinxijuese) string {
	err := yanzhengziduanchangdu(xinxijuese)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	xinxijuesefind := Chaxunxinxijuese(xinxijuese.Id)
	if xinxijuesefind != nil {

		if xinxijuese.Xinxibianma != kzf {
			xinxijuesefind.Xinxibianma = xinxijuese.Xinxibianma
		}
		if xinxijuese.Juesebianma != kzf {
			xinxijuesefind.Juesebianma = xinxijuese.Juesebianma
		}
		if xinxijuese.Biaoji != kzf {
			xinxijuesefind.Biaoji = xinxijuese.Biaoji
		}
		return zdxinxijuesedaos.Xiugaiyige(xinxijuesefind)

	}
	return baseinits.Cuowus[zf.Error04(false)].Zhi
}
func Shanchuxinxijuese(id int) string {
	return zdxinxijuesedaos.Shanchuyige(id)
}
func Chaxunxinxijuese(id int) *appmodels.Xinxijuese {
	return zdxinxijuesedaos.Chaxunyige(id)
}
