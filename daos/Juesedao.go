package daos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
)

func Tianjia_yige_juese(juese *appmodels.Juese) {
	appinits.Hanfuxinormer.Insert(juese)
}
func Tianjia_duoge_juese(jueses []appmodels.Juese) {
	appinits.Hanfuxinormer.InsertMulti(len(jueses), jueses)
}
