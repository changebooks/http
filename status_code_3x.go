package http

const (
	MultipleChoices   = 300 // Multiple Choices (HTTP/1.1 - RFC 2616)
	MovedPermanently  = 301 // Moved Permanently (HTTP/1.0 - RFC 1945)
	MovedTemporarily  = 302 // Moved Temporarily (Sometimes Found) (HTTP/1.0 - RFC 1945)
	SeeOther          = 303 // See Other (HTTP/1.1 - RFC 2616)
	NotModified       = 304 // Not Modified (HTTP/1.0 - RFC 1945)
	UseProxy          = 305 // Use Proxy (HTTP/1.1 - RFC 2616)
	TemporaryRedirect = 307 // Temporary Redirect (HTTP/1.1 - RFC 2616)
)
