package zdyanzhengleixingdaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/zf"
)

func Chaxunyige(id int) *appmodels.Yanzhengleixing {
	yanzhengleixing := &appmodels.Yanzhengleixing{Id: id}
	err := appinits.Hanfuxinormer.Read(yanzhengleixing)
	if err != nil {
		return nil
	}
	return yanzhengleixing
}
func Tianjiayige(yanzhengleixing *appmodels.Yanzhengleixing) string {
	_, err := appinits.Hanfuxinormer.Insert(yanzhengleixing)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(yanzhengleixingshuzu []appmodels.Yanzhengleixing) string {
	_, err := appinits.Hanfuxinormer.InsertMulti(len(yanzhengleixingshuzu), yanzhengleixingshuzu)
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
func Xiugaiyige(yanzhengleixing *appmodels.Yanzhengleixing) string {
	_, err := appinits.Hanfuxinormer.Update(yanzhengleixing)
	if err != nil {
		return baseinits.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return baseinits.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
