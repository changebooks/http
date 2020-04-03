package http

import (
	"crypto/tls"
	"errors"
	"net/http/httptrace"
	"time"
)

func NewTrace(s *Schema) (*httptrace.ClientTrace, error) {
	if s == nil {
		return nil, errors.New("schema can't be nil")
	}

	if s.Elapsed == nil {
		s.Elapsed = &Elapsed{}
		s.Elapsed.Start = time.Now()
	}

	return &httptrace.ClientTrace{
		// GetConn is called before a connection is created or
		// retrieved from an idle pool. The hostPort is the
		// "host:port" of the target or proxy. GetConn is called even
		// if there's already an idle cached connection available.
		GetConn: func(hostPort string) {
			s.HostPort = hostPort
		},

		// GotConn is called after a successful connection is
		// obtained. There is no hook for failure to obtain a
		// connection; instead, use the error from
		// Transport.RoundTrip.
		GotConn: func(info httptrace.GotConnInfo) {
			s.Reused = info.Reused
		},

		// DNSStart is called when a DNS lookup begins.
		DNSStart: func(info httptrace.DNSStartInfo) {
			s.Elapsed.DNSStart = time.Now()
			s.Host = info.Host
		},

		// DNSDone is called when a DNS lookup ends.
		DNSDone: func(info httptrace.DNSDoneInfo) {
			s.Elapsed.DNSDone = time.Now()
			s.Elapsed.DNSLookup = s.Elapsed.DNSDone.Sub(s.Elapsed.DNSStart)
		},

		// ConnectStart is called when a new connection's Dial begins.
		// If net.Dialer.DualStack (IPv6 "Happy Eyeballs") support is
		// enabled, this may be called multiple times.
		ConnectStart: func(network, addr string) {
			s.Elapsed.ConnectStart = time.Now()
			s.Ip = addr
		},

		// ConnectDone is called when a new connection's Dial
		// completes. The provided err indicates whether the
		// connection completedly successfully.
		// If net.Dialer.DualStack ("Happy Eyeballs") support is
		// enabled, this may be called multiple times.
		ConnectDone: func(network, addr string, err error) {
			s.Elapsed.ConnectDone = time.Now()
			s.Elapsed.Connect = s.Elapsed.ConnectDone.Sub(s.Elapsed.ConnectStart)
		},

		// TLSHandshakeStart is called when the TLS handshake is started. When
		// connecting to a HTTPS site via a HTTP proxy, the handshake happens after
		// the CONNECT request is processed by the proxy.
		TLSHandshakeStart: func() {
			s.Tls = true
			s.Elapsed.TLSHandshakeStart = time.Now()
		},

		// TLSHandshakeDone is called after the TLS handshake with either the
		// successful handshake's connection state, or a non-nil error on handshake
		// failure.
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			s.Elapsed.TLSHandshakeDone = time.Now()
			s.Elapsed.TLSHandshake = s.Elapsed.TLSHandshakeDone.Sub(s.Elapsed.TLSHandshakeStart)
		},

		// WroteRequest is called with the result of writing the
		// request and any body. It may be called multiple times
		// in the case of retried requests.
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			s.Elapsed.ServerStart = time.Now()
		},

		// GotFirstResponseByte is called when the first byte of the response
		// headers is available.
		GotFirstResponseByte: func() {
			s.Elapsed.ServerDone = time.Now()
			s.Elapsed.Server = s.Elapsed.ServerDone.Sub(s.Elapsed.ServerStart)
			s.Elapsed.TransferStart = s.Elapsed.ServerDone
		},
	}, nil
}
