package http

import (
	"errors"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Talk struct {
	Client          http.Client
	baseUrl         *BaseUrl
	maxRetries      int               // 最大重试次数
	timeout         time.Duration     // 缺省的限时
	header          http.Header       // 所有请求共同Header
	params          map[string]string // 所有请求共同参数
	useTrace        bool              // 是否记录请求详情
	retryMiddleware func(maxRetries int, retries int, req *http.Request, resp *http.Response, err error) bool
}

func (x *Talk) Request(req *http.Request, timeout time.Duration) *Schema {
	if timeout > 0 {
		x.Client.Timeout = timeout
	} else {
		x.Client.Timeout = x.timeout
	}

	return Do(&x.Client, req, x.useTrace, x.maxRetries, x.retryMiddleware)
}

func (x *Talk) NewRequest(method string, path string, params map[string]string) (*http.Request, string, error) {
	url := x.JoinUrl(path, params)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, url, err
	}

	if x.header != nil {
		req.Header = x.header
	}

	return req, url, nil
}

func (x *Talk) JoinUrl(path string, params map[string]string) string {
	if path != "" {
		path = strings.TrimLeft(path, "/")
	}

	return x.baseUrl.GetUrl() + UrlPath(path, x.MergeParams(params))
}

func (x *Talk) MergeParams(p map[string]string) map[string]string {
	if x.params == nil {
		return p
	}

	if p == nil {
		return x.params
	}

	r := x.params
	for k, v := range p {
		r[k] = v
	}

	return r
}

func (x *Talk) GetBaseUrl() *BaseUrl {
	return x.baseUrl
}

func (x *Talk) GetMaxRetries() int {
	return x.maxRetries
}

func (x *Talk) GetTimeout() time.Duration {
	return x.timeout
}

func (x *Talk) GetHeader() http.Header {
	return x.header
}

func (x *Talk) GetParams() map[string]string {
	return x.params
}

func (x *Talk) GetUseTrace() bool {
	return x.useTrace
}

func (x *Talk) GetRetryMiddleware() func(maxRetries int, retries int, req *http.Request, resp *http.Response, err error) bool {
	return x.retryMiddleware
}

type TalkBuilder struct {
	mu              sync.Mutex // ensures atomic writes; protects the following fields
	baseUrl         *BaseUrl
	maxRetries      int
	timeout         time.Duration
	header          http.Header
	params          map[string]string
	useTrace        bool
	retryMiddleware func(maxRetries int, retries int, req *http.Request, resp *http.Response, err error) bool
}

func (x *TalkBuilder) Build() (*Talk, error) {
	if x.baseUrl == nil {
		return nil, errors.New("base url can't be nil")
	}

	if x.maxRetries < 0 {
		return nil, errors.New("max retries can't be less than 0")
	}

	if x.timeout < 0 {
		return nil, errors.New("timeout can't be less than 0")
	}

	retryMiddleware := x.retryMiddleware
	if retryMiddleware == nil {
		retryMiddleware = RetryMiddleware
	}

	return &Talk{
		Client:          http.Client{},
		baseUrl:         x.baseUrl,
		maxRetries:      x.maxRetries,
		timeout:         x.timeout,
		header:          x.header,
		params:          x.params,
		useTrace:        x.useTrace,
		retryMiddleware: retryMiddleware,
	}, nil
}

func (x *TalkBuilder) SetBaseUrl(u *BaseUrl) *TalkBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.baseUrl = u
	return x
}

func (x *TalkBuilder) SetMaxRetries(retries int) *TalkBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.maxRetries = retries
	return x
}

func (x *TalkBuilder) SetTimeout(d time.Duration) *TalkBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.timeout = d
	return x
}

func (x *TalkBuilder) SetHeader(key string, value string) *TalkBuilder {
	if key = strings.TrimSpace(key); key == "" {
		return x
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	if x.header == nil {
		x.header = http.Header{}
	}

	x.header.Set(key, value)
	return x
}

func (x *TalkBuilder) AddHeader(key string, value string) *TalkBuilder {
	if key = strings.TrimSpace(key); key == "" {
		return x
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	if x.header == nil {
		x.header = http.Header{}
	}

	x.header.Add(key, value)
	return x
}

func (x *TalkBuilder) AddParam(key string, value string) *TalkBuilder {
	if key = strings.TrimSpace(key); key == "" {
		return x
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	if x.params == nil {
		x.params = make(map[string]string)
	}

	x.params[key] = value
	return x
}

func (x *TalkBuilder) SetUseTrace(useTrace bool) *TalkBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.useTrace = useTrace
	return x
}

func (x *TalkBuilder) SetRetryMiddleware(f func(maxRetries int, retries int, req *http.Request, resp *http.Response, err error) bool) *TalkBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.retryMiddleware = f
	return x
}
