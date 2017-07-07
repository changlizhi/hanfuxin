package tests
import(
"hanfuxin/appmodels"
"testing"
"log"
"hanfuxin/zdjuesedaos"
)
func TestChaxunyigeJuese(t *testing.T){
juese:=zdjuesedaos.Chaxunyige(1)
log.Println(juese)
}
func TestShanchuyigeJuese(t *testing.T){
zdjuesedaos.Shanchuyige(1)
}
func TestTianjiayigeJuese(t *testing.T){
juese:=&appmodels.Juese{
Biaoji:"BiaojiTianjiayigeTest1",
Mingcheng:"MingchengTianjiayigeTest1",
Id:1,
Bianma:"BianmaTianjiayigeTest1",
}
zdjuesedaos.Tianjiayige(juese)
}
func TestXiugaiyigeJuese(t *testing.T){
juese:=&appmodels.Juese{
Id:1,
Bianma:"BianmaXiugaiyigeTest1",
Mingcheng:"MingchengXiugaiyigeTest1",
Biaoji:"BiaojiXiugaiyigeTest1",
}
zdjuesedaos.Xiugaiyige(juese)
}
func TestTianjiaduoge(t *testing.T){
juese2:=appmodels.Juese{
Id:1,
Bianma:"BianmaTianjiaduogeTest2",
Mingcheng:"MingchengTianjiaduogeTest2",
Biaoji:"BiaojiTianjiaduogeTest2",
}
juese3:=appmodels.Juese{
Biaoji:"BiaojiTianjiaduogeTest3",
Bianma:"BianmaTianjiaduogeTest3",
Mingcheng:"MingchengTianjiaduogeTest3",
Id:1,
}
jueses:=[]appmodels.Juese{juese2,juese3}
zdjuesedaos.Tianjiaduoge(jueses)
}
