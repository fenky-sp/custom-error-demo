package customerror

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	ctxHlp "github.com/fenky-sp/custom-error-demo/helper/context"
)

type ErrorMetadata struct {
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
		err = &ErrorMetadata{}
	)

	if _, ok := rootErr.(*ErrorMetadata); !ok {
		if rootErr != nil {
			err.Err = rootErr
		}
	} else {
		err = rootErr.(*ErrorMetadata)
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return err
	}

	err.Lines = append(err.Lines, fmt.Sprintf("%v:%v", file, line))

	return err
}

func (m *ErrorMetadata) WithOptions(opts ...Option) CustomError {
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *ErrorMetadata) Error() string {
	data := make(map[string]interface{})

	if m.Request != "" {
		data["request"] = m.Request
	}

	if m.Response != "" {
		data["response"] = m.Response
	}

	if m.PIC != "" {
		data["pic"] = m.PIC
	}

	if m.ErrorType != "" {
		data["type"] = m.ErrorType
	}

	data["error"] = m.Err.Error()

	if m.Context != nil {
		data["trace"] = m.Context.Value(ctxHlp.KeyTrace)
		data["func"] = m.Context.Value(ctxHlp.KeyFunction)
	}

	if len(m.Lines) != 0 {
		data["lines"] = m.Lines
	}

	result, _ := json.Marshal(data)

	return string(result)
}

// GetError gets error only from custom error
func GetError(err error) error {
	switch err.(type) {
	case *ErrorMetadata:
		if em, ok := err.(*ErrorMetadata); ok {
			return em.Err
		}
	}

	return err
}
