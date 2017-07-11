package appinits

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"strconv"
)

// 全局ormer，或许应该做成一个方法，否则后期要改成多线程的时候会引起问题
var Hanfuxinormer orm.Ormer

// 初始化方法，请勿改变先后顺序
func init() {
	chushihua_ormer()
}

func chushihua_ormer() {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	// 设置orm是否为debug模式
	mh := zfzhi.Maohaozhi()
	qa := zfzhi.Quanazhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	xx := zfzhi.Xiexianzhi()
	orm.Debug, _ = strconv.ParseBool(baseinits.Chushihuas[zf.Ormdebug(false)].Zhi)
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
	// root:root@tcp(ip:duankou)/mingcheng
	url := baseinits.Shujukus[zf.Yonghu(false)].Zhi +
		mh +
		baseinits.Shujukus[zf.Mima(false)].Zhi +
		qa +
		zf.Tcp(true) +
		xkhz +
		baseinits.Shujukus[zf.Ip(false)].Zhi +
		mh +
		baseinits.Shujukus[zf.Duankou(false)].Zhi +
		xkhy +
		xx +
		baseinits.Shujukus[zf.Mingcheng(false)].Zhi
	orm.RegisterDataBase(zf.Default(true), baseinits.Shujukus[zf.Qudong(false)].Zhi, url)

	Hanfuxinormer = orm.NewOrm()
	Hanfuxinormer.Using(zf.Default(true))
}
