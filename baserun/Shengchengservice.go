package baserun
import(
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"bytes"
	"log"
	"strings"
)
func serviceimports(bianma string,buffer *bytes.Buffer){
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	syh := zfzhi.Shuangyinhaozhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	xx := zfzhi.Xiexianzhi()
	hhf := zfzhi.Huanhangfuzhi()

	buffer.WriteString(zf.Import(true))
	buffer.WriteString(xkhz)
	buffer.WriteString(hhf)
	//"hanfuxin/allerrors" \n
	buffer.WriteString(syh)
	buffer.WriteString(zf.Hanfuxin(true))
	buffer.WriteString(xx)
	buffer.WriteString(zf.Allerrors(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)

	// "bytes"
	buffer.WriteString(syh)
	buffer.WriteString(zf.Bytes(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)
	// "log"
	buffer.WriteString(syh)
	buffer.WriteString(zf.Log(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)
	// "time"
	buffer.WriteString(syh)
	buffer.WriteString(zf.Time(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)
	// "hanfuxin/appmodels"
	buffer.WriteString(syh)
	buffer.WriteString(zf.Hanfuxin(true))
	buffer.WriteString(xx)
	buffer.WriteString(zf.Appmodels(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)
	// "hanfuxin/apputils"
	buffer.WriteString(syh)
	buffer.WriteString(zf.Hanfuxin(true))
	buffer.WriteString(xx)
	buffer.WriteString(zf.Apputils(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)
	// "hanfuxin/baserun"
	buffer.WriteString(syh)
	buffer.WriteString(zf.Hanfuxin(true))
	buffer.WriteString(xx)
	buffer.WriteString(zf.Baserun(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)
	// "hanfuxin/zf"
	buffer.WriteString(syh)
	buffer.WriteString(zf.Hanfuxin(true))
	buffer.WriteString(xx)
	buffer.WriteString(zf.Zf(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)
	// "hanfuxin/zfzhi"
	buffer.WriteString(syh)
	buffer.WriteString(zf.Hanfuxin(true))
	buffer.WriteString(xx)
	buffer.WriteString(zf.Zfzhi(true))
	buffer.WriteString(syh)
	buffer.WriteString(hhf)
	// "hanfuxin/zdxxxdaos"
	buffer.WriteString(syh)
	buffer.WriteString(zf.Hanfuxin(true))
	buffer.WriteString(xx)
	daobao := zf.Zd(true) + strings.ToLower(bianma) + zf.Daos(true)
	buffer.WriteString(daobao)
	buffer.WriteString(syh)
	buffer.WriteString(hhf)
	
	buffer.WriteString(xkhy)
	buffer.WriteString(hhf)
}
func yanzhengchangdu(bianma string,buffer *bytes.Buffer){
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xh := zfzhi.Xinghaozhi()
	hhf := zfzhi.Huanhangfuzhi()
	mh :=zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	dh := zfzhi.Dianhaozhi()

	bmx := strings.ToLower(bianma)

	kgf := zfzhi.Konggefuzhi()
	buffer.WriteString(zf.Func(true))
	buffer.WriteString(kgf)
	buffer.WriteString(zf.Yanzhengziduanchangdu(true))
	//(bmx *appmodels.bm) error { \n	
	buffer.WriteString(xkhz)
	buffer.WriteString(bmx)
	buffer.WriteString(kgf)
	buffer.WriteString(xh)
	buffer.WriteString(zf.Appmodels(true))
	buffer.WriteString(bianma)
	buffer.WriteString(xkhy)
	buffer.WriteString(zf.Error(true))
	buffer.WriteString(dkhz)
	buffer.WriteString(hhf)

	// zf := zf.Zf{}
	buffer.WriteString(zf.Zf(true))
	buffer.WriteString(mh)
	buffer.WriteString(dyh)
	buffer.WriteString(zf.Zf(true))
	buffer.WriteString(dh)
	buffer.WriteString(zf.Zf(false))
	buffer.WriteString(dkhz)
	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)



	buffer.WriteString(hhf)
	buffer.WriteString(dkhy)
	
}
func Shengchengservice(){
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()
	for bk,bv := range Huoqubiaos(){
		buffer := bytes.Buffer{}
		bm := zf.Zd(true) + bv + zf.Services(true)
		//package xxxservices \n
		buffer.WriteString(zf.Package(true))
		buffer.WriteString(kgf)
		buffer.WriteString(bm)
		buffer.WriteString(hhf)

		serviceimports(bk,&buffer)
		yanzhengchangdu(bk,&buffer)
		log.Println("buffer--------",buffer.String())
	}
}
