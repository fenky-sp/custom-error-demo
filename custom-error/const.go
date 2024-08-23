package customerror

const (
	CtxErrTagKey      = "ctxerr"
	CtxErrTagValuePii = "pii"
)

const (
	ErrorTypeDB   = "db"
	ErrorTypeHTTP = "http"

	ErrorTypeAuthorization = "authorization"
	ErrorTypeParsing       = "parsing"
	ErrorTypeConversion    = "conversion"
	ErrorTypeValidation    = "validation"
)
