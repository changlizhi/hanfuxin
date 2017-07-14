package allerrors

import (
	"encoding/json"
	"time"
)

type Ziduanerror struct {
	Shijian time.Time
	Wenti   string
}

func (e Ziduanerror) Error() string {
	// 返回格式------时间:问题的json描述
	ret, _ := json.Marshal(e)
	return string(ret)
}
