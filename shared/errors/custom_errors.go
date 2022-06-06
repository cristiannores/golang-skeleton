package errors

type CustomError struct {
	id         []string
	Kind       ErrorTypes
	Msg        string
	InnerError error
}

type ErrorTypes string

const (
	INPUT_VALIDATION  ErrorTypes = "INPUT_VALIDATION"
	UNEXPECTED_ERROR  ErrorTypes = "UNEXPECTED_ERROR"
	TASK_NOT_FOUND    ErrorTypes = "TASK_NOT_FOUND"
	TASK_NOT_INSERTED ErrorTypes = "TASK_NOT_INSERTED"
)

func New(id []string, Msg string, Kind ErrorTypes) *CustomError {
	return &CustomError{id: id, Msg: Msg, Kind: Kind}
}

func NewWithError(id []string, Msg string, Kind ErrorTypes, InnerError error) *CustomError {
	return &CustomError{id: id, Msg: Msg, Kind: Kind, InnerError: InnerError}
}
func (e *CustomError) Error() string {
	return e.Msg
}
