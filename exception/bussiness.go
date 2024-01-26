package exception

// NewNotFound todo
// a 的作用 "Resource not found: %s", "exampleID"
func NewNotFound(format string, a ...interface{}) *APIException {
	return NewAPIException("", NotFound, codeReason(NotFound), format, a...)
}

// NewBadRequest todo
// a 的作用 "Resource not found: %s", "exampleID"
func NewBadRequest(format string, a ...interface{}) *APIException {
	return NewAPIException("", BadRequest, codeReason(BadRequest), format, a...)
}

// NewWebCookisNotFoundnRequest todo
// a 的作用 "Resource not found: %s", "exampleID"
func NewWebCookisNotFoundnRequest(format string, a ...interface{}) *APIException {
	return NewAPIException("", WebCookisNotFound, codeReason(WebCookisNotFound), format, a...)
}
