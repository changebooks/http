package http

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func HasScheme(url string) bool {
	return strings.HasPrefix(url, SchemeHttp+"://") || strings.HasPrefix(url, SchemeHttps+"://")
}

func IsDefaultPort(scheme string, port int) bool {
	if port <= 0 {
		return true
	}

	if port == PortHttp && (scheme == SchemeHttp || scheme == "") {
		return true
	}

	if port == PortHttps && scheme == SchemeHttps {
		return true
	}

	return false
}

type BaseUrl struct {
	url    string // join scheme, host, port
	scheme string // 协议，如：http or https
	host   string // 域名，如：t.cn
	port   int    // 端口，如：80 or 443
}

func (x *BaseUrl) IsDefaultPort() bool {
	return IsDefaultPort(x.scheme, x.port)
}

func (x *BaseUrl) GetUrl() string {
	return x.url
}

func (x *BaseUrl) GetScheme() string {
	return x.scheme
}

func (x *BaseUrl) GetHost() string {
	return x.host
}

func (x *BaseUrl) GetPort() int {
	return x.port
}

type BaseUrlBuilder struct {
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	scheme string
	host   string
	port   int
}

func (x *BaseUrlBuilder) Build() (*BaseUrl, error) {
	if x.host == "" {
		return nil, errors.New("host can't be empty")
	}

	host := x.host
	scheme := x.scheme

	// 从host分离scheme
	num := strings.Index(host, "://")
	if num >= 0 {
		if scheme == "" {
			scheme = host[:num]
		}

		host = host[num+3:]
	}

	if host == "" {
		return nil, errors.New("no scheme's host can't be empty")
	}

	port := x.port

	// 从host分离port
	num = strings.Index(host, ":")
	if num >= 0 {
		if port <= 0 {
			p, err := strconv.ParseInt(host[num+1:], 10, 32)
			if err != nil {
				return nil, err
			}

			port = int(p)
		}

		host = host[:num]
	}

	if host == "" {
		return nil, errors.New("no port's host can't be empty")
	}

	if scheme != "" {
		if scheme != SchemeHttp && scheme != SchemeHttps {
			return nil, fmt.Errorf(`unsupported scheme "%s", must be %s or %s`, scheme, SchemeHttp, SchemeHttps)
		}
	} else {
		scheme = SchemeHttp
	}

	url := ""
	if IsDefaultPort(scheme, port) {
		url = fmt.Sprintf("%s://%s/", scheme, host)
	} else {
		url = fmt.Sprintf("%s://%s:%d/", scheme, host, port)
	}

	return &BaseUrl{
		url:    url,
		scheme: scheme,
		host:   host,
		port:   port,
	}, nil
}

func (x *BaseUrlBuilder) SetScheme(s string) *BaseUrlBuilder {
	s = strings.TrimRight(strings.TrimSpace(s), "://")

	x.mu.Lock()
	defer x.mu.Unlock()

	x.scheme = s
	return x
}

func (x *BaseUrlBuilder) SetHost(s string) *BaseUrlBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.host = s
	return x
}

func (x *BaseUrlBuilder) SetPort(p int) *BaseUrlBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.port = p
	return x
}
