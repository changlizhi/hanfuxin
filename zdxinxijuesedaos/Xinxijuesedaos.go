package zdxinxijuesedaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/zf"
)

func Chaxunyige(id int) *appmodels.Xinxijuese {
	xinxijuese := &appmodels.Xinxijuese{Id: id}
	err := appinits.Hanfuxinormer.Read(xinxijuese)
	if err != nil {
		return nil
	}
	return xinxijuese
}
func Tianjiayige(xinxijuese *appmodels.Xinxijuese) string {
	_, err := appinits.Hanfuxinormer.Insert(xinxijuese)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(xinxijueseshuzu []appmodels.Xinxijuese) string {
	_, err := appinits.Hanfuxinormer.InsertMulti(len(xinxijueseshuzu), xinxijueseshuzu)
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
func Xiugaiyige(xinxijuese *appmodels.Xinxijuese) string {
	_, err := appinits.Hanfuxinormer.Update(xinxijuese)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
