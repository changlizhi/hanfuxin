package zdxinxijuesedaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"changliang/zf"
)

func Chaxunyige(id int) *appmodels.Xinxijuese {
	xinxijuese := &appmodels.Xinxijuese{Id: id}
	err := appinits.Defaultormer().Read(xinxijuese)
	if err != nil {
		return nil
	}
	return xinxijuese
}
func Tianjiayige(xinxijuese *appmodels.Xinxijuese) string {
	_, err := appinits.Defaultormer().Insert(xinxijuese)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(xinxijueseshuzu []appmodels.Xinxijuese) string {
	_, err := appinits.Defaultormer().InsertMulti(len(xinxijueseshuzu), xinxijueseshuzu)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Shanchuyige(id int) string {
	_, err := appinits.Defaultormer().Delete(Chaxunyige(id))
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi08(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi07(false)].Bianma
}
func Xiugaiyige(xinxijuese *appmodels.Xinxijuese) string {
	_, err := appinits.Defaultormer().Update(xinxijuese)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
