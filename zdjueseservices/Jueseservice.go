package zdjueseservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appmodels"
	"hanfuxin/baserun"
	"hanfuxin/zdjuesedaos"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"strconv"
	"time"
)

func shengchengerrorchangdu(leixing string, lenchangliang int64, lenshiji int) string {
	buffer := bytes.Buffer{}
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	douhao := zfzhi.Douhaozhi()
	mh := zfzhi.Maohaozhi()

	buffer.WriteString(dkhz)
	buffer.WriteString(zf.Leixing(false))
	buffer.WriteString(mh)

	buffer.WriteString(leixing)
	buffer.WriteString(douhao)
	buffer.WriteString(zf.Zuichang(false))
	buffer.WriteString(mh)
	buffer.WriteString(strconv.FormatInt(lenchangliang, 10))
	buffer.WriteString(douhao)
	buffer.WriteString(zf.Shiji(false))
	buffer.WriteString(mh)
	buffer.WriteString(strconv.Itoa(lenshiji))
	buffer.WriteString(dkhy)
	return buffer.String()
}

func yanzhengziduanchangdu(juese *appmodels.Juese) error {
	zf := zf.Zf{}
	lenbianma := baserun.Huoquyigechangdu(zf.Bianma(false))
	//strconv.ParseInt(baseinits.Lies[appinits.Yingyongzi.Bianma].Changdu, 10, 0)
	lenmingcheng := baserun.Huoquyigechangdu(zf.Mingcheng(false))
	// strconv.ParseInt(baseinits.Lies[appinits.Yingyongzi.Mingcheng].Changdu, 10, 0)
	lenbiaoji := baserun.Huoquyigechangdu(zf.Biaoji(false))
	//strconv.ParseInt(baseinits.Lies[appinits.Yingyongzi.Biaoji].Changdu, 10, 0)
	lenbianmashiti := len(juese.Bianma)
	lenmingchengshiti := len(juese.Mingcheng)
	lenbiaojishiti := len(juese.Biaoji)
	if lenbianmashiti > int(lenbianma) {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: shengchengerrorchangdu(zf.Bianma(false), int64(lenbianma), lenbianmashiti)}
	}
	if lenmingchengshiti > int(lenmingcheng) {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: shengchengerrorchangdu(zf.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti)}
	}
	if lenbiaojishiti > int(lenbiaoji) {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: shengchengerrorchangdu(zf.Biaoji(false), int64(lenbiaoji), lenbiaojishiti)}
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
