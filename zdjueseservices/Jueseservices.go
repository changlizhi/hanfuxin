package zdjueseservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/baserun"
	"hanfuxin/zdjuesedaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"time"
)

func yanzhengziduanchangdu(juese *appmodels.Juese) error {
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenbianma := baserun.Huoquyigechangdu(zf.Bianma(false))
	lenbianmashiti := len(juese.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(juese.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Mingcheng(false))
	lenmingchengshiti := len(juese.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiajuese(juese *appmodels.Juese) {
	err := yanzhengziduanchangdu(juese)
	if err != nil {
		log.Println(err)
		return
	}
	zdjuesedaos.Tianjiayige(juese)

}
func Xiugaijuese(juese *appmodels.Juese) {
	err := yanzhengziduanchangdu(juese)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	if err != nil {
		log.Println(err)
		return
	}
	juesefind := Chaxunjuese(juese.Id)
	if juesefind != nil {

		if juese.Mingcheng != kzf {
			juesefind.Mingcheng = juese.Mingcheng
		}
		if juese.Bianma != kzf {
			juesefind.Bianma = juese.Bianma
		}
		if juese.Biaoji != kzf {
			juesefind.Biaoji = juese.Biaoji
		}
		zdjuesedaos.Xiugaiyige(juesefind)
		return
	}
	log.Println(juese.Bianma, baseinits.Cuowus[zf.Error04(false)].Zhi)
}
func Chaxunjuese(id int) *appmodels.Juese {
	return zdjuesedaos.Chaxunyige(id)
}
func Shanchujuese(id int) {
	zdjuesedaos.Shanchuyige(id)
}
