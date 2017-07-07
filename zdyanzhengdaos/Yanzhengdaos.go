package zdyanzhengdaos
import (
"hanfuxin/appinits"
"hanfuxin/appmodels"
)
func Chaxunyige(id int) *appmodels.Yanzheng {
yanzheng := &appmodels.Yanzheng{Id:id}
appinits.Hanfuxinormer.Read(yanzheng)
return yanzheng
}
func Tianjiayige(yanzheng *appmodels.Yanzheng){
appinits.Hanfuxinormer.Insert(yanzheng)
}
func Tianjiaduoge(yanzhengshuzu []appmodels.Yanzheng){
appinits.Hanfuxinormer.InsertMulti(len(yanzhengshuzu),yanzhengshuzu)
}
func Shanchuyige(id int){
appinits.Hanfuxinormer.Delete(Chaxunyige(id))
}
func Xiugaiyige(yanzheng *appmodels.Yanzheng){
appinits.Hanfuxinormer.Update(yanzheng)
}
