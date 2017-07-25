package zdjuesedaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/zf"
)

func Chaxunyige(id int) *appmodels.Juese {
	juese := &appmodels.Juese{Id: id}
	err := appinits.Hanfuxinormer.Read(juese)
	if err != nil {
		return nil
	}
	return juese
}
func Tianjiayige(juese *appmodels.Juese) string {
	_, err := appinits.Hanfuxinormer.Insert(juese)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(jueseshuzu []appmodels.Juese) string {
	_, err := appinits.Hanfuxinormer.InsertMulti(len(jueseshuzu), jueseshuzu)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Shanchuyige(id int) string {
	_, err := appinits.Hanfuxinormer.Delete(Chaxunyige(id))
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi08(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi07(false)].Bianma
}
func Xiugaiyige(juese *appmodels.Juese) string {
	_, err := appinits.Hanfuxinormer.Update(juese)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
