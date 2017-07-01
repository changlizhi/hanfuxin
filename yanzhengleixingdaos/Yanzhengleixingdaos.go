package yanzhengleixingdaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
)

func Chaxunyige(id int) *appmodels.Yanzhengleixing {
	yanzhengleixing := &appmodels.Yanzhengleixing{Id: id}
	appinits.Hanfuxinormer.Read(yanzhengleixing)
	return yanzhengleixing
}
func Tianjiayige(yanzhengleixing *appmodels.Yanzhengleixing) {
	appinits.Hanfuxinormer.Insert(yanzhengleixing)
}
func Tianjiaduoge(yanzhengleixingshuzu []appmodels.Yanzhengleixing) {
	appinits.Hanfuxinormer.InsertMulti(len(yanzhengleixingshuzu), yanzhengleixingshuzu)
}
func Shanchuyige(id int) {
	appinits.Hanfuxinormer.Delete(Chaxunyige(id))
}
func Xiugaiyige(yanzhengleixing *appmodels.Yanzhengleixing) {
	appinits.Hanfuxinormer.Update(yanzhengleixing)
}
