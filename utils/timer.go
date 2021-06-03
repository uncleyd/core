package utils

import "time"

func NowMillisecond() int {
	return int(time.Now().UnixNano() / 1000000)
}

func GetDayUnix(nowUnix int64) int64 {
	return nowUnix - (nowUnix+8*3600)%86400
}

// 取过去的时间戳与现在的时间戳相差的天数，只要今天没登入就算一天，例如：2020-07-07 10:35:00 登入了，到2020-07-08 00:00:01，就算一天没登入，算08这一天还没登入过
func TimeSub(timeunix int64) int {
	timestamp := "2006-01-02"

	a := time.Unix(time.Now().Unix(), 0).Format(timestamp)
	b := time.Unix(timeunix, 0).Format(timestamp)

	a1, _ := time.ParseInLocation(timestamp, a, time.Local)
	b1, _ := time.ParseInLocation(timestamp, b, time.Local)

	return int(a1.Sub(b1).Hours() / 24)
}

// 获得x日O点时间时间戳 day=0今日，day=1明天....
func GetDayZeroTime(day int) int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.AddDate(0, 0, day).Unix()
}

// 获得今日24点时间剩余秒数
func GetDayZeroSurplusTime() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.AddDate(0, 0, 1).Unix() - time.Now().Unix()
}
