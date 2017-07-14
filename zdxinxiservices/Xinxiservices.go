package zdxinxiservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/baserun"
	"hanfuxin/zdxinxidaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"time"
)

func yanzhengziduanchangdu(xinxi *appmodels.Xinxi) error {
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Mingcheng(false))
	lenmingchengshiti := len(xinxi.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(xinxi.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenbianma := baserun.Huoquyigechangdu(zf.Bianma(false))
	lenbianmashiti := len(xinxi.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}

	lenyouxiang := baserun.Huoquyigechangdu(zf.Youxiang(false))
	lenyouxiangshiti := len(xinxi.Youxiang)
	if lenyouxiangshiti > int(lenyouxiang) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenyouxiang), lenyouxiangshiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiaxinxi(xinxi *appmodels.Xinxi) string {
	err := yanzhengziduanchangdu(xinxi)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	return zdxinxidaos.Tianjiayige(xinxi)

}
func Xiugaixinxi(xinxi *appmodels.Xinxi) string {
	err := yanzhengziduanchangdu(xinxi)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	xinxifind := Chaxunxinxi(xinxi.Id)
	if xinxifind != nil {

		if xinxi.Mingcheng != kzf {
			xinxifind.Mingcheng = xinxi.Mingcheng
		}
		if xinxi.Biaoji != kzf {
			xinxifind.Biaoji = xinxi.Biaoji
		}
		if xinxi.Youxiang != kzf {
			xinxifind.Youxiang = xinxi.Youxiang
		}
		if xinxi.Bianma != kzf {
			xinxifind.Bianma = xinxi.Bianma
		}
		return zdxinxidaos.Xiugaiyige(xinxifind)

	}
	return baseinits.Cuowus[zf.Error04(false)].Zhi
}
func Shanchuxinxi(id int) string {
	return zdxinxidaos.Shanchuyige(id)
}
func Chaxunxinxi(id int) *appmodels.Xinxi {
	return zdxinxidaos.Chaxunyige(id)
}
