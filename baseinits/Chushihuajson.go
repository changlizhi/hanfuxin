package baseinits

import (
	"hanfuxin/basemodels"
	"hanfuxin/baseutils"
)

// 这些常量也应该以方法的方式提供么，没必要吧，因为全都是从json里直接获取的

// Shezhis---用来初始化系统时进行设置一些值的参数提供
var Shezhis = make(map[string]basemodels.Tongyong)

// Mulus---项目的目录说明，除了自动生成的那些目录只提供了后缀：比如daos和services，其他的都原样保持在json中
var Mulus = make(map[string]basemodels.Tongyong)

// Biaos---这个是项目所有的表，只包含名称和注释
var Biaos = make(map[string]basemodels.Tongyong)

// Lies----这个是所有的字段名，有一个Biaoming字段用来保存所有用到此字段的所有表
// 在自动生成表时通过这个关系进行生成，一旦匹配到表名则说名该表应该有此字段
// 类型主要包括----int:INT,float32:FLOAT,string:VARCHAR(50),time:TIMESTAMP
// 后期将通过此关系进行数据库脚本自动生成
var Lies = make(map[string]basemodels.Biaojiegou)

// 基础对应：如有效:1，无效:0
var Jichus = make(map[string]basemodels.Tongyong)

// 网络规范常用词
var Wangluos = make(map[string]basemodels.Tongyong)

// 所有的符号
var Fuhaos = make(map[string]basemodels.Tongyong)

// 所有的文字串
var Wenzis = make(map[string]basemodels.Tongyong)

// 所有的Bianma对应
var Zifus = make(map[string]string)

// 包好最基础的字符，如关键字和生成第一个model时需要的字符常量
// 这里后期要添加到读取json那去么，需要考量
var Gen basemodels.Gen

// 中文
var Zhongwens = make(map[string]basemodels.Tongyong)

// 英文
var Yingwens = make(map[string]basemodels.Tongyong)

func init() {
	chushihua_json()
}
func chushihua_json() {
	// 读取changliang.json
	changliang := baseutils.Changliangjson()
	// 读取yingyong.json
	yingyong := baseutils.Yingyongjson()
	// 直接赋值给全局变量，为了所有的数据都从这个包提供，所以不把changliang做成全局
	Gen = changliang.Gen

	// 遍历两个json的各个定义并初始化全局变量
	for _, zhongwen := range changliang.Zhongwen {
		Zifus[zhongwen.Guilei] = zhongwen.Guilei // 冗余，为了简化json结构
		Zhongwens[zhongwen.Guilei+zhongwen.Bianma] = zhongwen
		Zifus[zhongwen.Guilei+zhongwen.Bianma] = zhongwen.Bianma
	}
	for _, yingwen := range changliang.Yingwen {
		Zifus[yingwen.Guilei] = yingwen.Guilei // 冗余，为了简化json结构
		Yingwens[yingwen.Guilei+yingwen.Bianma] = yingwen
		Zifus[yingwen.Guilei+yingwen.Bianma] = yingwen.Bianma
	}
	for _, fuhao := range changliang.Fuhao {
		Zifus[fuhao.Guilei] = fuhao.Guilei // 冗余，为了简化json结构
		Zifus[fuhao.Bianma] = fuhao.Bianma
		Fuhaos[fuhao.Bianma] = fuhao
	}
	for _, jichu := range changliang.Jichu {
		Zifus[jichu.Guilei] = jichu.Guilei
		Zifus[jichu.Bianma] = jichu.Bianma
		Jichus[jichu.Bianma] = jichu
	}
	for _, wangluo := range changliang.Wangluo {
		Zifus[wangluo.Guilei] = wangluo.Guilei
		Zifus[wangluo.Bianma] = wangluo.Bianma
		Wangluos[wangluo.Bianma] = wangluo
	}
	for _, wenzi := range changliang.Wenzi {
		Zifus[wenzi.Guilei] = wenzi.Guilei
		Zifus[wenzi.Bianma] = wenzi.Bianma
		Wenzis[wenzi.Bianma] = wenzi
	}
	for _, biao := range yingyong.Biao {
		Zifus[biao.Guilei] = biao.Guilei
		Zifus[biao.Bianma] = biao.Bianma
		Biaos[biao.Bianma] = biao
	}
	for _, lie := range yingyong.Lie {
		Zifus[lie.Guilei] = lie.Guilei
		Zifus[lie.Bianma] = lie.Bianma
		Lies[lie.Bianma] = lie
	}
	for _, shezhi := range yingyong.Shezhi {
		Zifus[shezhi.Guilei] = shezhi.Guilei
		Zifus[shezhi.Bianma] = shezhi.Bianma
		Shezhis[shezhi.Bianma] = shezhi
	}
	for _, mulu := range yingyong.Mulu {
		Zifus[mulu.Guilei] = mulu.Guilei
		Zifus[mulu.Bianma] = mulu.Bianma
		Mulus[mulu.Bianma] = mulu
	}
}
