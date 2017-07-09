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
Biaoji:"BiaojiTianjiayigeTest1",
Id:1,
Youxiang:"YouxiangTianjiayigeTest1",
}
zdxinxidaos.Tianjiayige(xinxi)
}
func TestXiugaiyigeXinxi(t *testing.T){
xinxi:=&appmodels.Xinxi{
Youxiang:"YouxiangXiugaiyigeTest1",
Bianma:"BianmaXiugaiyigeTest1",
Biaoji:"BiaojiXiugaiyigeTest1",
Id:1,
Mingcheng:"MingchengXiugaiyigeTest1",
}
zdxinxidaos.Xiugaiyige(xinxi)
}
func TestTianjiaduogeXinxi(t *testing.T){
xinxi2:=appmodels.Xinxi{
Id:1,
Mingcheng:"MingchengTianjiaduogeTest2",
Bianma:"BianmaTianjiaduogeTest2",
Biaoji:"BiaojiTianjiaduogeTest2",
Youxiang:"YouxiangTianjiaduogeTest2",
}
xinxi3:=appmodels.Xinxi{
Id:1,
Youxiang:"YouxiangTianjiaduogeTest3",
Biaoji:"BiaojiTianjiaduogeTest3",
Mingcheng:"MingchengTianjiaduogeTest3",
Bianma:"BianmaTianjiaduogeTest3",
}
xinxis:=[]appmodels.Xinxi{xinxi2,xinxi3}
zdxinxidaos.Tianjiaduoge(xinxis)
}
