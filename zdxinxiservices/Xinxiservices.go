package zdxinxiservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/appinits"
	"hanfuxin/baserun"
	"hanfuxin/zdxinxidaos"
	"changliang/zf"
	"changliang/zfzhi"
	"time"
)

func yanzhengziduanchangdu(xinxi *appmodels.Xinxi) error {
	cuowu := false
	buffer := bytes.Buffer{}

	lenyouxiang := baserun.Huoquyigechangdu(zf.Zfs.Youxiang(false))
	lenyouxiangshiti := len(xinxi.Youxiang)
	if lenyouxiangshiti > int(lenyouxiang) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenyouxiang), lenyouxiangshiti))
	}

	lenbianma := baserun.Huoquyigechangdu(zf.Zfs.Bianma(false))
	lenbianmashiti := len(xinxi.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Zfs.Biaoji(false))
	lenbiaojishiti := len(xinxi.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Zfs.Mingcheng(false))
	lenmingchengshiti := len(xinxi.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiaxinxi(xinxi *appmodels.Xinxi) string {
	err := yanzhengziduanchangdu(xinxi)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	return zdxinxidaos.Tianjiayige(xinxi)

}
func Xiugaixinxi(xinxi *appmodels.Xinxi) string {
	err := yanzhengziduanchangdu(xinxi)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	xinxifind := Chaxunxinxi(xinxi.Id)
	if xinxifind != nil {

		if xinxi.Youxiang != zfzhi.Zhi.Kzf() {
			xinxifind.Youxiang = xinxi.Youxiang
		}
		if xinxi.Biaoji != zfzhi.Zhi.Kzf() {
			xinxifind.Biaoji = xinxi.Biaoji
		}
		if xinxi.Bianma != zfzhi.Zhi.Kzf() {
			xinxifind.Bianma = xinxi.Bianma
		}
		if xinxi.Mingcheng != zfzhi.Zhi.Kzf() {
			xinxifind.Mingcheng = xinxi.Mingcheng
		}
		return zdxinxidaos.Xiugaiyige(xinxifind)

	}
	return appinits.Cuowus[zf.Zfs.Error04(false)].Zhi
}
func Shanchuxinxi(id int) string {
	return zdxinxidaos.Shanchuyige(id)
}
func Chaxunxinxi(id int) *appmodels.Xinxi {
	return zdxinxidaos.Chaxunyige(id)
}
