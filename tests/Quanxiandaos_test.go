package tests
import(
"hanfuxin/appmodels"
"testing"
"log"
"hanfuxin/zdquanxiandaos"
)
func TestChaxunyigeQuanxian(t *testing.T){
quanxian:=zdquanxiandaos.Chaxunyige(1)
log.Println(quanxian)
}
func TestShanchuyigeQuanxian(t *testing.T){
zdquanxiandaos.Shanchuyige(1)
}
func TestTianjiayigeQuanxian(t *testing.T){
quanxian:=&appmodels.Quanxian{
Biaoji:"BiaojiTianjiayigeTest1",
Lujing:"LujingTianjiayigeTest1",
Id:1,
Bianma:"BianmaTianjiayigeTest1",
Mingcheng:"MingchengTianjiayigeTest1",
}
zdquanxiandaos.Tianjiayige(quanxian)
}
func TestXiugaiyigeQuanxian(t *testing.T){
quanxian:=&appmodels.Quanxian{
Biaoji:"BiaojiXiugaiyigeTest1",
Lujing:"LujingXiugaiyigeTest1",
Id:1,
Bianma:"BianmaXiugaiyigeTest1",
Mingcheng:"MingchengXiugaiyigeTest1",
}
zdquanxiandaos.Xiugaiyige(quanxian)
}
func TestTianjiaduoge(t *testing.T){
quanxian2:=appmodels.Quanxian{
Biaoji:"BiaojiTianjiaduogeTest2",
Id:1,
Bianma:"BianmaTianjiaduogeTest2",
Mingcheng:"MingchengTianjiaduogeTest2",
Lujing:"LujingTianjiaduogeTest2",
}
quanxian3:=appmodels.Quanxian{
Id:1,
Bianma:"BianmaTianjiaduogeTest3",
Mingcheng:"MingchengTianjiaduogeTest3",
Lujing:"LujingTianjiaduogeTest3",
Biaoji:"BiaojiTianjiaduogeTest3",
}
quanxians:=[]appmodels.Quanxian{quanxian2,quanxian3}
zdquanxiandaos.Tianjiaduoge(quanxians)
}
