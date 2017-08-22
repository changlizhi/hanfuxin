package zdjuesedaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"changliang/zf"
	"log"
)

func Chaxunyige(id int) *appmodels.Juese {
	juese := &appmodels.Juese{Id: id}
	err := appinits.Defaultormer().Read(juese)
	if err != nil {
		log.Println(err)
		return nil
	}
	return juese
}
func Tianjiayige(juese *appmodels.Juese) string {
	_, err := appinits.Defaultormer().Insert(juese)
	if err != nil {
		log.Println(err)
		return appinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(jueseshuzu []appmodels.Juese) string {
	_, err := appinits.Defaultormer().InsertMulti(len(jueseshuzu), jueseshuzu)
	if err != nil {
		log.Println(err)
		return appinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Shanchuyige(id int) string {
	_, err := appinits.Defaultormer().Delete(Chaxunyige(id))
	if err != nil {
		log.Println(err)
		return appinits.Tishis[zf.Zfs.Tishi08(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi07(false)].Bianma
}
func Xiugaiyige(juese *appmodels.Juese) string {
	_, err := appinits.Defaultormer().Update(juese)
	if err != nil {
		log.Println(err)
		return appinits.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
