package customerror

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	ctxHlp "github.com/fenky-sp/custom-error-demo/helper/context"
	rtHlp "github.com/fenky-sp/custom-error-demo/helper/runtime"
)

type errorMetadata struct {
	Context   context.Context
	Err       error
	ErrorType string
	Lines     []string
	PIC       string // internal: service name e.g. mesocarp and so on; external: service name e.g. doku and so on.
	Request   string
	Response  string
}

func create(rootErr error) CustomError {
	var (
		err = &errorMetadata{}
	)

	if _, ok := rootErr.(*errorMetadata); !ok {
		if rootErr != nil {
			err.Err = rootErr
		}
	} else {
		err = rootErr.(*errorMetadata)
	}

	invoker := rtHlp.GetInvokerInformation(2)
	if !invoker.OK {
		return err
	}

	err.Lines = append(err.Lines, fmt.Sprintf("%v:%v", invoker.File, invoker.Line))

	return err
}

func (m *errorMetadata) WithOptions(opts ...Option) CustomError {
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *errorMetadata) Error() string {
	data := make(map[string]interface{})

	if m.Context != nil {
		data["trace"] = m.Context.Value(ctxHlp.KeyTrace)
		data["func"] = m.Context.Value(ctxHlp.KeyFunction)
	}

	data["error"] = m.Err.Error()

	if m.ErrorType != "" {
		data["type"] = m.ErrorType
	}

	if len(m.Lines) != 0 {
		data["lines"] = m.Lines
	}

	if m.PIC != "" {
		data["pic"] = m.PIC
	}

	if m.Request != "null" {
		data["request"] = m.Request
	}

	if m.Response != "null" {
		data["response"] = m.Response
	}

	result, _ := json.Marshal(data)

	return string(result)
}

func extractError(err error) error {
	switch err.(type) {
	case *errorMetadata:
		if em, ok := err.(*errorMetadata); ok {
			return em.Err
		}
	}

	return err
}

type errorData struct {
	errs []error
}

func (ed *errorData) getErrors(err error) {
	err = extractError(err)

	switch x := err.(type) {

	case interface{ Unwrap() error }:
		err = x.Unwrap()
		if err != nil {
			ed.getErrors(err) // check recursively if error is wrapped
		}

	case interface{ Unwrap() []error }:
		wrappedErrors := x.Unwrap()
		for _, wrappedError := range wrappedErrors {
			ed.getErrors(wrappedError)
		}

	default:
		ed.errs = append(ed.errs, err)

	}
}

// GetStandardError converts custom error to standard error
func GetStandardError(err error) error {
	ed := &errorData{}
	ed.getErrors(err)
	return errors.Join(ed.errs...)
}

// Is checks if error match the target error
func Is(err error, target error) bool {
	return errors.Is(GetStandardError(err), target)
}

func Init(
	ctx context.Context,
	rootErr error,
	pic string,
	errorType string,
) *errorMetadata {
	var (
		err = &errorMetadata{}
	)

	if _, ok := rootErr.(*errorMetadata); !ok {
		if rootErr != nil {
			err.Err = rootErr
		}
	} else {
		err = rootErr.(*errorMetadata)
	}

	invoker := rtHlp.GetInvokerInformation(2)
	if !invoker.OK {
		return err
	}

	err.Lines = append(err.Lines, fmt.Sprintf("%v:%v", invoker.File, invoker.Line))

	return err
}

func (m *errorMetadata) Create() CustomError {
	return m
}
