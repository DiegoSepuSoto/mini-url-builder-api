package shared

import "fmt"

const (
	DatabaseInsertError    = "DB_INSERT"
	DatabaseNotFoundError  = "DB_NOT_FOUND"
	DatabaseFindError      = "DB_FIND"
	SyncCommunicationError = "SYNC_COMM_ERROR"
)

type EchoErrorResponse struct {
	Message string `json:"message"`
}

type ApplicationError struct {
	ErrorCode        string
	ErrorDescription string
	ErrorOrigin      string
	HTTPStatusCode   int
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("[%s]: %s at %s - sending: %d", e.ErrorCode, e.ErrorDescription, e.ErrorOrigin, e.HTTPStatusCode)
}

func BuildError(httpStatusCode int, errorCode, errorDescription, errorOrigin string) *ApplicationError {
	return &ApplicationError{
		HTTPStatusCode:   httpStatusCode,
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
		ErrorOrigin:      errorOrigin,
	}
}
