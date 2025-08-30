package constants

type ErrorType string

const (
	ErrNone = "no error"

	ErrIDParamNotProvided = "id param not provided"
	ErrInvalidIDParam     = "invalid id param"

	ErrAuthUserNotFound        = "user not found"
	ErrAuthUnauthorized        = "unauthorized"
	ErrAuthGenerateTokenFailed = "failed to generate token"
)

func (e ErrorType) Error() string {
	return string(e)
}
