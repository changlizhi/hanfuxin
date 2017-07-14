package zdjuesequanxiandaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/zf"
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
	zf := zf.Zf{}
	_, err := appinits.Hanfuxinormer.Insert(juesequanxian)
	if err != nil {
		return baseinits.Tishis[zf.Tishi04(false)].Bianma
	}
	return baseinits.Tishis[zf.Tishi03(false)].Bianma
}
func Tianjiaduoge(juesequanxianshuzu []appmodels.Juesequanxian) string {
	zf := zf.Zf{}
	_, err := appinits.Hanfuxinormer.InsertMulti(len(juesequanxianshuzu), juesequanxianshuzu)
	if err != nil {
		return baseinits.Tishis[zf.Tishi04(false)].Bianma
	}
	return baseinits.Tishis[zf.Tishi03(false)].Bianma
}
func Shanchuyige(id int) string {
	zf := zf.Zf{}
	_, err := appinits.Hanfuxinormer.Delete(Chaxunyige(id))
	if err != nil {
		return baseinits.Tishis[zf.Tishi08(false)].Bianma
	}
	return baseinits.Tishis[zf.Tishi07(false)].Bianma
}
func Xiugaiyige(juesequanxian *appmodels.Juesequanxian) string {
	zf := zf.Zf{}
	_, err := appinits.Hanfuxinormer.Update(juesequanxian)
	if err != nil {
		return baseinits.Tishis[zf.Tishi06(false)].Bianma
	}
	return baseinits.Tishis[zf.Tishi05(false)].Bianma
}
