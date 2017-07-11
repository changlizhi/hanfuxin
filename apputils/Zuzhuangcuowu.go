package apputils

import (
	"bytes"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"strconv"
)

func Shengchengerrorchangdu(leixing string, lenchangliang int64, lenshiji int) string {
	buffer := bytes.Buffer{}
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	dkhz := zfzhi.Dakuohaozuozhi()
	dkhy := zfzhi.Dakuohaoyouzhi()
	douhao := zfzhi.Douhaozhi()
	mh := zfzhi.Maohaozhi()
	hhf := zfzhi.Huanhangfuzhi()

	buffer.WriteString(dkhz)
	buffer.WriteString(zf.Leixing(false))
	buffer.WriteString(mh)

	buffer.WriteString(leixing)
	buffer.WriteString(douhao)
	buffer.WriteString(zf.Zuichang(false))
	buffer.WriteString(mh)
	buffer.WriteString(strconv.FormatInt(lenchangliang, 10))
	buffer.WriteString(douhao)
	buffer.WriteString(zf.Shiji(false))
	buffer.WriteString(mh)
	buffer.WriteString(strconv.Itoa(lenshiji))
	buffer.WriteString(dkhy)
	buffer.WriteString(douhao)
	buffer.WriteString(hhf)
	return buffer.String()
}
