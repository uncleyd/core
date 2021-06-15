package rpc

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var SERVER *gin.Engine

func init() {
	SERVER = gin.Default()
}

const (
	MSG_OK          = 0
	MSG_ERR         = -1
	MSG_ERR_REQUEST = -2
	MSG_ERR_SYNC    = -3
)

//请求成功的时候 使用该方法返回信息
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": MSG_OK,
		"msg":  "",
		"data": v,
	})
}

//请求失败的时候, 使用该方法返回信息
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": MSG_ERR,
		"data": nil,
		"msg":  v,
	})
}

// 返回页面
func SuccessHtml(ctx *gin.Context, html, title string) {
	ctx.HTML(http.StatusOK, html, gin.H{
		"title": title,
	})
}

// 重定向
func SuccessRedirect(ctx *gin.Context, html string) {
	ctx.Redirect(http.StatusFound, html)
}

// 获取用户IP地址

//获取请求基本参数
func GetReq(ctx *gin.Context) Req {
	req := Req{map[string]interface{}{
		"guid":      ctx.Query("guid"),      // 设备id
		"systemver": ctx.Query("systemver"), // 固件版本号
		"ver":       ctx.Query("ver"),       // APP版本名
		"bid":       ctx.Query("bid"),       // 应用包名
		"device":    ctx.Query("bid"),       // 设备型号
		"chl":       ctx.Query("chl"),       // APP端（android/ios）
		"language":  ctx.Query("language"),  // 系统语言
		"locale":    ctx.Query("locale"),    // 国家域名
		"zone":      ctx.Query("zone"),      // 国家时区编码
		"net":       ctx.Query("net"),       // 网络类型（1是mobile 2是wifi 默认99）
		"mt":        ctx.Query("mt"),        // 网络类型 默认值2
		"channel":   ctx.Query("channel"),   // 应用渠道
		"hw":        ctx.Query("hw"),        // 是否华为设备（1华为，其他非华为）1)
		"ip":        ctx.ClientIP(),
	}}
	return req
}

type Req struct {
	Data map[string]interface{}
}

type Resp struct {
	Data interface{}
}

type RespMap map[string]interface{}

func RpcxReqModel(ctx *gin.Context, servicepath, serviceMethod string, p []string) {
	req := GetReq(ctx)

	// 接口专用参数
	if p != nil {
		for _, v := range p {
			req.Data[v] = ctx.Query(v)
		}
	}

	var resp = &Resp{}

	err := RpcxCall(req, resp, servicepath, serviceMethod)
	if err != nil {
		Failed(ctx, err.Error())
		return
	}

	Success(ctx, resp.Data)
}

func RpcxReqModelEx(ctx *gin.Context, servicepath, serviceMethod string, p []string, m map[string]interface{}) {
	req := GetReq(ctx)

	// 接口专用参数
	if p != nil {
		for _, v := range p {
			req.Data[v] = ctx.Query(v)
		}
	}

	// 入参
	if len(m) > 0 {
		for k, v := range m {
			req.Data[k] = v
		}
	}

	var resp = &Resp{}

	err := RpcxCall(req, resp, servicepath, serviceMethod)
	if err != nil {
		Failed(ctx, err.Error())
		return
	}

	Success(ctx, resp.Data)
}
