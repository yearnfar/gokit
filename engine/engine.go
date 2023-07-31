package engine

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func New() *Engine {
	e := &Engine{
		Engine: gin.New(),
	}
	return e
}

type Engine struct {
	*gin.Engine
}

// GET 注册GET请求的路由
func (e *Engine) GET(relativePath string, handler HandlerFunc, middlewares ...gin.HandlerFunc) {
	e.Engine.GET(relativePath, append(middlewares, Decorate(handler))...)
}

// POST 注册POST请求的路由
func (e *Engine) POST(relativePath string, handler HandlerFunc, middlewares ...gin.HandlerFunc) {
	e.Engine.POST(relativePath, append(middlewares, Decorate(handler))...)
}

// PUT 注册PUT请求的路由
func (e *Engine) PUT(relativePath string, handler HandlerFunc, middlewares ...gin.HandlerFunc) {
	e.Engine.PUT(relativePath, append(middlewares, Decorate(handler))...)
}

// PATCH 注册PATCH请求的路由
func (e *Engine) PATCH(relativePath string, handler HandlerFunc, middlewares ...gin.HandlerFunc) {
	e.Engine.PATCH(relativePath, append(middlewares, Decorate(handler))...)
}

// DELETE 注册DELETE请求的路由
func (e *Engine) DELETE(relativePath string, handler HandlerFunc, middlewares ...gin.HandlerFunc) {
	e.Engine.DELETE(relativePath, append(middlewares, Decorate(handler))...)
}

// Handle 自己指定 HTTP 方法注册路由
func (e *Engine) Handle(httpMethod, relativePath string, handler HandlerFunc, middlewares ...gin.HandlerFunc) {
	e.Engine.Handle(httpMethod, relativePath, append(middlewares, Decorate(handler))...)
}

// Group 注册组路由
func (e *Engine) Group(relativePath string, middlewares ...gin.HandlerFunc) *RouterGroup {
	g := e.Engine.Group(relativePath, middlewares...)
	return &RouterGroup{g}
}

// Decorate 把我们的 HandlerFunc 包装成 Gin 的 HandlerFunc
func Decorate(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := newContext(c)
		defer func() {
			if err := recover(); err != nil {
				// if e, ok := err.(codeError); ok {
				// 	if err2 := ctx.Message(e.Message(), e.Code()); err2 != nil {
				// 		log.Error(ctx, "message: ", err2)
				// 	}
				// } else {
				// 	panic(err)
				// }
			}
		}()
		var body []byte
		if ctx.Request.Method == "POST" && ctx.ContentType() == gin.MIMEJSON {
			body, _ = ioutil.ReadAll(ctx.Request.Body)
			ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		}
		if err := h(ctx); err != nil {
			// errorHandlerFunc(ctx, err)
		}
	}
}

type HandlerFunc func(ctx *Context) error
