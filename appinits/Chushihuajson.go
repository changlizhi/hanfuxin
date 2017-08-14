package appinits

import (
	"hanfuxin/zf"
)

var Cuowus = make(map[string]Tongyong)
var Tishis = make(map[string]Tongyong)
var Shujukus = make(map[string]Tongyong)
var Chushihuas = make(map[string]Tongyong)

func init() {
	chushihua_json()
}
func chushihua_json() {
	shezhi := Shezhijson()
	for _, c := range shezhi.Chushihua {
		Chushihuas[c.Bianma] = c
	}
	for _, s := range shezhi.Shujuku {
		Shujukus[s.Bianma] = s
	}

	guojihua := Guojihuajson(Chushihuas[zf.Zfs.Yuyan(false)].Zhi)
	Shezhiguojihua(guojihua)
}
func Shezhiguojihua(guojihua *Guojihua) {
	for _, cw := range guojihua.Cuowu {
		Cuowus[cw.Bianma] = cw
	}
	for _, ts := range guojihua.Tishi {
		Tishis[ts.Bianma] = ts
	}
}
