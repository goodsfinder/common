package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/goodsfinder/common/trace"
)

// InitTrace http请求初始化traceID
// 从header获取保存到header和gin上下文中
func InitTrace(c *gin.Context) {
	if len(c.Request.Header.Get(trace.HeaderTraceIdKey)) == 0 {
		c.Request.Header.Set(trace.HeaderTraceIdKey, trace.ID())
	}
	c.Set(trace.ContextTraceId, c.Request.Header.Get(trace.HeaderTraceIdKey))
	// 响应header加返回traceId 需要在业务响应之前
	c.Header(trace.HeaderTraceIdKey, c.Request.Header.Get(trace.HeaderTraceIdKey))
	c.Next()
}
