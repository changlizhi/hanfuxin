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
	"log"
	"time"
)

func yanzhengziduanchangdu(yanzheng *appmodels.Yanzheng) error {
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenzhi := baserun.Huoquyigechangdu(zf.Zhi(false))
	lenzhishiti := len(yanzheng.Zhi)
	if lenzhishiti > int(lenzhi) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenzhi), lenzhishiti))
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

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(yanzheng.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenleixing := baserun.Huoquyigechangdu(zf.Leixing(false))
	lenleixingshiti := len(yanzheng.Leixing)
	if lenleixingshiti > int(lenleixing) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenleixing), lenleixingshiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiayanzheng(yanzheng *appmodels.Yanzheng) {
	err := yanzhengziduanchangdu(yanzheng)
	if err != nil {
		log.Println(err)
		return
	}
	zdyanzhengdaos.Tianjiayige(yanzheng)

}
func Xiugaiyanzheng(yanzheng *appmodels.Yanzheng) {
	err := yanzhengziduanchangdu(yanzheng)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	if err != nil {
		log.Println(err)
		return
	}
	yanzhengfind := Chaxunyanzheng(yanzheng.Id)
	if yanzhengfind != nil {

		if yanzheng.Bianma != kzf {
			yanzhengfind.Bianma = yanzheng.Bianma
		}
		if yanzheng.Leixing != kzf {
			yanzhengfind.Leixing = yanzheng.Leixing
		}
		if yanzheng.Biaoji != kzf {
			yanzhengfind.Biaoji = yanzheng.Biaoji
		}
		if yanzheng.Zhi != kzf {
			yanzhengfind.Zhi = yanzheng.Zhi
		}
		if yanzheng.Mingcheng != kzf {
			yanzhengfind.Mingcheng = yanzheng.Mingcheng
		}
		zdyanzhengdaos.Xiugaiyige(yanzhengfind)
		return
	}
	log.Println(yanzheng.Bianma, baseinits.Cuowus[zf.Error04(false)].Zhi)
}
func Chaxunyanzheng(id int) *appmodels.Yanzheng {
	return zdyanzhengdaos.Chaxunyige(id)
}
func Shanchuyanzheng(id int) {
	zdyanzhengdaos.Shanchuyige(id)
}
