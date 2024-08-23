package customerror

import (
	"context"
	"errors"
)

type CustomError interface {
	error
	WithRequest(request any) CustomError
	WithResponse(response any) CustomError
}

func WrapError(
	ctx context.Context,
	rootErr error,
	pic string,
	errorType string,
	optional OptionalParameter,
) CustomError {
	if rootErr == nil {
		return nil
	}

	return initialize(rootErr).setMetadata(
		contextMetadataSetter(ctx),
		errorTypeMetadataSetter(errorType),
		picMetadataSetter(pic),
		requestMetadataSetter(optional.Request),
		responseMetadataSetter(optional.Response),
	)
}

func Create(
	ctx context.Context,
	rootErr error,
	pic string,
	errorType string,
) CustomError {
	if rootErr == nil {
		return nil
	}

	return initialize(rootErr).setMetadata(
		contextMetadataSetter(ctx),
		errorTypeMetadataSetter(errorType),
		picMetadataSetter(pic),
	)
}

// GetStandardError converts custom error to standard error
func GetStandardError(err error) error {
	ed := &errorData{}
	ed.getError(err)
	return errors.Join(ed.errs...)
}

// Is checks if error match the target error
func Is(err error, target error) bool {
	return errors.Is(GetStandardError(err), target)
}
