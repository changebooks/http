package http

const (
	Ok                          = 200 // OK (HTTP/1.0 - RFC 1945)
	Created                     = 201 // Created (HTTP/1.0 - RFC 1945)
	Accepted                    = 202 // Accepted (HTTP/1.0 - RFC 1945)
	NonAuthoritativeInformation = 203 // Non Authoritative Information (HTTP/1.1 - RFC 2616)
	NoContent                   = 204 // No Content (HTTP/1.0 - RFC 1945)
	ResetContent                = 205 // Reset Content (HTTP/1.1 - RFC 2616)
	PartialContent              = 206 // Partial Content (HTTP/1.1 - RFC 2616)
	MultiStatus                 = 207 // Multi-Status (WebDAV - RFC 2518) or 207 Partial Update OK (HTTP/1.1 - draft-ietf-http-v11-spec-rev-01?)
)
