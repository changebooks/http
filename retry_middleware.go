package http

import "net/http"

// 超过最大重试次数，不重试 - retries (已重试次数) >= maxRetries (最大重试次数)
// 丢失连接，重试 - NoHttpResponse
// SSL异常，不重试 - TLSHandshake
// 超时，不重试 - InterruptedIO
// 目标服务器不可达，不重试 - UnknownHost
// 拒绝连接，不重试 - ConnectTimeout
// 支持幂等，重试 - PUT、GET、HEAD
func RetryMiddleware(maxRetries int, retries int, req *http.Request, resp *http.Response, err error) bool {
	return false
}
