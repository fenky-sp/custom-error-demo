package context

import (
	"context"
	"fmt"
	"strings"

	rtHlp "github.com/fenky-sp/custom-error-demo/helper/runtime"
)

type KeyValue interface {
	SetContext(ctx context.Context) context.Context
	GetContext(ctx context.Context) string
}

// set context values from multiple options into context
func SetContext(ctx context.Context, kvs ...KeyValue) context.Context {
	for _, kv := range kvs {
		ctx = kv.SetContext(ctx)
	}
	return ctx
}

func DefaultTraceFunction() TraceFunction {
	return TraceFunction("")
}

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
