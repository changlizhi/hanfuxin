package zdxinxijueseservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/baserun"
	"hanfuxin/zdxinxijuesedaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"time"
)

func yanzhengziduanchangdu(xinxijuese *appmodels.Xinxijuese) error {
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(xinxijuese.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenjuesebianma := baserun.Huoquyigechangdu(zf.Juesebianma(false))
	lenjuesebianmashiti := len(xinxijuese.Juesebianma)
	if lenjuesebianmashiti > int(lenjuesebianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenjuesebianma), lenjuesebianmashiti))
	}

	lenxinxibianma := baserun.Huoquyigechangdu(zf.Xinxibianma(false))
	lenxinxibianmashiti := len(xinxijuese.Xinxibianma)
	if lenxinxibianmashiti > int(lenxinxibianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenxinxibianma), lenxinxibianmashiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiaxinxijuese(xinxijuese *appmodels.Xinxijuese) {
	err := yanzhengziduanchangdu(xinxijuese)
	if err != nil {
		log.Println(err)
		return
	}
	zdxinxijuesedaos.Tianjiayige(xinxijuese)

}
func Xiugaixinxijuese(xinxijuese *appmodels.Xinxijuese) {
	err := yanzhengziduanchangdu(xinxijuese)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	if err != nil {
		log.Println(err)
		return
	}
	xinxijuesefind := Chaxunxinxijuese(xinxijuese.Id)
	if xinxijuesefind != nil {

		if xinxijuese.Juesebianma != kzf {
			xinxijuesefind.Juesebianma = xinxijuese.Juesebianma
		}
		if xinxijuese.Biaoji != kzf {
			xinxijuesefind.Biaoji = xinxijuese.Biaoji
		}
		if xinxijuese.Xinxibianma != kzf {
			xinxijuesefind.Xinxibianma = xinxijuese.Xinxibianma
		}
		zdxinxijuesedaos.Xiugaiyige(xinxijuesefind)
		return
	}
	log.Println(xinxijuese.Bianma, baseinits.Cuowus[zf.Error04(false)].Zhi)
}
func Chaxunxinxijuese(id int) *appmodels.Xinxijuese {
	return zdxinxijuesedaos.Chaxunyige(id)
}
func Shanchuxinxijuese(id int) {
	zdxinxijuesedaos.Shanchuyige(id)
}
