package zdjuesequanxianservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/baserun"
	"hanfuxin/zdjuesequanxiandaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"time"
)

func yanzhengziduanchangdu(juesequanxian *appmodels.Juesequanxian) error {
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenquanxianbianma := baserun.Huoquyigechangdu(zf.Quanxianbianma(false))
	lenquanxianbianmashiti := len(juesequanxian.Quanxianbianma)
	if lenquanxianbianmashiti > int(lenquanxianbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenquanxianbianma), lenquanxianbianmashiti))
	}

	lenjuesebianma := baserun.Huoquyigechangdu(zf.Juesebianma(false))
	lenjuesebianmashiti := len(juesequanxian.Juesebianma)
	if lenjuesebianmashiti > int(lenjuesebianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenjuesebianma), lenjuesebianmashiti))
	}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(juesequanxian.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiajuesequanxian(juesequanxian *appmodels.Juesequanxian) {
	err := yanzhengziduanchangdu(juesequanxian)
	if err != nil {
		log.Println(err)
		return
	}
	zdjuesequanxiandaos.Tianjiayige(juesequanxian)

}
func Xiugaijuesequanxian(juesequanxian *appmodels.Juesequanxian) {
	err := yanzhengziduanchangdu(juesequanxian)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	if err != nil {
		log.Println(err)
		return
	}
	juesequanxianfind := Chaxunjuesequanxian(juesequanxian.Id)
	if juesequanxianfind != nil {

		if juesequanxian.Juesebianma != kzf {
			juesequanxianfind.Juesebianma = juesequanxian.Juesebianma
		}
		if juesequanxian.Biaoji != kzf {
			juesequanxianfind.Biaoji = juesequanxian.Biaoji
		}
		if juesequanxian.Quanxianbianma != kzf {
			juesequanxianfind.Quanxianbianma = juesequanxian.Quanxianbianma
		}
		zdjuesequanxiandaos.Xiugaiyige(juesequanxianfind)
		return
	}
	log.Println(juesequanxian.Bianma, baseinits.Cuowus[zf.Error04(false)].Zhi)
}
func Chaxunjuesequanxian(id int) *appmodels.Juesequanxian {
	return zdjuesequanxiandaos.Chaxunyige(id)
}
func Shanchujuesequanxian(id int) {
	zdjuesequanxiandaos.Shanchuyige(id)
}
