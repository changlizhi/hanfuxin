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
	"log"
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

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(quanxian.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Mingcheng(false))
	lenmingchengshiti := len(quanxian.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
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
func Tianjiaquanxian(quanxian *appmodels.Quanxian) {
	err := yanzhengziduanchangdu(quanxian)
	if err != nil {
		log.Println(err)
		return
	}
	zdquanxiandaos.Tianjiayige(quanxian)

}
func Xiugaiquanxian(quanxian *appmodels.Quanxian) {
	err := yanzhengziduanchangdu(quanxian)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	if err != nil {
		log.Println(err)
		return
	}
	quanxianfind := Chaxunquanxian(quanxian.Id)
	if quanxianfind != nil {

		if quanxian.Lujing != kzf {
			quanxianfind.Lujing = quanxian.Lujing
		}
		if quanxian.Mingcheng != kzf {
			quanxianfind.Mingcheng = quanxian.Mingcheng
		}
		if quanxian.Biaoji != kzf {
			quanxianfind.Biaoji = quanxian.Biaoji
		}
		if quanxian.Bianma != kzf {
			quanxianfind.Bianma = quanxian.Bianma
		}
		zdquanxiandaos.Xiugaiyige(quanxianfind)
		return
	}
	log.Println(quanxian.Bianma, baseinits.Cuowus[zf.Error04(false)].Zhi)
}
func Chaxunquanxian(id int) *appmodels.Quanxian {
	return zdquanxiandaos.Chaxunyige(id)
}
func Shanchuquanxian(id int) {
	zdquanxiandaos.Shanchuyige(id)
}
