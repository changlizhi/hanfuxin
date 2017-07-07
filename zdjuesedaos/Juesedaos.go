package zdjuesedaos
import (
"hanfuxin/appinits"
"hanfuxin/appmodels"
)
func Chaxunyige(id int) *appmodels.Juese {
juese := &appmodels.Juese{Id:id}
appinits.Hanfuxinormer.Read(juese)
return juese
}
func Tianjiayige(juese *appmodels.Juese){
appinits.Hanfuxinormer.Insert(juese)
}
func Tianjiaduoge(jueseshuzu []appmodels.Juese){
appinits.Hanfuxinormer.InsertMulti(len(jueseshuzu),jueseshuzu)
}
func Shanchuyige(id int){
appinits.Hanfuxinormer.Delete(Chaxunyige(id))
}
func Xiugaiyige(juese *appmodels.Juese){
appinits.Hanfuxinormer.Update(juese)
}
