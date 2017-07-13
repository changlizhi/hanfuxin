package zdxinxijuesedaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
)

func Chaxunyige(id int) *appmodels.Xinxijuese {
	xinxijuese := &appmodels.Xinxijuese{Id: id}
	appinits.Hanfuxinormer.Read(xinxijuese)
	return xinxijuese
}
func Tianjiayige(xinxijuese *appmodels.Xinxijuese) {
	appinits.Hanfuxinormer.Insert(xinxijuese)
}
func Tianjiaduoge(xinxijueseshuzu []appmodels.Xinxijuese) {
	appinits.Hanfuxinormer.InsertMulti(len(xinxijueseshuzu), xinxijueseshuzu)
}
func Shanchuyige(id int) {
	appinits.Hanfuxinormer.Delete(Chaxunyige(id))
}
func Xiugaiyige(xinxijuese *appmodels.Xinxijuese) {
	appinits.Hanfuxinormer.Update(xinxijuese)
}
