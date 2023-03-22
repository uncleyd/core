package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uncleyd/core/logger"
	"net/http"
	"strconv"
)

// message server instance
var Web *WebServer

type IController interface {
	Init(*gin.Engine)
}

const (
	MSG_OK          = 0
	MSG_ERR         = -1
	MSG_ERR_REQUEST = -2
	MSG_ERR_SYNC    = -3
)

// 包装为，自定义的GinContext
func BuildHandle(handler func(ctx *GinContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(GinContext)
		context.Context = c
		context.AuthCookie = &AuthCookie{}
		context.AuthCookie.Parse(context)
		//context.Request = &Request{}
		c.BindQuery(&context.Request)
		handler(context)
	}
}

// 通用请求结构
type Request struct {
	Guid    string `form:"guid"`     // 设备id
	Openid  string `form:"openid"`   // 微信openid
	Ver     string `form:"ver"`      // APP版本号
	Bid     string `form:"bid"`      // 应用包名
	Channel string `form:"channel" ` // 渠道
	Chl     string `form:"chl"`      // APP端（ios/android）
}

type GinContext struct {
	*gin.Context
	AuthCookie *AuthCookie
	Request    *Request
}

func (c *GinContext) QueryInt(name string, defaultValue int) int {
	if value, ok := c.GetQuery(name); ok {
		if v, e := strconv.Atoi(value); e == nil {
			return v
		}
	}
	return defaultValue
}

func (c *GinContext) QueryInt64(name string, defaultValue int64) int64 {
	if value, ok := c.GetQuery(name); ok {
		if v, e := strconv.ParseInt(value, 10, 64); e == nil {
			return v
		}
	}
	return defaultValue
}

func (c *GinContext) PostInt(name string, defaultValue int) int {
	if value, ok := c.GetPostForm(name); ok {
		if v, e := strconv.Atoi(value); e == nil {
			return v
		}
	}
	return defaultValue
}

func (c *GinContext) PostInt64(name string, defaultValue int64) int64 {
	if value, ok := c.GetPostForm(name); ok {
		if v, e := strconv.ParseInt(value, 10, 64); e == nil {
			return v
		}
	}
	return defaultValue
}

func (c *GinContext) SetCookie(name, value string) {
	c.Context.SetCookie(name, value, -1, "", "", false, false)
}

func (c *GinContext) CookieInt(name string, defaultValue int) int {
	if value, err := c.Context.Cookie(name); err == nil {
		if v, e := strconv.Atoi(value); e == nil {
			return v
		}
	}
	return defaultValue
}

func (c *GinContext) CookieInt64(name string, defaultValue int64) int64 {
	if value, err := c.Context.Cookie(name); err == nil {
		if v, e := strconv.ParseInt(value, 10, 64); e == nil {
			return v
		}
	}

	return defaultValue
}

func (c *GinContext) Error(msg string) {
	// 添加错误日志
	logger.Sugar.Errorf("%v err,err:%v", c.Context.Request.URL, msg)

	c.Context.JSON(http.StatusOK, gin.H{
		"code": MSG_ERR,
		"msg":  msg,
	})
}

func (c *GinContext) Ok() {
	c.JSON("")
}

func (c *GinContext) JSON(value ...interface{}) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code": MSG_OK,
		"msg":  value[0],
		"data": value[1],
	})
}

func (c *GinContext) JSONWithCount(msg string, count int, data interface{}) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code":  MSG_OK,
		"msg":   msg,
		"data":  data,
		"count": count,
	})
}

func (c *GinContext) JSONS(msg string, count int, data interface{}) {
	out := make(map[string]interface{})
	out["msg"] = msg
	out["count"] = count
	if data != nil {
		out["data"] = data
	}

	c.Context.JSON(http.StatusOK,
		gin.H{
			"code": MSG_OK,
			"msg":  msg,
			"data": out,
		})
}

func (c *GinContext) JSONAli(code int, msg interface{}) {
	c.Context.JSON(http.StatusOK, msg)
}

func (c *GinContext) Pagination(total int, rows interface{}) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code":  MSG_OK,
		"msg":   "",
		"count": total,
		"data":  rows,
	})
}

func (c *GinContext) PaginationMsg(total int, rows interface{}, msg string) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code":  MSG_OK,
		"msg":   msg,
		"total": total,
		"rows":  rows,
	})
}

func (c *GinContext) Object(obj interface{}) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code": MSG_OK,
		"data": obj,
	})
}

func (c *GinContext) HTML(name string, obj interface{}) {
	c.Context.HTML(http.StatusOK, name, obj)
}

func (c *GinContext) Redirect(url string) {
	c.Context.Redirect(http.StatusMovedPermanently, url)
}

func (c *GinContext) JsonMsgAndEncrypt(msg string, data interface{}) {
	resp := gin.H{
		"code": MSG_OK,
		"msg":  msg,
		"data": data,
	}

	content, err := json.Marshal(resp)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	eData := Encrypt(content)

	fmt.Println("eData:", string(eData))
	c.Context.Writer.WriteString(string(eData))
}
func (c *GinContext) ErrorAndEncrypt(msg string) {
	// 添加错误日志
	logger.Sugar.Errorf("%v err,err:%v", c.Context.Request.URL, msg)

	resp := gin.H{
		"code": MSG_ERR,
		"msg":  msg,
		"data": nil,
	}

	content, err := json.Marshal(resp)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Context.Writer.WriteString(string(Encrypt(content)))
}
