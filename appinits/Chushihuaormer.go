package appinits

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"hanfuxin/appmodels"
	"changliang/zf"
	"changliang/zfzhi"
	"strconv"
)

// 初始化方法，请勿改变先后顺序
func init() {
	ormerdebug()
	ormermoxing()
	ormershujuku()
}

func ormerdebug() {
	// 设置orm是否为debug模式
	Setormerdebug(Chushihuas[zf.Zfs.Ormdebug(false)].Zhi)
}
func Setormerdebug(boolstr string) {
	orm.Debug, _ = strconv.ParseBool(boolstr)
}

func ormermoxing() {
	// 注册所有的实体，这些实体全部都是在baserun里生成的，请使用自动生成再在这里添加，
	// 后期这个初始化将拆分并自动生成
	orm.RegisterModel(
		new(appmodels.Juese),
		new(appmodels.Juesequanxian),
		new(appmodels.Quanxian),
		new(appmodels.Xinxi),
		new(appmodels.Xinxijuese),
		new(appmodels.Yanzheng),
		new(appmodels.Yanzhengleixing,
		),
	)
}
func ormershujuku() {
	// orm注册数据库
	// root:root@tcp(ip:duankou)/mingcheng
	url := Shujukus[zf.Zfs.Yonghu(false)].Zhi +
		zfzhi.Zhi.Mh() +
		Shujukus[zf.Zfs.Mima(false)].Zhi +
		zfzhi.Zhi.Qa() +
		zf.Zfs.Tcp(true) +
		zfzhi.Zhi.Xkhz() +
		Shujukus[zf.Zfs.Ip(false)].Zhi +
		zfzhi.Zhi.Mh() +
		Shujukus[zf.Zfs.Duankou(false)].Zhi +
		zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Xx() +
		Shujukus[zf.Zfs.Mingcheng(false)].Zhi
	orm.RegisterDataBase(zf.Zfs.Default(true), Shujukus[zf.Zfs.Qudong(false)].Zhi, url)
}
func Defaultormer() orm.Ormer {
	return Ormerbyname(zf.Zfs.Default(true))
}
func Ormerbyname(bieming string) orm.Ormer {
	ret := orm.NewOrm()
	ret.Using(bieming)
	return ret
}