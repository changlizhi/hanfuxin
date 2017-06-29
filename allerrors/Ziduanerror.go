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
	return fmt.Sprintf("Ziduanerror--------%v:%v", e.Shijian, e.Wenti)
}
