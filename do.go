package http

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"
)

func Do(c *http.Client, req *http.Request, useTrace bool, maxRetries int,
	retryMiddleware func(maxRetries int, retries int, req *http.Request, resp *http.Response, err error) bool) *Schema {
	s := &Schema{}

	if c == nil {
		s.Error = errors.New("http client can't be nil")
		return s
	}

	if req == nil {
		s.Error = errors.New("http request can't be nil")
		return s
	}

	if useTrace {
		t, err := NewTrace(s)
		if err != nil {
			s.Error = err
			return s
		}

		req = req.WithContext(httptrace.WithClientTrace(req.Context(), t))
	}

	retries := 0
	for {
		resp, err := c.Do(req)
		if err != nil && retryMiddleware != nil && maxRetries > 0 && retryMiddleware(maxRetries, retries, req, resp, err) {
			retries++
			s.Elapsed = nil
			continue
		}

		s.Response = resp
		s.Error = err
		break
	}

	s.Retries = retries
	s.Request = req

	if s.Elapsed == nil {
		s.Elapsed = &Elapsed{}
	}

	s.Elapsed.TransferDone = time.Now()
	s.Elapsed.Transfer = s.Elapsed.TransferDone.Sub(s.Elapsed.TransferStart)

	s.Elapsed.Done = s.Elapsed.TransferDone
	s.Elapsed.Total = s.Elapsed.Done.Sub(s.Elapsed.Start)

	if s.Response != nil {
		s.StatusCode = s.Response.StatusCode
		s.Data, s.ReadError = ioutil.ReadAll(s.Response.Body)
		s.FreeError = s.Response.Body.Close()
	} else {
		if s.Error == nil {
			s.Error = errors.New("response is nil")
		}
	}

	return s
}
