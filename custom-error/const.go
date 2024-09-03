package customerror

const (
	ctxErrTagKey      = "ctxerr"
	ctxErrTagValuePii = "pii"
)

const (
	ErrorTypeDB   = "db"
	ErrorTypeHTTP = "http"

	ErrorTypeAuthorization = "authorization"
	ErrorTypeParsing       = "parsing"
	ErrorTypeConversion    = "conversion"
	ErrorTypeValidation    = "validation"
)
