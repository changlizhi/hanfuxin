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
Biaoji:"BiaojiTianjiayigeTest1",
Bianma:"BianmaTianjiayigeTest1",
Mingcheng:"MingchengTianjiayigeTest1",
Leixing:"LeixingTianjiayigeTest1",
Zhi:"ZhiTianjiayigeTest1",
Id:1,
}
zdyanzhengdaos.Tianjiayige(yanzheng)
}
func TestXiugaiyigeYanzheng(t *testing.T){
yanzheng:=&appmodels.Yanzheng{
Mingcheng:"MingchengXiugaiyigeTest1",
Leixing:"LeixingXiugaiyigeTest1",
Zhi:"ZhiXiugaiyigeTest1",
Id:1,
Bianma:"BianmaXiugaiyigeTest1",
Biaoji:"BiaojiXiugaiyigeTest1",
}
zdyanzhengdaos.Xiugaiyige(yanzheng)
}
func TestTianjiaduoge(t *testing.T){
yanzheng2:=appmodels.Yanzheng{
Zhi:"ZhiTianjiaduogeTest2",
Id:1,
Bianma:"BianmaTianjiaduogeTest2",
Mingcheng:"MingchengTianjiaduogeTest2",
Leixing:"LeixingTianjiaduogeTest2",
Biaoji:"BiaojiTianjiaduogeTest2",
}
yanzheng3:=appmodels.Yanzheng{
Biaoji:"BiaojiTianjiaduogeTest3",
Leixing:"LeixingTianjiaduogeTest3",
Zhi:"ZhiTianjiaduogeTest3",
Id:1,
Bianma:"BianmaTianjiaduogeTest3",
Mingcheng:"MingchengTianjiaduogeTest3",
}
yanzhengs:=[]appmodels.Yanzheng{yanzheng2,yanzheng3}
zdyanzhengdaos.Tianjiaduoge(yanzhengs)
}
