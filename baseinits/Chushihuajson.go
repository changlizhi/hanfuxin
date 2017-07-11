package baseinits

import (
	"hanfuxin/basemodels"
	"hanfuxin/baseutils"
	"hanfuxin/zf"
)

var Cuowus = make(map[string]basemodels.Tongyong)
var Tishis = make(map[string]basemodels.Tongyong)
var Shujukus = make(map[string]basemodels.Tongyong)
var Chushihuas = make(map[string]basemodels.Tongyong)

func init() {
	chushihua_json()
}
func chushihua_json() {
	zf := zf.Zf{}
	shezhi := baseutils.Shezhijson()
	for _, c := range shezhi.Chushihua {
		Chushihuas[c.Bianma] = c
	}
	for _, s := range shezhi.Shujuku {
		Shujukus[s.Bianma] = s
	}

	guojihua := baseutils.Guojihuajson(Chushihuas[zf.Yuyan(false)].Zhi)
	Shezhiguojihua(guojihua)
}
func Shezhiguojihua(guojihua *basemodels.Guojihua) {
	for _, cw := range guojihua.Cuowu {
		Cuowus[cw.Bianma] = cw
	}
	for _, ts := range guojihua.Tishi {
		Tishis[ts.Bianma] = ts
	}
}
