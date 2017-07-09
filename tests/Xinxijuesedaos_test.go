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
Juesebianma:"JuesebianmaTianjiayigeTest1",
Id:1,
Xinxibianma:"XinxibianmaTianjiayigeTest1",
Biaoji:"BiaojiTianjiayigeTest1",
}
zdxinxijuesedaos.Tianjiayige(xinxijuese)
}
func TestXiugaiyigeXinxijuese(t *testing.T){
xinxijuese:=&appmodels.Xinxijuese{
Xinxibianma:"XinxibianmaXiugaiyigeTest1",
Biaoji:"BiaojiXiugaiyigeTest1",
Juesebianma:"JuesebianmaXiugaiyigeTest1",
Id:1,
}
zdxinxijuesedaos.Xiugaiyige(xinxijuese)
}
func TestTianjiaduogeXinxijuese(t *testing.T){
xinxijuese2:=appmodels.Xinxijuese{
Biaoji:"BiaojiTianjiaduogeTest2",
Juesebianma:"JuesebianmaTianjiaduogeTest2",
Xinxibianma:"XinxibianmaTianjiaduogeTest2",
Id:1,
}
xinxijuese3:=appmodels.Xinxijuese{
Xinxibianma:"XinxibianmaTianjiaduogeTest3",
Id:1,
Biaoji:"BiaojiTianjiaduogeTest3",
Juesebianma:"JuesebianmaTianjiaduogeTest3",
}
xinxijueses:=[]appmodels.Xinxijuese{xinxijuese2,xinxijuese3}
zdxinxijuesedaos.Tianjiaduoge(xinxijueses)
}
