package zdyanzhengleixingservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/baserun"
	"hanfuxin/zdyanzhengleixingdaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"time"
)

func yanzhengziduanchangdu(yanzhengleixing *appmodels.Yanzhengleixing) error {
	zf := zf.Zf{}
	cuowu := false
	buffer := bytes.Buffer{}

	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	lenbiaojishiti := len(yanzhengleixing.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenbianma := baserun.Huoquyigechangdu(zf.Bianma(false))
	lenbianmashiti := len(yanzhengleixing.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}

	lenmingcheng := baserun.Huoquyigechangdu(zf.Mingcheng(false))
	lenmingchengshiti := len(yanzhengleixing.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(apputils.Shengchengerrorchangdu(zf.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}
	if cuowu {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiayanzhengleixing(yanzhengleixing *appmodels.Yanzhengleixing) {
	err := yanzhengziduanchangdu(yanzhengleixing)
	if err != nil {
		log.Println(err)
		return
	}
	zdyanzhengleixingdaos.Tianjiayige(yanzhengleixing)

}
func Xiugaiyanzhengleixing(yanzhengleixing *appmodels.Yanzhengleixing) {
	err := yanzhengziduanchangdu(yanzhengleixing)
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	kzf := zfzhi.Kongzifuzhi()
	if err != nil {
		log.Println(err)
		return
	}
	yanzhengleixingfind := Chaxunyanzhengleixing(yanzhengleixing.Id)
	if yanzhengleixingfind != nil {

		if yanzhengleixing.Biaoji != kzf {
			yanzhengleixingfind.Biaoji = yanzhengleixing.Biaoji
		}
		if yanzhengleixing.Bianma != kzf {
			yanzhengleixingfind.Bianma = yanzhengleixing.Bianma
		}
		if yanzhengleixing.Mingcheng != kzf {
			yanzhengleixingfind.Mingcheng = yanzhengleixing.Mingcheng
		}
		zdyanzhengleixingdaos.Xiugaiyige(yanzhengleixingfind)
		return
	}
	log.Println(yanzhengleixing.Bianma, baseinits.Cuowus[zf.Error04(false)].Zhi)
}
func Chaxunyanzhengleixing(id int) *appmodels.Yanzhengleixing {
	return zdyanzhengleixingdaos.Chaxunyige(id)
}
func Shanchuyanzhengleixing(id int) {
	zdyanzhengleixingdaos.Shanchuyige(id)
}
