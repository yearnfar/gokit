package engine

import (
	"github.com/gin-gonic/gin"
)

// newContext 封装上下文对象
func newContext(c *gin.Context) *Context {
	ctx := &Context{Context: c}
	return ctx
}

// Context 上下文对象，对 gin.Context 进行扩展
type Context struct {
	*gin.Context
}

// Response 输出响应结果
// func (c *Context) Response(data interface{}, msg string, code int) error {
// 	obj := &JsonResult{
// 		Code:    code,
// 		Msg:     msg,
// 		Data:    data,
// 		TraceId: trace.GetTraceID(c),
// 	}
// 	if code > 0 {
// 		c.JSON(errStatusCode, obj)
// 	} else {
// 		c.JSON(200, obj)
// 	}
// 	return nil
// }

// // Message 输出错误消息
// func (c *Context) Message(msg string, code ...int) error {
// 	if len(code) > 0 {
// 		return c.Response(nil, msg, code[0])
// 	}
// 	return c.Response(nil, msg, MsgErr)
// }

// // MsgOK 输出成功消息
// func (c *Context) MsgOK(msg ...string) error {
// 	if len(msg) > 0 {
// 		return c.Response(nil, msg[0], MsgOK)
// 	} else {
// 		return c.Response(nil, "ok", MsgOK)
// 	}
// }

// // ServeJSON 输出成功结果
// func (c *Context) ServeJSON(data interface{}) error {
// 	return c.Response(data, "", MsgOK)
// }
