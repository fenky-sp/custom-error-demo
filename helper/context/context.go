package helper

import (
	"context"
	"fmt"
	"runtime"
	"strings"
)

type (
	TraceFunction string
	FuncName      string
	Trace         string
	contextKey    string
)

const (
	KeyFuncName contextKey = "func_name"
	KeyTrace    contextKey = "trace"
)

// TraceFunction.SetContext
// set TraceFunction context value
func (tf TraceFunction) SetContext(ctx context.Context) context.Context {
	// TODO fenky check traceString
	var traceString string

	// get file module trace function
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		traceString = string(tf)
		return SetContext(ctx, FuncName(traceString), Trace(traceString))
	}

	str := strings.Split(file, "/")
	if len(str) < 5 {
		traceString = string(tf)
		return SetContext(ctx, FuncName(traceString), Trace(traceString))
	}

	traceString = fmt.Sprintf("%s.%s.%s.%s.%s", str[len(str)-5], str[len(str)-4], str[len(str)-3], str[len(str)-2], string(tf))
	return SetContext(ctx, FuncName(traceString), Trace(traceString))
}

// get TraceFunction context value
func (tf TraceFunction) GetContext(ctx context.Context) string {
	return string(tf)
}

// set FuncName context value
func (f FuncName) SetContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, KeyFuncName, string(f))
}

// get FuncName context value
func (f FuncName) GetContext(ctx context.Context) string {
	var value string
	if val, ok := ctx.Value(KeyFuncName).(string); ok {
		value = val
	}
	return value
}

// set Trace context value
func (t Trace) SetContext(ctx context.Context) context.Context {
	var traceString string
	if _, ok := ctx.Value(KeyTrace).(string); ok {
		traceString = fmt.Sprintf("%s-%s", ctx.Value(KeyTrace).(string), string(t))
	} else {
		traceString = string(t)
	}
	temp := context.WithValue(ctx, KeyTrace, traceString)
	return temp
}

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
func SetContext(ctx context.Context, opt ...Option) context.Context {
	for _, o := range opt {
		ctx = o.SetContext(ctx)
	}
	return ctx
}
