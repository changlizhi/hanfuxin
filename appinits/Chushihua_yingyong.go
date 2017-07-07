package appinits

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/yingyongzimodel"
	"hanfuxin/zfz"
	"hanfuxin/zfzhi"
	"log"
	"strconv"
)

// 应用字全局变量，保存了从json里素有Bianma字段的字符化实体
var Yingyongzi yingyongzimodel.Yingyongzi

// 全局ormer，或许应该做成一个方法，否则后期要改成多线程的时候会引起问题
var Hanfuxinormer orm.Ormer

// 初始化方法，请勿改变先后顺序
func init() {
	chushihua_yingyongzi()
	chushihua_ormer()
}

// 初始化应用字，应用字的重点在于实现了所有字符的实体化表达，避免了魔数的存在
//这个实体会越来越大，但很方便拆分
func chushihua_yingyongzi() {
	// 从基础的初始化里读取json的Zifus,Zifus是所有Bianma的key和value一样的map，
	// 这个map保证了通过yingyognzi的实体进行取值时属性和值是一样的
	// 先把map转成json
	jsonobj, err := json.Marshal(baseinits.Zifus)
	if err != nil {
		log.Println(err)
		log.Println("解析Zifus[]到json格式出错")
	}
	// 再把json 转成实体
	json.Unmarshal(jsonobj, &Yingyongzi)
}
func chushihua_ormer() {
	zf := zfz.Zf{}
	// 设置orm是否为debug模式
	mh := zfzhi.Maohaozhi()
	qa := zfzhi.Quanazhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	xx := zfzhi.Xiexianzhi()
	orm.Debug, _ = strconv.ParseBool(baseinits.Shezhis[Yingyongzi.Ormdebug].Zhi)
	// 注册所有的实体，这些实体全部都是在baserun里生成的，请使用自动生成再在这里添加，
	// 后期这个初始化将拆分并自动生成
	orm.RegisterModel(
		new(appmodels.Juese),
		new(appmodels.Juesequanxian),
		new(appmodels.Quanxian),
		new(appmodels.Xinxi),
		new(appmodels.Xinxijuese),
		new(appmodels.Yanzheng),
		new(appmodels.Yanzhengleixing),
	)
	// orm注册数据库
	orm.RegisterDataBase(
		// 数据库名称，conf/changliang.json里Wenzi需要有一个default数据库
		zf.Default(true),
		// 数据库驱动，在conf/yingyong.jsond的Shezhi里
		baseinits.Shezhis[Yingyongzi.Shujukuqudong].Zhi,
		// 数据库链接，yonghu:mima@tcp(ip:duankou)/shujukuming
		// yonghu:mima
		baseinits.Shezhis[Yingyongzi.Shujukuyonghu].Zhi+
			mh+
			baseinits.Shezhis[Yingyongzi.Shujukumima].Zhi+
			//@tcp
			qa+
			zf.Tcp(true)+
			//(ip:duankou)
			xkhz +
			baseinits.Shezhis[Yingyongzi.Shujukuip].Zhi+
			mh+
			baseinits.Shezhis[Yingyongzi.Shujukuduankou].Zhi+
			xkhy+
			// /shujukuming
			xx+
			baseinits.Shezhis[Yingyongzi.Shujuku].Zhi)
	Hanfuxinormer = orm.NewOrm()
	Hanfuxinormer.Using(baseinits.Shezhis[Yingyongzi.Default].Zhi)
}
