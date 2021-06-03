// 请求频繁黑名单
package rdmodels

import (
	"strings"
	"time"
)

const (
	BLACK_LIST = "black_list" // 黑名单
)

// 是否存在黑名单中
func ExistsBlackList(openid string) bool {
	in, err := RdClient.Exists(strings.Join([]string{BLACK_LIST, openid}, ":")).Result()
	if err != nil {
		return false
	}

	return in > 0
}

// 设置请求黑名单
func SetBlackList(openid string, times int) error {
	_, err := RdClient.Set(strings.Join([]string{BLACK_LIST, openid}, ":"), 1, time.Duration(times)*time.Second).Result()
	if err != nil {
		return err
	}
	return nil
}
