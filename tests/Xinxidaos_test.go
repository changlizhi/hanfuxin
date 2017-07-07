package tests
import(
"hanfuxin/appmodels"
"testing"
"log"
"hanfuxin/zdxinxidaos"
)
func TestChaxunyigeXinxi(t *testing.T){
xinxi:=zdxinxidaos.Chaxunyige(1)
log.Println(xinxi)
}
func TestShanchuyigeXinxi(t *testing.T){
zdxinxidaos.Shanchuyige(1)
}
func TestTianjiayigeXinxi(t *testing.T){
xinxi:=&appmodels.Xinxi{
Bianma:"BianmaTianjiayigeTest1",
Mingcheng:"MingchengTianjiayigeTest1",
Youxiang:"YouxiangTianjiayigeTest1",
Id:1,
Biaoji:"BiaojiTianjiayigeTest1",
}
zdxinxidaos.Tianjiayige(xinxi)
}
func TestXiugaiyigeXinxi(t *testing.T){
xinxi:=&appmodels.Xinxi{
Biaoji:"BiaojiXiugaiyigeTest1",
Bianma:"BianmaXiugaiyigeTest1",
Mingcheng:"MingchengXiugaiyigeTest1",
Youxiang:"YouxiangXiugaiyigeTest1",
Id:1,
}
zdxinxidaos.Xiugaiyige(xinxi)
}
func TestTianjiaduoge(t *testing.T){
xinxi2:=appmodels.Xinxi{
Biaoji:"BiaojiTianjiaduogeTest2",
Bianma:"BianmaTianjiaduogeTest2",
Mingcheng:"MingchengTianjiaduogeTest2",
Youxiang:"YouxiangTianjiaduogeTest2",
Id:1,
}
xinxi3:=appmodels.Xinxi{
Biaoji:"BiaojiTianjiaduogeTest3",
Youxiang:"YouxiangTianjiaduogeTest3",
Id:1,
Bianma:"BianmaTianjiaduogeTest3",
Mingcheng:"MingchengTianjiaduogeTest3",
}
xinxis:=[]appmodels.Xinxi{xinxi2,xinxi3}
zdxinxidaos.Tianjiaduoge(xinxis)
}
