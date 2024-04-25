package customerror

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	ctxHlp "github.com/fenky-sp/custom-error-demo/helper/context"
)

type Metadata struct {
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
		err = &Metadata{}
	)

	if _, ok := rootErr.(*Metadata); !ok {
		if rootErr != nil {
			err.Err = rootErr
		}
	} else {
		err = rootErr.(*Metadata)
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return err
	}

	err.Lines = append(err.Lines, fmt.Sprintf("%v:%v", file, line))

	return err
}

func (m *Metadata) Error() string {
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
		// TODO fenky check
		data["trace"] = m.Context.Value(ctxHlp.KeyTrace)
		data["func_name"] = m.Context.Value(ctxHlp.KeyFuncName)
	}

	if len(m.Lines) != 0 {
		data["lines"] = m.Lines
	}

	result, _ := json.Marshal(data)

	return string(result)
}

func (m *Metadata) WithOption(opts ...Option) CustomError {
	for _, opt := range opts {
		opt(m)
	}
	return m
}
