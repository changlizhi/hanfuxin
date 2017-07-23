package zdxinxijuesedaos
import (
"hanfuxin/appinits"
"hanfuxin/appmodels"
"hanfuxin/baseinits"
"hanfuxin/zf")
func Chaxunyige(id int) *appmodels.Xinxijuese {
xinxijuese := &appmodels.Xinxijuese{Id:id}
err:=appinits.Hanfuxinormer.Read(xinxijuese)
if err !=nil{
return nil
}
return xinxijuese
}
func Tianjiayige(xinxijuese *appmodels.Xinxijuese)string{
zf:=zf.Zf{}
_,err:=appinits.Hanfuxinormer.Insert(xinxijuese)
if err!=nil{
return baseinits.Tishis[zf.Tishi04(false)].Bianma
}
return baseinits.Tishis[zf.Tishi03(false)].Bianma
}
func Tianjiaduoge(xinxijueseshuzu []appmodels.Xinxijuese)string{
zf:=zf.Zf{}
_,err:=appinits.Hanfuxinormer.InsertMulti(len(xinxijueseshuzu),xinxijueseshuzu)
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
func Xiugaiyige(xinxijuese *appmodels.Xinxijuese)string{
zf:=zf.Zf{}
_,err:=appinits.Hanfuxinormer.Update(xinxijuese)
if err!=nil{
return baseinits.Tishis[zf.Tishi06(false)].Bianma
}
return baseinits.Tishis[zf.Tishi05(false)].Bianma
}
