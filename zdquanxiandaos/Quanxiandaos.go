package zdquanxiandaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"changliang/zf"
)

func Chaxunyige(id int) *appmodels.Quanxian {
	quanxian := &appmodels.Quanxian{Id: id}
	err := appinits.Defaultormer().Read(quanxian)
	if err != nil {
		return nil
	}
	return quanxian
}
func Tianjiayige(quanxian *appmodels.Quanxian) string {
	_, err := appinits.Defaultormer().Insert(quanxian)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(quanxianshuzu []appmodels.Quanxian) string {
	_, err := appinits.Defaultormer().InsertMulti(len(quanxianshuzu), quanxianshuzu)
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
func Xiugaiyige(quanxian *appmodels.Quanxian) string {
	_, err := appinits.Defaultormer().Update(quanxian)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
