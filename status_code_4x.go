package http

const (
	BadRequest                   = 400 // Bad Request (HTTP/1.1 - RFC 2616)
	Unauthorized                 = 401 // Unauthorized (HTTP/1.0 - RFC 1945)
	PaymentRequired              = 402 // Payment Required (HTTP/1.1 - RFC 2616)
	Forbidden                    = 403 // Forbidden (HTTP/1.0 - RFC 1945)
	NotFound                     = 404 // Not Found (HTTP/1.0 - RFC 1945)
	MethodNotAllowed             = 405 // Method Not Allowed (HTTP/1.1 - RFC 2616)
	NotAcceptable                = 406 // Not Acceptable (HTTP/1.1 - RFC 2616)
	ProxyAuthenticationRequired  = 407 // Proxy Authentication Required (HTTP/1.1 - RFC 2616)
	RequestTimeout               = 408 // Request Timeout (HTTP/1.1 - RFC 2616)
	Conflict                     = 409 // Conflict (HTTP/1.1 - RFC 2616)
	Gone                         = 410 // Gone (HTTP/1.1 - RFC 2616)
	LengthRequired               = 411 // Length Required (HTTP/1.1 - RFC 2616)
	PreconditionFailed           = 412 // Precondition Failed (HTTP/1.1 - RFC 2616)
	RequestTooLong               = 413 // Request Entity Too Large (HTTP/1.1 - RFC 2616)
	RequestUriTooLong            = 414 // Request-URI Too Long (HTTP/1.1 - RFC 2616)
	UnsupportedMediaType         = 415 // Unsupported Media Type (HTTP/1.1 - RFC 2616)
	RequestedRangeNotSatisfiable = 416 // Requested Range Not Satisfiable (HTTP/1.1 - RFC 2616)
	ExpectationFailed            = 417 // Expectation Failed (HTTP/1.1 - RFC 2616)
	InsufficientSpaceOnResource  = 419 // Static constant for a 419 error.
	MethodFailure                = 420 // Static constant for a 420 error.
	UnprocessableEntity          = 422 // Unprocessable Entity (WebDAV - RFC 2518)
	Locked                       = 423 // Locked (WebDAV - RFC 2518)
	FailedDependency             = 424 // Failed Dependency (WebDAV - RFC 2518)
)
