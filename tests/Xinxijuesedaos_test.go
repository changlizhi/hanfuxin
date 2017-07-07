package tests
import(
"hanfuxin/appmodels"
"testing"
"log"
"hanfuxin/zdxinxijuesedaos"
)
func TestChaxunyigeXinxijuese(t *testing.T){
xinxijuese:=zdxinxijuesedaos.Chaxunyige(1)
log.Println(xinxijuese)
}
func TestShanchuyigeXinxijuese(t *testing.T){
zdxinxijuesedaos.Shanchuyige(1)
}
func TestTianjiayigeXinxijuese(t *testing.T){
xinxijuese:=&appmodels.Xinxijuese{
Biaoji:"BiaojiTianjiayigeTest1",
Juesebianma:"JuesebianmaTianjiayigeTest1",
Id:1,
Xinxibianma:"XinxibianmaTianjiayigeTest1",
}
zdxinxijuesedaos.Tianjiayige(xinxijuese)
}
func TestXiugaiyigeXinxijuese(t *testing.T){
xinxijuese:=&appmodels.Xinxijuese{
Id:1,
Xinxibianma:"XinxibianmaXiugaiyigeTest1",
Biaoji:"BiaojiXiugaiyigeTest1",
Juesebianma:"JuesebianmaXiugaiyigeTest1",
}
zdxinxijuesedaos.Xiugaiyige(xinxijuese)
}
func TestTianjiaduoge(t *testing.T){
xinxijuese2:=appmodels.Xinxijuese{
Juesebianma:"JuesebianmaTianjiaduogeTest2",
Biaoji:"BiaojiTianjiaduogeTest2",
Id:1,
Xinxibianma:"XinxibianmaTianjiaduogeTest2",
}
xinxijuese3:=appmodels.Xinxijuese{
Biaoji:"BiaojiTianjiaduogeTest3",
Juesebianma:"JuesebianmaTianjiaduogeTest3",
Id:1,
Xinxibianma:"XinxibianmaTianjiaduogeTest3",
}
xinxijueses:=[]appmodels.Xinxijuese{xinxijuese2,xinxijuese3}
zdxinxijuesedaos.Tianjiaduoge(xinxijueses)
}
