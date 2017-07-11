package zdjuesedaos

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"log"
)

func Chaxunyige(id int) *appmodels.Juese {
	juese := &appmodels.Juese{Id: id}
	err := appinits.Hanfuxinormer.Read(juese)
	if err != nil {
		log.Println(err)
		return nil
	}
	return juese
}
func Tianjiayige(juese *appmodels.Juese) {
	appinits.Hanfuxinormer.Insert(juese)
}
func Tianjiaduoge(jueseshuzu []appmodels.Juese) {
	appinits.Hanfuxinormer.InsertMulti(len(jueseshuzu), jueseshuzu)
}
func Shanchuyige(id int) {
	appinits.Hanfuxinormer.Delete(Chaxunyige(id))
}
func Xiugaiyige(juese *appmodels.Juese) {
	appinits.Hanfuxinormer.Update(juese)
}
