package http

const (
	InternalServerError = 500 // Server Error (HTTP/1.0 - RFC 1945)
	NotImplemented      = 501 // Not Implemented (HTTP/1.0 - RFC 1945)
	BadGateway          = 502 // Bad Gateway (HTTP/1.0 - RFC 1945)
	ServiceUnavailable  = 503 // Service Unavailable (HTTP/1.0 - RFC 1945)
	GatewayTimeout      = 504 // Gateway Timeout (HTTP/1.1 - RFC 2616)
	VersionNotSupported = 505 // HTTP Version Not Supported (HTTP/1.1 - RFC 2616)
	InsufficientStorage = 507 // Insufficient Storage (WebDAV - RFC 2518)
)
