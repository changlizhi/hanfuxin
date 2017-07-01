package quanxiandaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
)

func Chaxunyige(id int) *appmodels.Quanxian {
	quanxian := &appmodels.Quanxian{Id: id}
	appinits.Hanfuxinormer.Read(quanxian)
	return quanxian
}
func Tianjiayige(quanxian *appmodels.Quanxian) {
	appinits.Hanfuxinormer.Insert(quanxian)
}
func Tianjiaduoge(quanxianshuzu []appmodels.Quanxian) {
	appinits.Hanfuxinormer.InsertMulti(len(quanxianshuzu), quanxianshuzu)
}
func Shanchuyige(id int) {
	appinits.Hanfuxinormer.Delete(Chaxunyige(id))
}
func Xiugaiyige(quanxian *appmodels.Quanxian) {
	appinits.Hanfuxinormer.Update(quanxian)
}
