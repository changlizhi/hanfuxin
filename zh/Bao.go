package zh

import (
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
)

func (zh *Zh) Beego() string {
	//"github.com/astaxie/beego"
	ret :=
		zfzhi.Zhi.Syh() +
			zf.Zfs.Github(true) + zfzhi.Zhi.Dh() + zf.Zfs.Com(true) +
			zfzhi.Zhi.Xx() +
			zf.Zfs.Astaxie(true) +
			zfzhi.Zhi.Xx() +
			zf.Zfs.Beego(true) +
			zfzhi.Zhi.Syh() +
			zfzhi.Zhi.Hhf()
	return ret
}
