package zdxinxidaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/zf"
)

func Chaxunyige(id int) *appmodels.Xinxi {
	xinxi := &appmodels.Xinxi{Id: id}
	err := appinits.Hanfuxinormer.Read(xinxi)
	if err != nil {
		return nil
	}
	return xinxi
}
func Tianjiayige(xinxi *appmodels.Xinxi) string {
	_, err := appinits.Hanfuxinormer.Insert(xinxi)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(xinxishuzu []appmodels.Xinxi) string {
	_, err := appinits.Hanfuxinormer.InsertMulti(len(xinxishuzu), xinxishuzu)
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
func Xiugaiyige(xinxi *appmodels.Xinxi) string {
	_, err := appinits.Hanfuxinormer.Update(xinxi)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
