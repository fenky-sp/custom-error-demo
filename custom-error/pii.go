package customerror

var (
	// map of field name marked as PII
	// key must be in lowercase
	piiFieldNameMap = map[string]bool{
		"phoneno": true,
	}
)
