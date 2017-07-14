package zdquanxianservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/baserun"
	"hanfuxin/zdquanxiandaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"time"
)

func yanzhengziduanchangdu(quanxian *appmodels.Quanxian) error {
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenbianma := baserun.Huoquyigechangdu(zf.Bianma(false))
	lenbianmashiti := len(quanxian.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Mingcheng(false))
	lenmingchengshiti := len(quanxian.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(quanxian.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenlujing := baserun.Huoquyigechangdu(zf.Lujing(false))
	lenlujingshiti := len(quanxian.Lujing)
	if lenlujingshiti > int(lenlujing) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenlujing), lenlujingshiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiaquanxian(quanxian *appmodels.Quanxian) string {
	err := yanzhengziduanchangdu(quanxian)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	return zdquanxiandaos.Tianjiayige(quanxian)

}
func Xiugaiquanxian(quanxian *appmodels.Quanxian) string {
	err := yanzhengziduanchangdu(quanxian)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	quanxianfind := Chaxunquanxian(quanxian.Id)
	if quanxianfind != nil {

		if quanxian.Biaoji != kzf {
			quanxianfind.Biaoji = quanxian.Biaoji
		}
		if quanxian.Bianma != kzf {
			quanxianfind.Bianma = quanxian.Bianma
		}
		if quanxian.Mingcheng != kzf {
			quanxianfind.Mingcheng = quanxian.Mingcheng
		}
		if quanxian.Lujing != kzf {
			quanxianfind.Lujing = quanxian.Lujing
		}
		return zdquanxiandaos.Xiugaiyige(quanxianfind)

	}
	return baseinits.Cuowus[zf.Error04(false)].Zhi
}
func Shanchuquanxian(id int) string {
	return zdquanxiandaos.Shanchuyige(id)
}
func Chaxunquanxian(id int) *appmodels.Quanxian {
	return zdquanxiandaos.Chaxunyige(id)
}
