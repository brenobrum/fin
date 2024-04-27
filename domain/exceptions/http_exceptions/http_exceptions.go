package http_exceptions

const (
	// Client errors (4xx)
	BAD_REQUEST                     = "bad request"
	UNAUTHORIZED                    = "unauthorized"
	FORBIDDEN                       = "forbidden"
	NOT_FOUND                       = "not found"
	METHOD_NOT_ALLOWED              = "method not allowed"
	NOT_ACCEPTABLE                  = "not acceptable"
	PROXY_AUTH_REQUIRED             = "proxy authentication required"
	REQUEST_TIMEOUT                 = "request timeout"
	CONFLICT                        = "conflict"
	GONE                            = "gone"
	LENGTH_REQUIRED                 = "length required"
	PRECONDITION_FAILED             = "precondition failed"
	REQUEST_ENTITY_TOO_LARGE        = "request entity too large"
	REQUEST_URI_TOO_LONG            = "request URI too long"
	UNSUPPORTED_MEDIA_TYPE          = "unsupported media type"
	REQUESTED_RANGE_NOT_SATISFIABLE = "requested range not satisfiable"
	EXPECTATION_FAILED              = "expectation failed"
	UNPROCESSABLE_ENTITY            = "unprocessable entity"
	LOCKED                          = "locked"
	FAILED_DEPENDENCY               = "failed dependency"
	UPGRADE_REQUIRED                = "upgrade required"
	PRECONDITION_REQUIRED           = "precondition required"
	TOO_MANY_REQUESTS               = "too many requests"
	REQUEST_HEADER_FIELDS_TOO_LARGE = "request header fields too large"

	// Server errors (5xx)
	INTERNAL_SERVER_ERROR      = "internal server error"
	NOT_IMPLEMENTED            = "not implemented"
	BAD_GATEWAY                = "bad gateway"
	SERVICE_UNAVAILABLE        = "service unavailable"
	GATEWAY_TIMEOUT            = "gateway timeout"
	HTTP_VERSION_NOT_SUPPORTED = "http version not supported"
)

const (
	// Client errors (4xx)
	BAD_REQUEST_CODE                     = "BAD_REQUEST"
	UNAUTHORIZED_CODE                    = "UNAUTHORIZED"
	FORBIDDEN_CODE                       = "FORBIDDEN"
	NOT_FOUND_CODE                       = "NOT_FOUND"
	METHOD_NOT_ALLOWED_CODE              = "METHOD_NOT_ALLOWED"
	NOT_ACCEPTABLE_CODE                  = "NOT_ACCEPTABLE"
	PROXY_AUTH_REQUIRED_CODE             = "PROXY_AUTH_REQUIRED"
	REQUEST_TIMEOUT_CODE                 = "REQUEST_TIMEOUT"
	CONFLICT_CODE                        = "CONFLICT"
	GONE_CODE                            = "GONE"
	LENGTH_REQUIRED_CODE                 = "LENGTH_REQUIRED"
	PRECONDITION_FAILED_CODE             = "PRECONDITION_FAILED"
	REQUEST_ENTITY_TOO_LARGE_CODE        = "REQUEST_ENTITY_TOO_LARGE"
	REQUEST_URI_TOO_LONG_CODE            = "REQUEST_URI_TOO_LONG"
	UNSUPPORTED_MEDIA_TYPE_CODE          = "UNSUPPORTED_MEDIA_TYPE"
	REQUESTED_RANGE_NOT_SATISFIABLE_CODE = "REQUESTED_RANGE_NOT_SATISFIABLE"
	EXPECTATION_FAILED_CODE              = "EXPECTATION_FAILED"
	UNPROCESSABLE_ENTITY_CODE            = "UNPROCESSABLE_ENTITY"
	LOCKED_CODE                          = "LOCKED"
	FAILED_DEPENDENCY_CODE               = "FAILED_DEPENDENCY"
	UPGRADE_REQUIRED_CODE                = "UPGRADE_REQUIRED"
	PRECONDITION_REQUIRED_CODE           = "PRECONDITION_REQUIRED"
	TOO_MANY_REQUESTS_CODE               = "TOO_MANY_REQUESTS"
	REQUEST_HEADER_FIELDS_TOO_LARGE_CODE = "REQUEST_HEADER_FIELDS_TOO_LARGE"

	// Server errors (5xx)
	INTERNAL_SERVER_ERROR_CODE      = "INTERNAL_SERVER_ERROR"
	NOT_IMPLEMENTED_CODE            = "NOT_IMPLEMENTED"
	BAD_GATEWAY_CODE                = "BAD_GATEWAY"
	SERVICE_UNAVAILABLE_CODE        = "SERVICE_UNAVAILABLE"
	GATEWAY_TIMEOUT_CODE            = "GATEWAY_TIMEOUT"
	HTTP_VERSION_NOT_SUPPORTED_CODE = "HTTP_VERSION_NOT_SUPPORTED"
)
