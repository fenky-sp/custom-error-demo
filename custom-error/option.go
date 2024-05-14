package customerror

import (
	"context"
	"fmt"
)

type Option func(*ErrorMetadata)

func WithContext(input context.Context) Option {
	return func(m *ErrorMetadata) {
		m.Context = input
	}
}

func WithErrorType(input string) Option {
	return func(m *ErrorMetadata) {
		m.ErrorType = input
	}
}

func WithPIC(input string) Option {
	return func(m *ErrorMetadata) {
		m.PIC = input
	}
}

func WithRequest(input interface{}) Option {
	return func(t *ErrorMetadata) {
		t.Request = fmt.Sprintf("%+v", input)
	}
}

func WithResponse(input interface{}) Option {
	return func(t *ErrorMetadata) {
		t.Response = fmt.Sprintf("%+v", input)
	}
}

func (m *ErrorMetadata) WithRequest(input interface{}) *ErrorMetadata {
	m.Request = fmt.Sprintf("%+v", input)
	return m
}

func (m *ErrorMetadata) WithResponse(input interface{}) *ErrorMetadata {
	m.Response = fmt.Sprintf("%+v", input)
	return m
}
