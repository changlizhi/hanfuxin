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
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(yanzheng.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Mingcheng(false))
	lenmingchengshiti := len(yanzheng.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}

	lenbianma := baserun.Huoquyigechangdu(zf.Bianma(false))
	lenbianmashiti := len(yanzheng.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}

	lenleixing := baserun.Huoquyigechangdu(zf.Leixing(false))
	lenleixingshiti := len(yanzheng.Leixing)
	if lenleixingshiti > int(lenleixing) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenleixing), lenleixingshiti))
	}

	lenzhi := baserun.Huoquyigechangdu(zf.Zhi(false))
	lenzhishiti := len(yanzheng.Zhi)
	if lenzhishiti > int(lenzhi) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenzhi), lenzhishiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiayanzheng(yanzheng *appmodels.Yanzheng) string {
	err := yanzhengziduanchangdu(yanzheng)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	return zdyanzhengdaos.Tianjiayige(yanzheng)

}
func Xiugaiyanzheng(yanzheng *appmodels.Yanzheng) string {
	err := yanzhengziduanchangdu(yanzheng)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	xhx := zfzhi.Xiahuaxianzhi()
	if err != nil {
		return baseinits.Tishis[zf.Tishi09(false)].Bianma + xhx + err.Error()

	}
	yanzhengfind := Chaxunyanzheng(yanzheng.Id)
	if yanzhengfind != nil {

		if yanzheng.Biaoji != kzf {
			yanzhengfind.Biaoji = yanzheng.Biaoji
		}
		if yanzheng.Mingcheng != kzf {
			yanzhengfind.Mingcheng = yanzheng.Mingcheng
		}
		if yanzheng.Zhi != kzf {
			yanzhengfind.Zhi = yanzheng.Zhi
		}
		if yanzheng.Leixing != kzf {
			yanzhengfind.Leixing = yanzheng.Leixing
		}
		if yanzheng.Bianma != kzf {
			yanzhengfind.Bianma = yanzheng.Bianma
		}
		return zdyanzhengdaos.Xiugaiyige(yanzhengfind)

	}
	return baseinits.Cuowus[zf.Error04(false)].Zhi
}
func Shanchuyanzheng(id int) string {
	return zdyanzhengdaos.Shanchuyige(id)
}
func Chaxunyanzheng(id int) *appmodels.Yanzheng {
	return zdyanzhengdaos.Chaxunyige(id)
}
