package zdjuesequanxianservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/baserun"
	"hanfuxin/zdjuesequanxiandaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"time"
)

func yanzhengziduanchangdu(juesequanxian *appmodels.Juesequanxian) error {
	cuowu := false
	buffer := bytes.Buffer{}

	lenjuesebianma := baserun.Huoquyigechangdu(zf.Zfs.Juesebianma(false))
	lenjuesebianmashiti := len(juesequanxian.Juesebianma)
	if lenjuesebianmashiti > int(lenjuesebianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenjuesebianma), lenjuesebianmashiti))
	}

	lenquanxianbianma := baserun.Huoquyigechangdu(zf.Zfs.Quanxianbianma(false))
	lenquanxianbianmashiti := len(juesequanxian.Quanxianbianma)
	if lenquanxianbianmashiti > int(lenquanxianbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenquanxianbianma), lenquanxianbianmashiti))
	}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Zfs.Biaoji(false))
	lenbiaojishiti := len(juesequanxian.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiajuesequanxian(juesequanxian *appmodels.Juesequanxian) string {
	err := yanzhengziduanchangdu(juesequanxian)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	return zdjuesequanxiandaos.Tianjiayige(juesequanxian)

}
func Xiugaijuesequanxian(juesequanxian *appmodels.Juesequanxian) string {
	err := yanzhengziduanchangdu(juesequanxian)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	juesequanxianfind := Chaxunjuesequanxian(juesequanxian.Id)
	if juesequanxianfind != nil {

		if juesequanxian.Juesebianma != zfzhi.Zhi.Kzf() {
			juesequanxianfind.Juesebianma = juesequanxian.Juesebianma
		}
		if juesequanxian.Quanxianbianma != zfzhi.Zhi.Kzf() {
			juesequanxianfind.Quanxianbianma = juesequanxian.Quanxianbianma
		}
		if juesequanxian.Biaoji != zfzhi.Zhi.Kzf() {
			juesequanxianfind.Biaoji = juesequanxian.Biaoji
		}
		return zdjuesequanxiandaos.Xiugaiyige(juesequanxianfind)

	}
	return baseinits.Cuowus[zf.Zfs.Error04(false)].Zhi
}
func Shanchujuesequanxian(id int) string {
	return zdjuesequanxiandaos.Shanchuyige(id)
}
func Chaxunjuesequanxian(id int) *appmodels.Juesequanxian {
	return zdjuesequanxiandaos.Chaxunyige(id)
}
