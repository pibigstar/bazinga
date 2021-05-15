package timex

import "time"

// 获取毫秒时间戳
func UnixMill() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
