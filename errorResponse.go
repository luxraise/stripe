package stripe

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	// The document URL explaining the error type
	DocumentURL string `json:"doc_url"`
	// The type of error returned. One of api_connection_error, api_error, authentication_error, card_error, idempotency_error, invalid_request_error, or rate_limit_error
	Type string `json:"type"`
	// For some errors that could be handled programmatically, a short string indicating the error code reported.
	Code string `json:"code"`
	// For card errors resulting from a card issuer decline, a short string indicating the card issuer’s reason for the decline if they provide one.
	DeclineCode string `json:"decline_code"`
	// A human-readable message providing more details about the error. For card errors, these messages can be shown to your users.
	Message string `json:"message"`
	// If the error is parameter-specific, the parameter related to the error. For example, you can use this to display a message near the correct form field.
	Param string `json:"param"`
}

func (e *Error) Error() string {
	return e.Message
}
