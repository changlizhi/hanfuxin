package zdquanxianservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/appinits"
	"hanfuxin/baserun"
	"hanfuxin/zdquanxiandaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"time"
)

func yanzhengziduanchangdu(quanxian *appmodels.Quanxian) error {
	cuowu := false
	buffer := bytes.Buffer{}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Zfs.Biaoji(false))
	lenbiaojishiti := len(quanxian.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Zfs.Mingcheng(false))
	lenmingchengshiti := len(quanxian.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}

	lenbianma := baserun.Huoquyigechangdu(zf.Zfs.Bianma(false))
	lenbianmashiti := len(quanxian.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}

	lenlujing := baserun.Huoquyigechangdu(zf.Zfs.Lujing(false))
	lenlujingshiti := len(quanxian.Lujing)
	if lenlujingshiti > int(lenlujing) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenlujing), lenlujingshiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiaquanxian(quanxian *appmodels.Quanxian) string {
	err := yanzhengziduanchangdu(quanxian)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	return zdquanxiandaos.Tianjiayige(quanxian)

}
func Xiugaiquanxian(quanxian *appmodels.Quanxian) string {
	err := yanzhengziduanchangdu(quanxian)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	quanxianfind := Chaxunquanxian(quanxian.Id)
	if quanxianfind != nil {

		if quanxian.Lujing != zfzhi.Zhi.Kzf() {
			quanxianfind.Lujing = quanxian.Lujing
		}
		if quanxian.Biaoji != zfzhi.Zhi.Kzf() {
			quanxianfind.Biaoji = quanxian.Biaoji
		}
		if quanxian.Mingcheng != zfzhi.Zhi.Kzf() {
			quanxianfind.Mingcheng = quanxian.Mingcheng
		}
		if quanxian.Bianma != zfzhi.Zhi.Kzf() {
			quanxianfind.Bianma = quanxian.Bianma
		}
		return zdquanxiandaos.Xiugaiyige(quanxianfind)

	}
	return appinits.Cuowus[zf.Zfs.Error04(false)].Zhi
}
func Shanchuquanxian(id int) string {
	return zdquanxiandaos.Shanchuyige(id)
}
func Chaxunquanxian(id int) *appmodels.Quanxian {
	return zdquanxiandaos.Chaxunyige(id)
}
