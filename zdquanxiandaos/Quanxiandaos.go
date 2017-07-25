package zdquanxiandaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/zf"
)

func Chaxunyige(id int) *appmodels.Quanxian {
	quanxian := &appmodels.Quanxian{Id: id}
	err := appinits.Hanfuxinormer.Read(quanxian)
	if err != nil {
		return nil
	}
	return quanxian
}
func Tianjiayige(quanxian *appmodels.Quanxian) string {
	_, err := appinits.Hanfuxinormer.Insert(quanxian)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(quanxianshuzu []appmodels.Quanxian) string {
	_, err := appinits.Hanfuxinormer.InsertMulti(len(quanxianshuzu), quanxianshuzu)
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
func Xiugaiyige(quanxian *appmodels.Quanxian) string {
	_, err := appinits.Hanfuxinormer.Update(quanxian)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
