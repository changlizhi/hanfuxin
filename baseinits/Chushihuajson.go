package baseinits

import (
	"hanfuxin/basemodels"
	"hanfuxin/baseutils"
)

// 这些常量也应该以方法的方式提供么，没必要吧，因为全都是从json里直接获取的
// Shezhis---用来初始化系统时进行设置一些值的参数提供
var Shezhis = make(map[string]basemodels.Tongyong)

// 所有的Bianma对应
var Zifus = make(map[string]string)

// 包好最基础的字符，如关键字和生成第一个model时需要的字符常量
// 这里后期要添加到读取json那去么，需要考量

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
	for _, shezhi := range yingyong.Shezhi {
		Zifus[shezhi.Guilei] = shezhi.Guilei
		Zifus[shezhi.Bianma] = shezhi.Bianma
		Shezhis[shezhi.Bianma] = shezhi
	}
}
