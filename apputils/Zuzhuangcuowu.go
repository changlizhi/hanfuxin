package apputils

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"strconv"
)

func Shengchengerrorchangdu(leixing string, lenchangliang int64, lenshiji int) string {
	buffer := bytes.Buffer{}

	buffer.WriteString(zfzhi.Zhi.Dkhz())
	buffer.WriteString(zf.Zfs.Leixing(false))
	buffer.WriteString(zfzhi.Zhi.Mh())

	buffer.WriteString(leixing)
	buffer.WriteString(zfzhi.Zhi.Dou())
	buffer.WriteString(zf.Zfs.Zuichang(false))
	buffer.WriteString(zfzhi.Zhi.Mh())
	buffer.WriteString(strconv.FormatInt(lenchangliang, 10))
	buffer.WriteString(zfzhi.Zhi.Dou())
	buffer.WriteString(zf.Zfs.Shiji(false))
	buffer.WriteString(zfzhi.Zhi.Mh())
	buffer.WriteString(strconv.Itoa(lenshiji))
	buffer.WriteString(zfzhi.Zhi.Dkhy())
	buffer.WriteString(zfzhi.Zhi.Dou())
	buffer.WriteString(zfzhi.Zhi.Hhf())
	return buffer.String()
}
