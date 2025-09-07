package format

import (
	"time"
)

func DayStrToTimestamp(day string) (timestamp int64) {
	// 解析日期字符串，使用本地时区
	t, err := time.ParseInLocation("2006-01-02", day, time.Local)
	if err != nil {
		return 0
	}

	// 转换为时间戳（秒）
	timestamp = t.Unix()
	return timestamp
}
