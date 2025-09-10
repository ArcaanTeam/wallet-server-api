package constants

type ErrorType string

const (
	ErrNone ErrorType = "no error"

	ErrInternal ErrorType = "internal server error"

	ErrIDParamNotProvided ErrorType = "id param not provided"
	ErrInvalidIDParam     ErrorType = "invalid id param"

	ErrAuthUserNotFound        ErrorType = "user not found"
	ErrAuthUnauthorized        ErrorType = "unauthorized"
	ErrAuthGenerateTokenFailed ErrorType = "failed to generate token"

	ErrInvalidUserRoleString ErrorType = "invalid user role string"
)

func (e ErrorType) Error() string {
	return string(e)
}
