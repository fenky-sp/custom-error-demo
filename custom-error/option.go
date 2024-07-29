package customerror

import (
	"context"
)

type Option func(*errorMetadata)

func WithContext(input context.Context) Option {
	return func(m *errorMetadata) {
		m.Context = input
	}
}

func WithErrorType(input string) Option {
	return func(m *errorMetadata) {
		m.ErrorType = input
	}
}

func WithPIC(input string) Option {
	return func(m *errorMetadata) {
		m.PIC = input
	}
}

func WithRequest(input interface{}) Option {
	return func(t *errorMetadata) {
		t.Request = convertContextualErrorDataToString(input)
	}
}

func WithResponse(input interface{}) Option {
	return func(t *errorMetadata) {
		t.Response = convertContextualErrorDataToString(input)
	}
}

func (m *errorMetadata) WithRequest(input interface{}) *errorMetadata {
	m.Request = convertContextualErrorDataToString(input)
	return m
}

func (m *errorMetadata) WithResponse(input interface{}) *errorMetadata {
	m.Response = convertContextualErrorDataToString(input)
	return m
}
