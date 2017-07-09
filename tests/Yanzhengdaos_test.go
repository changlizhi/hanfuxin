package tests
import(
"hanfuxin/appmodels"
"testing"
"log"
"hanfuxin/zdyanzhengdaos"
)
func TestChaxunyigeYanzheng(t *testing.T){
yanzheng:=zdyanzhengdaos.Chaxunyige(1)
log.Println(yanzheng)
}
func TestShanchuyigeYanzheng(t *testing.T){
zdyanzhengdaos.Shanchuyige(1)
}
func TestTianjiayigeYanzheng(t *testing.T){
yanzheng:=&appmodels.Yanzheng{
Mingcheng:"MingchengTianjiayigeTest1",
Zhi:"ZhiTianjiayigeTest1",
Id:1,
Leixing:"LeixingTianjiayigeTest1",
Biaoji:"BiaojiTianjiayigeTest1",
Bianma:"BianmaTianjiayigeTest1",
}
zdyanzhengdaos.Tianjiayige(yanzheng)
}
func TestXiugaiyigeYanzheng(t *testing.T){
yanzheng:=&appmodels.Yanzheng{
Zhi:"ZhiXiugaiyigeTest1",
Id:1,
Leixing:"LeixingXiugaiyigeTest1",
Bianma:"BianmaXiugaiyigeTest1",
Biaoji:"BiaojiXiugaiyigeTest1",
Mingcheng:"MingchengXiugaiyigeTest1",
}
zdyanzhengdaos.Xiugaiyige(yanzheng)
}
func TestTianjiaduogeYanzheng(t *testing.T){
yanzheng2:=appmodels.Yanzheng{
Bianma:"BianmaTianjiaduogeTest2",
Leixing:"LeixingTianjiaduogeTest2",
Id:1,
Mingcheng:"MingchengTianjiaduogeTest2",
Biaoji:"BiaojiTianjiaduogeTest2",
Zhi:"ZhiTianjiaduogeTest2",
}
yanzheng3:=appmodels.Yanzheng{
Id:1,
Bianma:"BianmaTianjiaduogeTest3",
Zhi:"ZhiTianjiaduogeTest3",
Mingcheng:"MingchengTianjiaduogeTest3",
Leixing:"LeixingTianjiaduogeTest3",
Biaoji:"BiaojiTianjiaduogeTest3",
}
yanzhengs:=[]appmodels.Yanzheng{yanzheng2,yanzheng3}
zdyanzhengdaos.Tianjiaduoge(yanzhengs)
}
