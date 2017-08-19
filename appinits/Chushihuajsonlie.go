package appinits

func init() {
	Shezhilie()
	Guojihualie()
}

var Shujukus = make(map[string]Tongyong)
var Chushihuas = make(map[string]Tongyong)

func Shezhilie() {
	shezhi := Shezhijson()
	for _, c := range shezhi.Chushihua {
		Chushihuas[c.Bianma] = c
	}
	for _, s := range shezhi.Shujuku {
		Shujukus[s.Bianma] = s
	}
}

var Cuowus = make(map[string]Tongyong)
var Tishis = make(map[string]Tongyong)

func Guojihualie() {
	guojihua := Guojihuajson()
	for _, cw := range guojihua.Cuowu {
		Cuowus[cw.Bianma] = cw
	}
	for _, ts := range guojihua.Tishi {
		Tishis[ts.Bianma] = ts
	}
}
