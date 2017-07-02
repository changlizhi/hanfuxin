package juesequanxiandaos
import (
"hanfuxin/appinits"
"hanfuxin/appmodels"
)
func Chaxunyige(id int) *appmodels.Juesequanxian {
juesequanxian := &appmodels.Juesequanxian{Id:id}
appinits.Hanfuxinormer.Read(juesequanxian)
return juesequanxian
}
func Tianjiayige(juesequanxian *appmodels.Juesequanxian){
appinits.Hanfuxinormer.Insert(juesequanxian)
}
func Tianjiaduoge(juesequanxianshuzu []appmodels.Juesequanxian){ 
appinits.Hanfuxinormer.InsertMulti(len(juesequanxianshuzu),juesequanxianshuzu)
}
func Shanchuyige(id int) {
appinits.Hanfuxinormer.Delete(Chaxunyige(id))
}
func Xiugaiyige(juesequanxian *appmodels.Juesequanxian){
appinits.Hanfuxinormer.Update(juesequanxian)
}