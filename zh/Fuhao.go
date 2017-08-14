package zh

import (
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
)

func (zh *Zh) Zhistr(bianma string) string {
	return zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
}
