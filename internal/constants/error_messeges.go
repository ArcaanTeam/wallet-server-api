package constants

type ErrorType uint

const (
	ErrIDParamNotProvided ErrorType = iota
	ErrInvalidIDParam

	ErrAuthUserNotFound
	ErrAuthUnauthorized
	ErrAuthGenerateTokenFailed
)
