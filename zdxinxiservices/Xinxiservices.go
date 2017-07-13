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
	"log"
	"time"
)

func yanzhengziduanchangdu(xinxi *appmodels.Xinxi) error {
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(xinxi.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Mingcheng(false))
	lenmingchengshiti := len(xinxi.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
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
func Tianjiaxinxi(xinxi *appmodels.Xinxi) {
	err := yanzhengziduanchangdu(xinxi)
	if err != nil {
		log.Println(err)
		return
	}
	zdxinxidaos.Tianjiayige(xinxi)

}
func Xiugaixinxi(xinxi *appmodels.Xinxi) {
	err := yanzhengziduanchangdu(xinxi)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	if err != nil {
		log.Println(err)
		return
	}
	xinxifind := Chaxunxinxi(xinxi.Id)
	if xinxifind != nil {

		if xinxi.Bianma != kzf {
			xinxifind.Bianma = xinxi.Bianma
		}
		if xinxi.Mingcheng != kzf {
			xinxifind.Mingcheng = xinxi.Mingcheng
		}
		if xinxi.Biaoji != kzf {
			xinxifind.Biaoji = xinxi.Biaoji
		}
		if xinxi.Youxiang != kzf {
			xinxifind.Youxiang = xinxi.Youxiang
		}
		zdxinxidaos.Xiugaiyige(xinxifind)
		return
	}
	log.Println(xinxi.Bianma, baseinits.Cuowus[zf.Error04(false)].Zhi)
}
func Chaxunxinxi(id int) *appmodels.Xinxi {
	return zdxinxidaos.Chaxunyige(id)
}
func Shanchuxinxi(id int) {
	zdxinxidaos.Shanchuyige(id)
}
