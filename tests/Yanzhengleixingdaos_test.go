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
Biaoji:"BiaojiTianjiayigeTest1",
Id:1,
Bianma:"BianmaTianjiayigeTest1",
Mingcheng:"MingchengTianjiayigeTest1",
}
zdyanzhengleixingdaos.Tianjiayige(yanzhengleixing)
}
func TestXiugaiyigeYanzhengleixing(t *testing.T){
yanzhengleixing:=&appmodels.Yanzhengleixing{
Biaoji:"BiaojiXiugaiyigeTest1",
Id:1,
Bianma:"BianmaXiugaiyigeTest1",
Mingcheng:"MingchengXiugaiyigeTest1",
}
zdyanzhengleixingdaos.Xiugaiyige(yanzhengleixing)
}
func TestTianjiaduoge(t *testing.T){
yanzhengleixing2:=appmodels.Yanzhengleixing{
Mingcheng:"MingchengTianjiaduogeTest2",
Id:1,
Bianma:"BianmaTianjiaduogeTest2",
Biaoji:"BiaojiTianjiaduogeTest2",
}
yanzhengleixing3:=appmodels.Yanzhengleixing{
Biaoji:"BiaojiTianjiaduogeTest3",
Bianma:"BianmaTianjiaduogeTest3",
Mingcheng:"MingchengTianjiaduogeTest3",
Id:1,
}
yanzhengleixings:=[]appmodels.Yanzhengleixing{yanzhengleixing2,yanzhengleixing3}
zdyanzhengleixingdaos.Tianjiaduoge(yanzhengleixings)
}
