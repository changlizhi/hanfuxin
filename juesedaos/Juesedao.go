package juesedaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
)

func Chaxun_yige(id int) *appmodels.Juese {
	juese := &appmodels.Juese{Id: id}
	appinits.Hanfuxinormer.Read(juese)
	return juese
}

func Tianjia_yige_juese(juese *appmodels.Juese) {
	appinits.Hanfuxinormer.Insert(juese)
}

func Tianjia_duoge_juese(jueses []appmodels.Juese) {
	appinits.Hanfuxinormer.InsertMulti(len(jueses), jueses)
}

func Shanchu_yige(id int) {
	appinits.Hanfuxinormer.Delete(Chaxun_yige(id))
}

func Xiugai_yige(juese *appmodels.Juese) {
	appinits.Hanfuxinormer.Update(juese)
}
