package baseinits

import (
	"hanfuxin/apputils"
	"hanfuxin/basemodels"
)

var Shezhis = make(map[string]basemodels.Tongyong)
var Mulus = make(map[string]basemodels.Tongyong)
var Biaos = make(map[string]basemodels.Tongyong)
var Lies = make(map[string]basemodels.Biaojiegou)
var Jichus = make(map[string]basemodels.Tongyong)
var Wangluos = make(map[string]basemodels.Tongyong)
var Fuhaos = make(map[string]basemodels.Tongyong)
var Wenzis = make(map[string]basemodels.Tongyong)
var Zifus = make(map[string]string)
var Gen basemodels.Gen
var Zhongwens = make(map[string]basemodels.Tongyong)
var Yingwens = make(map[string]basemodels.Tongyong)
func init() {
	chushihua_json()
}

func chushihua_json() {
	changliang := apputils.Changliangjson()
	yingyong := apputils.Yingyongjson()
	Gen = changliang.Gen

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
