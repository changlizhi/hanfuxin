package zdjueseservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baserun"
	"hanfuxin/zdjuesedaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"time"
)

func yanzhengziduanchangdu(juese *appmodels.Juese) error {
	zf := zf.Zf{}
	lenbianma := baserun.Huoquyigechangdu(zf.Bianma(false))
	lenmingcheng := baserun.Huoquyigechangdu(zf.Mingcheng(false))
	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbianmashiti := len(juese.Bianma)
	lenmingchengshiti := len(juese.Mingcheng)
	lenbiaojishiti := len(juese.Biaoji)
	cuowu := false
	buffer := bytes.Buffer{}
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Bianma(false), int64(lenbianma), lenbianmashiti))
	}
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Biaoji(false), int64(lenbiaoji), lenbiaojishiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Chaxunjuese(id int) *appmodels.Juese {
	return zdjuesedaos.Chaxunyige(id)
}
func Shanchujuese(id int) {
	zdjuesedaos.Shanchuyige(id)
}
func Xiugaijuese(juese *appmodels.Juese) {
	err := yanzhengziduanchangdu(juese)
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	if err != nil {
		log.Println(err)
		return
	}
	juesefind := Chaxunjuese(juese.Id)
	if juesefind != nil {
		if juese.Bianma != kzf {
			juesefind.Bianma = juese.Bianma
		}
		if juese.Mingcheng != kzf {
			juesefind.Mingcheng = juese.Mingcheng
		}
		if juese.Biaoji != kzf {
			juesefind.Biaoji = juese.Biaoji
		}
		zdjuesedaos.Xiugaiyige(juesefind)
		return
	}
	log.Println("角色不存在！")
}
func Tianjiajuese(juese *appmodels.Juese) {
	err := yanzhengziduanchangdu(juese)
	if err != nil {
		log.Println(err)
		return
	}
	zdjuesedaos.Tianjiayige(juese)
}
