package appinits

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"hanfuxin/yingyongzimodel"
	"log"
	"strconv"
)

var Yingyongzi yingyongzimodel.Yingyongzi
var Hanfuxinormer orm.Ormer

func init() {
	chushihua_yingyongzi()
	chushihua_ormer()
}
func chushihua_yingyongzi() {
	jsonobj, err := json.Marshal(baseinits.Zifus)
	if err != nil {
		log.Println(err)
		log.Println("解析Zifus[]到json格式出错")
	}
	json.Unmarshal(jsonobj, &Yingyongzi)
}
func chushihua_ormer() {
	orm.Debug, _ = strconv.ParseBool(baseinits.Shezhis[Yingyongzi.Ormdebug].Zhi)
	orm.RegisterModel(
		new(appmodels.Juese),
		new(appmodels.Juesequanxian),
		new(appmodels.Quanxian),
		new(appmodels.Xinxi),
		new(appmodels.Xinxijuese),
		new(appmodels.Yanzheng),
		new(appmodels.Yanzhengleixing),
	)
	orm.RegisterDataBase(
		baseinits.Wenzis[Yingyongzi.Default].Zhi,
		baseinits.Shezhis[Yingyongzi.Shujukuqudong].Zhi,
		baseinits.Shezhis[Yingyongzi.Shujukuyonghu].Zhi+
			baseinits.Fuhaos[Yingyongzi.Maohao].Zhi+
			baseinits.Shezhis[Yingyongzi.Shujukumima].Zhi+
			baseinits.Fuhaos[Yingyongzi.Quana].Zhi+
			baseinits.Wangluos[Yingyongzi.Xieyitcp].Zhi+
			baseinits.Fuhaos[Yingyongzi.Xiaokuohaozuo].Zhi+
			baseinits.Shezhis[Yingyongzi.Shujukuip].Zhi+
			baseinits.Fuhaos[Yingyongzi.Maohao].Zhi+
			baseinits.Shezhis[Yingyongzi.Shujukuduankou].Zhi+
			baseinits.Fuhaos[Yingyongzi.Xiaokuohaoyou].Zhi+
			baseinits.Fuhaos[Yingyongzi.Xiexian].Zhi+
			baseinits.Shezhis[Yingyongzi.Shujuku].Zhi)
	log.Println(
		baseinits.Shezhis[Yingyongzi.Shujukuyonghu].Zhi+
			baseinits.Fuhaos[Yingyongzi.Maohao].Zhi+
			baseinits.Shezhis[Yingyongzi.Shujukumima].Zhi+
			baseinits.Fuhaos[Yingyongzi.Quana].Zhi+
			baseinits.Wangluos[Yingyongzi.Xieyitcp].Zhi+
			baseinits.Fuhaos[Yingyongzi.Xiaokuohaozuo].Zhi+
			baseinits.Shezhis[Yingyongzi.Shujukuip].Zhi+
			baseinits.Fuhaos[Yingyongzi.Maohao].Zhi+
			baseinits.Shezhis[Yingyongzi.Shujukuduankou].Zhi+
			baseinits.Fuhaos[Yingyongzi.Xiaokuohaoyou].Zhi+
			baseinits.Fuhaos[Yingyongzi.Xiexian].Zhi+
			baseinits.Shezhis[Yingyongzi.Shujuku].Zhi,
		"---------------数据库链接-------"+baseinits.Wenzis[Yingyongzi.Default].Zhi)
	Hanfuxinormer = orm.NewOrm()
	Hanfuxinormer.Using(baseinits.Shezhis[Yingyongzi.Default].Zhi)
}
