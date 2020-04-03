package http

import (
	"fmt"
	"net/http"
)

type Schema struct {
	StatusCode int
	Data       []byte
	Error      error
	ReadError  error // 读取 Response Body 错误
	FreeError  error // 关闭 Response Body 错误
	Tls        bool  // Https ?
	Host       string
	HostPort   string
	Ip         string
	Reused     bool     // 此连接以前用于另一个Http请求
	Retries    int      // 重试次数，无重试 = 0
	Elapsed    *Elapsed // 耗时
	Request    *http.Request
	Response   *http.Response
}

func (x *Schema) ToString() string {
	return fmt.Sprintf("StatusCode: %v\n"+
		"Data:       %v\n"+
		"Error:      %v\n"+
		"ReadError:  %v\n"+
		"FreeError:  %v\n"+
		"Tls:        %v\n"+
		"Host:       %v\n"+
		"HostPort:   %v\n"+
		"Ip:         %v\n"+
		"Reused:     %v\n"+
		"Retries:    %v\n"+
		"Request:    %v\n"+
		"Response:   %v\n",
		x.StatusCode, x.Data, x.Error, x.ReadError, x.FreeError, x.Tls, x.Host, x.HostPort, x.Ip, x.Reused, x.Retries, x.Request, x.Response,
	)
}
