package xinxidaos
import (
"hanfuxin/appinits"
"hanfuxin/appmodels"
)
func Chaxunyige(id int) *appmodels.Xinxi {
xinxi := &appmodels.Xinxi{Id:id}
appinits.Hanfuxinormer.Read(xinxi)
return xinxi
}
func Tianjiayige(xinxi *appmodels.Xinxi){
appinits.Hanfuxinormer.Insert(xinxi)
}
func Tianjiaduoge(xinxishuzu []appmodels.Xinxi){ 
appinits.Hanfuxinormer.InsertMulti(len(xinxishuzu),xinxishuzu)
}
func Shanchuyige(id int) {
appinits.Hanfuxinormer.Delete(Chaxunyige(id))
}
func Xiugaiyige(xinxi *appmodels.Xinxi){
appinits.Hanfuxinormer.Update(xinxi)
}