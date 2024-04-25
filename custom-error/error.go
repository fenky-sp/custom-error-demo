package customerror

import "context"

type CustomError interface {
	error
	WithOption(...Option) CustomError
}

func WrapError(
	ctx context.Context,
	err error,
	pic string,
	errorType string,
	optional OptionalParameter,
) CustomError {
	if err == nil {
		return nil
	}

	return create(err).WithOption(
		WithContext(ctx),
		WithErrorType(errorType),
		WithPIC(pic),
		WithRequest(optional.Request),
		WithResponse(optional.Response),
	)
}
