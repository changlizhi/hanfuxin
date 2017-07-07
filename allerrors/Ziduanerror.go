package allerrors

import (
	"fmt"
	"time"
)

type Ziduanerror struct {
	Shijian time.Time
	Wenti   string
}

func (e Ziduanerror) Error() string {
	// 返回格式------时间:问题的json描述
	return fmt.Sprintf("Ziduanerror--------%v:%v", e.Shijian, e.Wenti)
}
