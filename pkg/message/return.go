package message

import (
	"github.com/gin-gonic/gin"
	"github.com/uncleyd/core/pkg/e"
	"net/http"
)

//请求成功的时候 使用该方法返回信息
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": e.MSG_OK,
		"msg":  "",
		"data": v,
	})
}

//请求成功的时候 使用该方法返回信息
func JSON(ctx *gin.Context, code int, msg string, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": v,
	})
}


//请求成功的时候 使用该方法返回信息
func AdminSuccess(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": e.MSG_OK,
		"msg":  "成功",
		"data": v,
	})
}

//请求失败的时候, 使用该方法返回信息
func Failed(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": e.MSG_ERR,
		"data": nil,
		"msg":  msg,
	})
}

//请求成功的时候 使用该方法返回信息
func JsonList(ctx *gin.Context, code int, count int, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":  code,
		"msg":   "",
		"count": count,
		"data":  data,
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

func HTML(ctx *gin.Context, name string, obj interface{}) {
	ctx.HTML(http.StatusOK, name, obj)
}
