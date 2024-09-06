package customerror

import (
	"context"
	"errors"
)

type CustomError interface {
	error
	Unwrap() error
	WithErrorType(errorType string) CustomError
	WithPIC(pic string) CustomError
	WithRequest(request any) CustomError
	WithResponse(response any) CustomError
}

func Wrap(
	ctx context.Context,
	rootErr error,
	metadataSetters ...metadataSetter,
) CustomError {
	if rootErr == nil {
		return nil
	}

	var ms []metadataSetter
	ms = append(ms, contextMetadataSetter(ctx))
	ms = append(ms, metadataSetters...)

	return initialize(rootErr).setMetadata(ms...)
}

func Create(
	ctx context.Context,
	rootErr error,
) CustomError {
	if rootErr == nil {
		return nil
	}

	var ms []metadataSetter
	ms = append(ms, contextMetadataSetter(ctx))

	return initialize(rootErr).setMetadata(ms...)
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
