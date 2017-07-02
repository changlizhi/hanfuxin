package baserun
import(
	"log"
	"bytes"
	"hanfuxin/baseinits"
	"hanfuxin/appinits"
	"strings"
)
func testsimports(bianma string,buffer *bytes.Buffer){
	buffer.WriteString(baseinits.Gen.Package)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Tests].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Gen.Import)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Hanfuxin].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiexian].Zhi)
	buffer.WriteString(baseinits.Mulus[appinits.Yingyongzi.Appmodels].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Testing].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Log].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Hanfuxin].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiexian].Zhi)
	baoming := strings.ToLower(bianma) + baseinits.Mulus[appinits.Yingyongzi.Daos].Zhi
	buffer.WriteString(baoming)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Shuangyinhao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
}

func testchaxunyige(bianma string,buffer *bytes.Buffer){
	buffer.WriteString(baseinits.Gen.Func)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	fangfaming := baseinits.Wenzis[appinits.Yingyongzi.Test].Bianma +
	baseinits.Wenzis[appinits.Yingyongzi.Chaxunyige].Bianma +
	bianma
	buffer.WriteString(fangfaming)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.T].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Konggefu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xinghao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Testing].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.T].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaozuo].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Maohao].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dengyuhao].Zhi)
	baoming := strings.ToLower(bianma) + baseinits.Mulus[appinits.Yingyongzi.Daos].Zhi
	buffer.WriteString(baoming)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Chaxunyige].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(baseinits.Jichus[appinits.Yingyongzi.Shuzi1].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)

	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Log].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dianhao].Zhi)
	buffer.WriteString(baseinits.Wenzis[appinits.Yingyongzi.Println].Bianma)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaozuo].Zhi)
	buffer.WriteString(strings.ToLower(bianma))
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Xiaokuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Dakuohaoyou].Zhi)
	buffer.WriteString(baseinits.Fuhaos[appinits.Yingyongzi.Huanhangfu].Zhi)
}


func Shengchengdaostests(){
	for _,biao := range baseinits.Biaos{
		buffer := bytes.Buffer{}
	  testsimports(biao.Bianma,&buffer)
	  testchaxunyige(biao.Bianma,&buffer)

	log.Println("buffer-------",buffer.String())
	}
}
