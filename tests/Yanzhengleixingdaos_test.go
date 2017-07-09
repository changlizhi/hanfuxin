package tests
import(
"hanfuxin/appmodels"
"testing"
"log"
"hanfuxin/zdyanzhengleixingdaos"
)
func TestChaxunyigeYanzhengleixing(t *testing.T){
yanzhengleixing:=zdyanzhengleixingdaos.Chaxunyige(1)
log.Println(yanzhengleixing)
}
func TestShanchuyigeYanzhengleixing(t *testing.T){
zdyanzhengleixingdaos.Shanchuyige(1)
}
func TestTianjiayigeYanzhengleixing(t *testing.T){
yanzhengleixing:=&appmodels.Yanzhengleixing{
Id:1,
Mingcheng:"MingchengTianjiayigeTest1",
Biaoji:"BiaojiTianjiayigeTest1",
Bianma:"BianmaTianjiayigeTest1",
}
zdyanzhengleixingdaos.Tianjiayige(yanzhengleixing)
}
func TestXiugaiyigeYanzhengleixing(t *testing.T){
yanzhengleixing:=&appmodels.Yanzhengleixing{
Id:1,
Biaoji:"BiaojiXiugaiyigeTest1",
Bianma:"BianmaXiugaiyigeTest1",
Mingcheng:"MingchengXiugaiyigeTest1",
}
zdyanzhengleixingdaos.Xiugaiyige(yanzhengleixing)
}
func TestTianjiaduogeYanzhengleixing(t *testing.T){
yanzhengleixing2:=appmodels.Yanzhengleixing{
Mingcheng:"MingchengTianjiaduogeTest2",
Id:1,
Biaoji:"BiaojiTianjiaduogeTest2",
Bianma:"BianmaTianjiaduogeTest2",
}
yanzhengleixing3:=appmodels.Yanzhengleixing{
Bianma:"BianmaTianjiaduogeTest3",
Biaoji:"BiaojiTianjiaduogeTest3",
Id:1,
Mingcheng:"MingchengTianjiaduogeTest3",
}
yanzhengleixings:=[]appmodels.Yanzhengleixing{yanzhengleixing2,yanzhengleixing3}
zdyanzhengleixingdaos.Tianjiaduoge(yanzhengleixings)
}
