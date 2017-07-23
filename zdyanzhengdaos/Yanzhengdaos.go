package zdyanzhengdaos
import (
"hanfuxin/appinits"
"hanfuxin/appmodels"
"hanfuxin/baseinits"
"hanfuxin/zf")
func Chaxunyige(id int) *appmodels.Yanzheng {
yanzheng := &appmodels.Yanzheng{Id:id}
err:=appinits.Hanfuxinormer.Read(yanzheng)
if err !=nil{
return nil
}
return yanzheng
}
func Tianjiayige(yanzheng *appmodels.Yanzheng)string{
zf:=zf.Zf{}
_,err:=appinits.Hanfuxinormer.Insert(yanzheng)
if err!=nil{
return baseinits.Tishis[zf.Tishi04(false)].Bianma
}
return baseinits.Tishis[zf.Tishi03(false)].Bianma
}
func Tianjiaduoge(yanzhengshuzu []appmodels.Yanzheng)string{
zf:=zf.Zf{}
_,err:=appinits.Hanfuxinormer.InsertMulti(len(yanzhengshuzu),yanzhengshuzu)
if err!=nil{
return baseinits.Tishis[zf.Tishi04(false)].Bianma
}
return baseinits.Tishis[zf.Tishi03(false)].Bianma
}
func Shanchuyige(id int)string{
zf:=zf.Zf{}
_,err:=appinits.Hanfuxinormer.Delete(Chaxunyige(id))
if err!=nil{
return baseinits.Tishis[zf.Tishi08(false)].Bianma
}
return baseinits.Tishis[zf.Tishi07(false)].Bianma
}
func Xiugaiyige(yanzheng *appmodels.Yanzheng)string{
zf:=zf.Zf{}
_,err:=appinits.Hanfuxinormer.Update(yanzheng)
if err!=nil{
return baseinits.Tishis[zf.Tishi06(false)].Bianma
}
return baseinits.Tishis[zf.Tishi05(false)].Bianma
}
