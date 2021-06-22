package server

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

const (
	AppKey    = "animal"
	AppSecret = "-=bj64ac"
)

// MD5 組合加密
func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}

	//return func(c *gin.Context) {
	//	err := verifySign(c)
	//	if err != nil {
	//		ctx := GinContext{
	//			Context: c,
	//		}
	//		ctx.JSON(err.Error(), nil)
	//		c.Abort()
	//		return
	//	}
	//	c.Next()
	//}
}

// 驗證簽名
func verifySign(c *gin.Context) error {
	_ = c.Request.ParseForm()
	req := c.Request.URL
	sn := c.Query("sign")

	// 驗證簽名
	if sn == "" || sn != createSign(req.String()) {
		return errors.New("sn Error")
	}

	return nil
}

// 創建簽名
func createSign(url string) string {
	// 自定義 MD5 組合
	return MD5(AppKey + createEncryptStr(url) + AppSecret)
}

// 获取待签字符串
func createEncryptStr(params string) string {
	str := strings.Split(params, "&sign")
	return str[0]
}
