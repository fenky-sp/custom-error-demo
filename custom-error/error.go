package customerror

import "context"

type CustomError interface {
	error
	WithOptions(...Option) CustomError
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

	return create(err).WithOptions(
		WithContext(ctx),
		WithErrorType(errorType),
		WithPIC(pic),
		WithRequest(optional.Request),
		WithResponse(optional.Response),
	)
}
