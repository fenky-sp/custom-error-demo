package helper

import (
	"context"
	"fmt"
	"strings"

	rtHlp "github.com/fenky-sp/custom-error-demo/helper/runtime"
)

type (
	TraceFunction string
	Function      string
	Trace         string
	contextKey    string
)

const (
	KeyFunction contextKey = "func"
	KeyTrace    contextKey = "trace"
)

// TraceFunction.SetContext
// set TraceFunction context value
func (tf TraceFunction) SetContext(ctx context.Context) context.Context {
	var (
		function string = string(tf)
	)

	// get invoker information
	invoker := rtHlp.GetInvokerInformation(2)
	if !invoker.OK {
		return SetContext(ctx, Trace(function), Function(function))
	}

	// get function name if empty
	if function == "" && invoker.Fn != nil {
		detailsSplit := strings.Split(invoker.Fn.Name(), ".")
		if len(detailsSplit) != 0 {
			function = detailsSplit[len(detailsSplit)-1]
		}
	}

	return SetContext(ctx, Trace(fmt.Sprintf("%s-(%s)", invoker.File, function)), Function(function))
}

// TraceFunction.GetContext
// get TraceFunction context value
func (tf TraceFunction) GetContext(ctx context.Context) string {
	return string(tf)
}

func DefaultTraceFunctionOption() Option {
	return TraceFunction("")
}

// Function.SetContext
// set Function context value
func (f Function) SetContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, KeyFunction, string(f))
}

// Function.GetContext
// get Function context value
func (f Function) GetContext(ctx context.Context) string {
	var value string
	if val, ok := ctx.Value(KeyFunction).(string); ok {
		value = val
	}
	return value
}

// Trace.SetContext
// set Trace context value
func (t Trace) SetContext(ctx context.Context) context.Context {
	var traceString string
	if _, ok := ctx.Value(KeyTrace).(string); ok {
		traceString = fmt.Sprintf("%s#%s", ctx.Value(KeyTrace).(string), string(t))
	} else {
		traceString = string(t)
	}
	temp := context.WithValue(ctx, KeyTrace, traceString)
	return temp
}

// Trace.GetContext
// get Trace context value
func (t Trace) GetContext(ctx context.Context) string {
	var value string
	if val, ok := ctx.Value(KeyTrace).(string); ok {
		value = val
	}
	return value
}

type Option interface {
	SetContext(ctx context.Context) context.Context
	GetContext(ctx context.Context) string
}

// set context values from multiple options into context
func SetContext(ctx context.Context, opts ...Option) context.Context {
	for _, opt := range opts {
		ctx = opt.SetContext(ctx)
	}
	return ctx
}
