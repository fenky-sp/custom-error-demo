package customerror

import (
	"context"
	"fmt"
)

type Option func(*Metadata)

func WithContext(input context.Context) Option {
	return func(m *Metadata) {
		m.Context = input
	}
}

func WithErrorType(input string) Option {
	return func(m *Metadata) {
		m.ErrorType = input
	}
}

func WithPIC(input string) Option {
	return func(m *Metadata) {
		m.PIC = input
	}
}

func WithRequest(input interface{}) Option {
	return func(t *Metadata) {
		t.Request = fmt.Sprintf("%+v", input)
	}
}

func WithResponse(input interface{}) Option {
	return func(t *Metadata) {
		t.Response = fmt.Sprintf("%+v", input)
	}
}
