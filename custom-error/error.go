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
	if err == nil {
		return nil
	}
	ed := &errorData{}
	ed.getError(err)
	if len(ed.errs) == 1 {
		return ed.errs[0]
	}
	return errors.Join(ed.errs...)
}

// Is checks if error match the target error
func Is(err error, target error) bool {
	errStd := GetStandardError(err)
	targetStd := GetStandardError(target)
	output := errors.Is(errStd, targetStd) // compare error
	if !output && errStd != nil && targetStd != nil {
		output = errStd.Error() == targetStd.Error() // compare error message
	}
	return output
}
