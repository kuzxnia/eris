package exception

type ValidationError struct {
	Message string
}

func (ex ValidationError) Error() string {
	return ex.Message
}

type NotFoundError struct {
	Message string
}

func (ex NotFoundError) Error() string {
	return ex.Message
}

type UnauthorizedError struct {
	Message string
}

func (unauthorizedError UnauthorizedError) Error() string {
	return unauthorizedError.Message
}

type ResourceAlreadyExistsError struct {
	Message string
}

func (ex ResourceAlreadyExistsError) Error() string {
	return ex.Message
}
