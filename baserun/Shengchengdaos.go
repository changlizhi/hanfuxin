package baserun

import (
	"bytes"
	"hanfuxin/basemodels"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"io/ioutil"
	"os"
	"strings"
)

func imports(bianma string, buffer *bytes.Buffer) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	xx := zfzhi.Xiexianzhi()
	hfx := zf.Hanfuxin(true)
	syh := zfzhi.Shuangyinhaozhi()
	// \n import (\n
	buffer.WriteString(hhf)
	buffer.WriteString(zf.Import(true))
	buffer.WriteString(zfzhi.Konggefuzhi())
	buffer.WriteString(zfzhi.Xiaokuohaozuozhi())
	buffer.WriteString(hhf)

	// "hanfuxin/appinits" \n
	buffer.WriteString(syh)               // "
	buffer.WriteString(hfx)               // hanfuxin
	buffer.WriteString(xx)                // /
	buffer.WriteString(zf.Appinits(true)) // appinits
	buffer.WriteString(syh)               // "
	buffer.WriteString(hhf)

	// "hanfuxin/appmodels" \n
	buffer.WriteString(syh)                // "
	buffer.WriteString(hfx)                // hanfuxin
	buffer.WriteString(xx)                 // /
	buffer.WriteString(zf.Appmodels(true)) // appmodels
	buffer.WriteString(syh)                // "
	buffer.WriteString(hhf)

	// ) \n
	buffer.WriteString(zfzhi.Xiaokuohaoyouzhi()) // )
	buffer.WriteString(hhf)                      // \n
}
func chaxunyige(bianma string, buffer *bytes.Buffer) {
	// func Chaxunyige (id int)
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	idx := zf.Id(true)
	idd := zf.Id(false)
	appmodels := zf.Appmodels(true)

	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kongge := zfzhi.Konggefuzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	mh := zfzhi.Maohaozhi()
	dyh := zfzhi.Dengyuhaozhi()
	qh := zfzhi.Qiehaozhi()
	xh := zfzhi.Xinghaozhi()

	bmx := strings.ToLower(bianma)

	buffer.WriteString(zf.Func(true))
	buffer.WriteString(kongge)               // kongge
	buffer.WriteString(zf.Chaxunyige(false)) //Chaxunyige
	buffer.WriteString(xkhz)                 // (
	buffer.WriteString(idx)                  // id
	buffer.WriteString(kongge)               // kongge
	buffer.WriteString(zf.Int(true))         // int
	buffer.WriteString(xkhy)                 // )

	// *appmodels.Juese{ \n
	buffer.WriteString(kongge)    // kongge
	buffer.WriteString(xh)        // *
	buffer.WriteString(appmodels) // appmodels
	buffer.WriteString(dh)        // .
	buffer.WriteString(bianma)    // Juese
	buffer.WriteString(kongge)    // kongge
	buffer.WriteString(dkhz)      // {
	buffer.WriteString(hhf)       // \n

	// juese := &appmodels.Juese
	buffer.WriteString(bmx)       // juese
	buffer.WriteString(kongge)    // kongge
	buffer.WriteString(mh)        // :
	buffer.WriteString(dyh)       // =
	buffer.WriteString(kongge)    // kongge
	buffer.WriteString(qh)        // &
	buffer.WriteString(appmodels) // appmodels
	buffer.WriteString(dh)        // .
	buffer.WriteString(bianma)    //Juese

	// {Id:id}
	buffer.WriteString(dkhz) // {
	buffer.WriteString(idd)  // Id
	buffer.WriteString(mh)   // :
	buffer.WriteString(idx)  // id
	buffer.WriteString(dkhy) // }

	// \n appinits.Hanfuxinormer.Read(juese)
	buffer.WriteString(hhf)
	buffer.WriteString(zf.Appinits(true))       // appinits
	buffer.WriteString(dh)                      // .
	buffer.WriteString(zf.Hanfuxinormer(false)) // Hanfuxinormer
	buffer.WriteString(dh)                      // .
	buffer.WriteString(zf.Read(false))          // Read
	buffer.WriteString(xkhz)                    // (
	buffer.WriteString(bmx)                     // juese
	buffer.WriteString(xkhy)                    // )

	// \n return juese
	buffer.WriteString(hhf)
	buffer.WriteString(zf.Return(true)) // return
	buffer.WriteString(kongge)          // kongge
	buffer.WriteString(bmx)             // juese

	//\n } \n
	buffer.WriteString(hhf)
	buffer.WriteString(dkhy) // }
	buffer.WriteString(hhf)
}
func tianjiayige(bianma string, buffer *bytes.Buffer) {
	// func Tianjiayige
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()

	bmx := strings.ToLower(bianma)

	buffer.WriteString(zf.Func(true))         //func
	buffer.WriteString(kgf)                   // kongge
	buffer.WriteString(zf.Tianjiayige(false)) // Tianjiayige

	//(juese *appmodels.Juese){ \n
	buffer.WriteString(xkhz)               //(
	buffer.WriteString(bmx)                //juese
	buffer.WriteString(kgf)                // kongge
	buffer.WriteString(zfzhi.Xinghaozhi()) // *
	buffer.WriteString(zf.Appmodels(true)) //appmodels
	buffer.WriteString(dh)                 // .
	buffer.WriteString(bianma)             // Juese
	buffer.WriteString(xkhy)               // (
	buffer.WriteString(dkhz)               //{
	buffer.WriteString(hhf)

	// appinits.Hanfuxinormer.Insert
	buffer.WriteString(zf.Appinits(true))       // appinits
	buffer.WriteString(dh)                      // .
	buffer.WriteString(zf.Hanfuxinormer(false)) //Hanfuxinormer
	buffer.WriteString(dh)                      // .
	buffer.WriteString(zf.Insert(false))        //Insert

	// (juese)
	buffer.WriteString(xkhz) // (
	buffer.WriteString(bmx)  // jese
	buffer.WriteString(xkhy) //)

	// \n } \n
	buffer.WriteString(hhf)
	buffer.WriteString(dkhy) // }
	buffer.WriteString(hhf)
}
func tianjiaduoge(bianma string, buffer *bytes.Buffer) {
	//func Tianjiaduoge(
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	douhao := zfzhi.Douhaozhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	kgf := zfzhi.Konggefuzhi()

	buffer.WriteString(zf.Func(true))          //func
	buffer.WriteString(kgf)                    //kongge
	buffer.WriteString(zf.Tianjiaduoge(false)) //Tianjiayige
	buffer.WriteString(xkhz)                   // (

	// jueseshuzu
	canshu := strings.ToLower(bianma) + zf.Shuzu(true)
	buffer.WriteString(canshu)

	// []appmodels.Juese){ \n
	buffer.WriteString(kgf)                //k ongge
	buffer.WriteString(zkhz)               // [
	buffer.WriteString(zkhy)               // ]
	buffer.WriteString(zf.Appmodels(true)) //appmodels
	buffer.WriteString(dh)                 // .
	buffer.WriteString(bianma)             // Juese
	buffer.WriteString(xkhy)               // )
	buffer.WriteString(dkhz)               // {
	buffer.WriteString(hhf)

	// appinits.Hanfuxinormer.InsertMulti
	buffer.WriteString(zf.Appinits(true))       // appinits
	buffer.WriteString(dh)                      // .
	buffer.WriteString(zf.Hanfuxinormer(false)) // Hanfuxinormer
	buffer.WriteString(dh)                      // .
	buffer.WriteString(zf.InsertMulti(false))   //InsertMulti

	//(len(jueseshuzu),jueseshuzu)
	buffer.WriteString(xkhz)         //(
	buffer.WriteString(zf.Len(true)) //len
	buffer.WriteString(xkhz)         //(
	buffer.WriteString(canshu)       //jueseshuzu
	buffer.WriteString(xkhy)         //)
	buffer.WriteString(douhao)       //,
	buffer.WriteString(canshu)       //jueseshuzu
	buffer.WriteString(xkhy)         //)

	//\n } \n
	buffer.WriteString(hhf)
	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)
}
func shanchuyige(bianma string, buffer *bytes.Buffer) {
	// func Shanchuyige
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	id := zf.Id(true)

	buffer.WriteString(zf.Func(true))         // func
	buffer.WriteString(kgf)                   // kongge
	buffer.WriteString(zf.Shanchuyige(false)) //Shanchuyige

	//(ic int){ \n
	buffer.WriteString(xkhz) //(
	buffer.WriteString(id)
	buffer.WriteString(kgf)
	buffer.WriteString(zf.Int(true)) // int
	buffer.WriteString(xkhy)         //)
	buffer.WriteString(dkhz)         //{
	buffer.WriteString(hhf)

	// appinits.Hanfuxinormer.Delete
	buffer.WriteString(zf.Appinits(true))       // appinits
	buffer.WriteString(dh)                      // .
	buffer.WriteString(zf.Hanfuxinormer(false)) // Hanfuxinormer
	buffer.WriteString(dh)                      // .
	buffer.WriteString(zf.Delete(false))        // Delete

	//(Chaxunyige(id))
	buffer.WriteString(xkhz)
	buffer.WriteString(zf.Chaxunyige(false)) // Chaxunyige
	buffer.WriteString(xkhz)
	buffer.WriteString(id) //id
	buffer.WriteString(xkhy)
	buffer.WriteString(xkhy)

	// \n } \n
	buffer.WriteString(hhf)
	buffer.WriteString(dkhy)
	buffer.WriteString(hhf)
}
func xiugaiyige(bianma string, buffer *bytes.Buffer) {
	// func Xiugaiyige
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	hhf := zfzhi.Huanhangfuzhi()
	dh := zfzhi.Dianhaozhi()
	kgf := zfzhi.Konggefuzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	xh := zfzhi.Xinghaozhi()

	bmx := strings.ToLower(bianma)

	buffer.WriteString(zf.Func(true))        // func
	buffer.WriteString(kgf)                  // kongge
	buffer.WriteString(zf.Xiugaiyige(false)) //Xiugaiyige

	//(juese *appmodels.Juese){ \n
	buffer.WriteString(xkhz)               //(
	buffer.WriteString(bmx)                //juese
	buffer.WriteString(kgf)                // kongge
	buffer.WriteString(xh)                 // *
	buffer.WriteString(zf.Appmodels(true)) //appmodels
	buffer.WriteString(dh)                 // .
	buffer.WriteString(bianma)             //Juese
	buffer.WriteString(xkhy)               // )
	buffer.WriteString(dkhz)               //{
	buffer.WriteString(hhf)

	//appinits.Hanfuxinormer.Update
	buffer.WriteString(zf.Appinits(true))       //appinits
	buffer.WriteString(dh)                      //.
	buffer.WriteString(zf.Hanfuxinormer(false)) //hanfuxinormer
	buffer.WriteString(dh)                      // .
	buffer.WriteString(zf.Update(false))        //Update

	//(juese)
	buffer.WriteString(xkhz) //(
	buffer.WriteString(bmx)  //juese
	buffer.WriteString(xkhy) //)

	// \n } \n
	buffer.WriteString(hhf)
	buffer.WriteString(dkhy) //}
	buffer.WriteString(hhf)

}
func Shengchengdaos() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	dh := zfzhi.Dianhaozhi()
	xx := zfzhi.Xiexianzhi()

	daos := zf.Daos(true)
	for biao, _ := range Huoqubiaos() {
		buffer := bytes.Buffer{}
		buffer.WriteString(zf.Package(true))    //package
		buffer.WriteString(zfzhi.Konggefuzhi()) // kongge
		lujing := zf.Zd(true) + strings.ToLower(biao) + daos
		buffer.WriteString(lujing) // zdjuesedaos

		imports(biao, &buffer)      //import all
		chaxunyige(biao, &buffer)   // Chaxunyige
		tianjiayige(biao, &buffer)  // Tianjiayige
		tianjiaduoge(biao, &buffer) // Tianjiaduoge
		shanchuyige(biao, &buffer)  // Shanchuyige
		xiugaiyige(biao, &buffer)   // Xiugaiyige

		dir := basemodels.Getapppath() + xx + lujing + xx
		path := dir + biao + daos + dh + zf.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
