package engine

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

// POST is a shortcut for router.Handle("POST", path, handlers).
func (group *RouterGroup) Use(handle HandlerFunc, handlers ...gin.HandlerFunc) gin.IRoutes {
	return group.RouterGroup.Use(append(handlers, Decorate(handle))...)
}

// POST is a shortcut for router.Handle("POST", path, handlers).
func (group *RouterGroup) POST(relativePath string, handle HandlerFunc, handlers ...gin.HandlerFunc) gin.IRoutes {
	return group.RouterGroup.POST(relativePath, append(handlers, Decorate(handle))...)
}

// GET is a shortcut for router.Handle("GET", path, handlers).
func (group *RouterGroup) GET(relativePath string, handle HandlerFunc, handlers ...gin.HandlerFunc) gin.IRoutes {
	return group.RouterGroup.GET(relativePath, append(handlers, Decorate(handle))...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handlers).
func (group *RouterGroup) DELETE(relativePath string, handle HandlerFunc, handlers ...gin.HandlerFunc) gin.IRoutes {
	return group.RouterGroup.DELETE(relativePath, append(handlers, Decorate(handle))...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handlers).
func (group *RouterGroup) PATCH(relativePath string, handle HandlerFunc, handlers ...gin.HandlerFunc) gin.IRoutes {
	return group.RouterGroup.PATCH(relativePath, append(handlers, Decorate(handle))...)
}

// PUT is a shortcut for router.Handle("PUT", path, handlers).
func (group *RouterGroup) PUT(relativePath string, handle HandlerFunc, handlers ...gin.HandlerFunc) gin.IRoutes {
	return group.RouterGroup.PUT(relativePath, append(handlers, Decorate(handle))...)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handlers).
func (group *RouterGroup) OPTIONS(relativePath string, handle HandlerFunc, handlers ...gin.HandlerFunc) gin.IRoutes {
	return group.RouterGroup.OPTIONS(relativePath, append(handlers, Decorate(handle))...)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handlers).
func (group *RouterGroup) HEAD(relativePath string, handle HandlerFunc, handlers ...gin.HandlerFunc) gin.IRoutes {
	return group.RouterGroup.HEAD(relativePath, append(handlers, Decorate(handle))...)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (group *RouterGroup) Any(relativePath string, handle HandlerFunc, handlers ...gin.HandlerFunc) gin.IRoutes {
	return group.RouterGroup.Any(relativePath, append(handlers, Decorate(handle))...)
}

// Match registers a route that matches the specified methods that you declared.
func (group *RouterGroup) Match(methods []string, relativePath string, handle HandlerFunc, handlers ...gin.HandlerFunc) gin.IRoutes {
	return group.RouterGroup.Match(methods, relativePath, append(handlers, Decorate(handle))...)
}
