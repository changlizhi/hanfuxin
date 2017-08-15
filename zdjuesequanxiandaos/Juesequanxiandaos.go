package zdjuesequanxiandaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"changliang/zf"
)

func Chaxunyige(id int) *appmodels.Juesequanxian {
	juesequanxian := &appmodels.Juesequanxian{Id: id}
	err := appinits.Hanfuxinormer.Read(juesequanxian)
	if err != nil {
		return nil
	}
	return juesequanxian
}
func Tianjiayige(juesequanxian *appmodels.Juesequanxian) string {
	_, err := appinits.Hanfuxinormer.Insert(juesequanxian)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(juesequanxianshuzu []appmodels.Juesequanxian) string {
	_, err := appinits.Hanfuxinormer.InsertMulti(len(juesequanxianshuzu), juesequanxianshuzu)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Shanchuyige(id int) string {
	_, err := appinits.Hanfuxinormer.Delete(Chaxunyige(id))
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi08(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi07(false)].Bianma
}
func Xiugaiyige(juesequanxian *appmodels.Juesequanxian) string {
	_, err := appinits.Hanfuxinormer.Update(juesequanxian)
	if err != nil {
		return appinits.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return appinits.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
