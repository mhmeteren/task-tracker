package util

type ValidationError struct {
	Errors map[string]string
}

func (v *ValidationError) Error() string {
	return "validation failed"
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type ConflictError struct {
	Message string
}

func (e *ConflictError) Error() string {
	return e.Message
}

type AuthError struct {
	Message string
}

func (e *AuthError) Error() string {
	return e.Message
}
