package jueseservices

import (
	"bytes"
	"hanfuxin/allerrors"
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/juesedaos"
	"log"
	"strconv"
	"time"
)

func shengchengerrorchangdu(leixing string, lenchangliang int64, lenshiji int) string {
	buffer := bytes.Buffer{}
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)
	buffer.WriteString(appinits.Yingyongzi.Leixing)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Maohao].Zhi)
	buffer.WriteString(leixing)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Douhao].Zhi)
	buffer.WriteString("zuichang")
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Maohao].Zhi)
	buffer.WriteString(strconv.FormatInt(lenchangliang, 10))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Douhao].Zhi)
	buffer.WriteString("shiji")
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Maohao].Zhi)
	buffer.WriteString(strconv.Itoa(lenshiji))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)
	return buffer.String()
}

func yanzhengziduanchangdu(juese *appmodels.Juese) error {
	lenbianma, _ := strconv.ParseInt(baseinits.Lies[appinits.Yingyongzi.Bianma].Changdu, 10, 0)
	lenmingcheng, _ := strconv.ParseInt(baseinits.Lies[appinits.Yingyongzi.Mingcheng].Changdu, 10, 0)
	lenbiaoji, _ := strconv.ParseInt(baseinits.Lies[appinits.Yingyongzi.Biaoji].Changdu, 10, 0)
	lenbianmashiti := len(juese.Bianma)
	lenmingchengshiti := len(juese.Mingcheng)
	lenbiaojishiti := len(juese.Biaoji)
	if lenbianmashiti > int(lenbianma) {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: shengchengerrorchangdu(appinits.Yingyongzi.Bianma, lenbianma, lenbianmashiti)}
	}
	if lenmingchengshiti > int(lenmingcheng) {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: shengchengerrorchangdu(appinits.Yingyongzi.Mingcheng, lenmingcheng, lenmingchengshiti)}
	}
	if lenbiaojishiti > int(lenbiaoji) {
		return allerrors.Ziduanerror{Shijian: time.Now(), Wenti: shengchengerrorchangdu(appinits.Yingyongzi.Biaoji, lenbiaoji, lenbiaojishiti)}
	}
	return nil
}

func Tianjiajuese(juese *appmodels.Juese) {
	err := yanzhengziduanchangdu(juese)
	if err != nil {
		log.Println(err)
		return
	}
	juesedaos.Tianjia_yige_juese(juese)
}
