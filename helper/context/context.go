package helper

import (
	"context"
	"fmt"
	"runtime"
	"strings"
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
		trace    string
	)

	// get caller file path
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		return SetContext(ctx, Trace(function), Function(function))
	}

	// split file path
	filePaths := strings.Split(file, "/")
	if len(filePaths) <= 5 {
		trace = strings.Join(filePaths, "/")
		return SetContext(ctx, Trace(fmt.Sprintf("%s-(%s)", trace, function)), Function(function))
	}

	// trim file path
	filePaths = filePaths[len(filePaths)-5:]
	trace = strings.Join(filePaths, "/")
	return SetContext(ctx, Trace(fmt.Sprintf("%s-(%s)", trace, function)), Function(function))
}

// TraceFunction.GetContext
// get TraceFunction context value
func (tf TraceFunction) GetContext(ctx context.Context) string {
	return string(tf)
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
