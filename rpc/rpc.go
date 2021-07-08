package rpc

import (
	"github.com/gin-gonic/gin"
	"github.com/uncleyd/core/logger"
	"github.com/uncleyd/core/pkg/message"
	"github.com/uncleyd/core/server"
)

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

func RpcxReqModel(ctx *server.GinContext, servicepath, serviceMethod string, p []string) {
	logger.Sugar.Debugf("%v:%v  on '%s'", servicepath, serviceMethod, ctx.Context.Request.URL.String())
	req := GetReq(ctx.Context)

	// 接口专用参数
	if p != nil {
		for _, v := range p {
			req.Data[v] = ctx.Query(v)
		}
	}

	var resp = &Resp{}

	err := RpcxCall(req, resp, servicepath, serviceMethod)
	if err != nil {
		logger.Sugar.Errorf("%v:%v  on '%s' err:%v", servicepath, serviceMethod, ctx.Context.Request.URL.String(), err)
		message.Failed(ctx.Context, err.Error())
		return
	}

	message.Success(ctx.Context, resp.Data)
}

func RpcxReqModelEx(ctx *server.GinContext, servicepath, serviceMethod string, p []string, m map[string]interface{}) {
	logger.Sugar.Debugf("%v:%v  on '%s'", servicepath, serviceMethod, ctx.Context.Request.URL.String())
	req := GetReq(ctx.Context)

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
		logger.Sugar.Errorf("%v:%v  on '%s' err:%v", servicepath, serviceMethod, ctx.Context.Request.URL.String(), err)
		message.Failed(ctx.Context, err.Error())
		return
	}
	message.Success(ctx.Context, resp.Data)
}

// 后台使用
func RpcxReqModelJsonList(ctx *server.GinContext, servicepath, serviceMethod string, p []string) {
	logger.Sugar.Debugf("%v:%v  on '%s'", servicepath, serviceMethod, ctx.Context.Request.URL.String())
	req := Req{
		Data: map[string]interface{}{},
	}

	// 接口专用参数
	if p != nil {
		for _, v := range p {
			req.Data[v] = ctx.Query(v)
		}
	}

	var resp = &RespJsonList{}

	err := RpcxCall(req, resp, servicepath, serviceMethod)
	if err != nil {
		logger.Sugar.Errorf("%v:%v  on '%s' err:%v", servicepath, serviceMethod, ctx.Context.Request.URL.String(), err)
		message.Failed(ctx.Context, err.Error())
		return
	}
	message.JsonList(ctx.Context, resp.Code, resp.Count, resp.Data)
}

// 后台请求单个服务
func RpcxReqAdminModel(ctx *server.GinContext, servicepath, serviceMethod string, p []string) {
	logger.Sugar.Debugf("%v:%v  on '%s'", servicepath, serviceMethod, ctx.Context.Request.URL.String())
	req := Req{
		Data: map[string]interface{}{},
	}

	// 接口专用参数
	if p != nil {
		for _, v := range p {
			req.Data[v] = ctx.Query(v)
		}
	}

	var resp = &Resp{}

	err := RpcxAdminCall(req, resp, servicepath, serviceMethod)
	if err != nil {
		logger.Sugar.Errorf("%v:%v  on '%s' err:%v", servicepath, serviceMethod, ctx.Context.Request.URL.String(), err)
		message.Failed(ctx.Context, err.Error())
		return
	}
	message.JSON(ctx.Context, resp.Code, resp.Msg, resp.Data)
}

// 广播请求所有服务
func RpcxReqBroadcastModel(ctx *server.GinContext, servicepath, serviceMethod string, p []string) {
	logger.Sugar.Debugf("%v:%v  on '%s'", servicepath, serviceMethod, ctx.Context.Request.URL.String())
	req := Req{
		Data: map[string]interface{}{},
	}

	// 接口专用参数
	if p != nil {
		for _, v := range p {
			req.Data[v] = ctx.Query(v)
		}
	}

	var resp = &Resp{}

	err := RpcxBroadcast(req, resp, servicepath, serviceMethod)
	if err != nil {
		logger.Sugar.Errorf("%v:%v  on '%s' err:%v", servicepath, serviceMethod, ctx.Context.Request.URL.String(), err)
		message.Failed(ctx.Context, err.Error())
		return
	}
	message.JSON(ctx.Context, resp.Code, resp.Msg, resp.Data)
}

// 请求单个服务返回HTML
func RpcxReqAdminModelHtml(ctx *server.GinContext, servicepath, serviceMethod string, p []string, name string) {
	logger.Sugar.Debugf("%v:%v  on '%s'", servicepath, serviceMethod, ctx.Context.Request.URL.String())
	req := Req{
		Data: map[string]interface{}{},
	}

	// 接口专用参数
	if p != nil {
		for _, v := range p {
			req.Data[v] = ctx.Query(v)
		}
	}

	var resp = &Resp{}

	err := RpcxAdminCall(req, resp, servicepath, serviceMethod)
	if err != nil {
		logger.Sugar.Errorf("%v:%v  on '%s' err:%v", servicepath, serviceMethod, ctx.Context.Request.URL.String(), err)
		message.Failed(ctx.Context, err.Error())
		return
	}
	message.HTML(ctx.Context, name, resp.Data)
}
