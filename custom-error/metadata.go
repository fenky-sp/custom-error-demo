package customerror

import (
	"context"
	"encoding/json"
	"fmt"

	ctxHlp "github.com/fenky-sp/custom-error-demo/helper/context"
	rtHlp "github.com/fenky-sp/custom-error-demo/helper/runtime"
)

type metadata struct {
	ctx     context.Context
	err     error
	errType string
	lines   []string
	pic     string // internal: service name e.g. mesocarp and so on; external: service name e.g. doku and so on.
	req     string
	res     string
}

func initialize(rootErr error) *metadata {
	var (
		err = &metadata{}
	)

	if _, ok := rootErr.(*metadata); !ok {
		if rootErr != nil {
			err.err = rootErr
		}
	} else {
		err = rootErr.(*metadata)
	}

	invoker := rtHlp.GetInvokerInformation(2)
	if !invoker.OK {
		return err
	}

	err.lines = append(err.lines, fmt.Sprintf("%v:%v", invoker.File, invoker.Line))

	return err
}

func (md *metadata) setMetadata(setters ...metadataSetter) CustomError {
	for _, setter := range setters {
		setter(md)
	}
	return md
}

func (md *metadata) Error() string {
	data := make(map[string]interface{})

	if md.ctx != nil {
		data["trace"] = md.ctx.Value(ctxHlp.KeyTrace)
		data["func"] = md.ctx.Value(ctxHlp.KeyFunction)
	}

	data["error"] = md.err.Error()

	if md.errType != "" {
		data["type"] = md.errType
	}

	if len(md.lines) != 0 {
		data["lines"] = md.lines
	}

	if md.pic != "" {
		data["pic"] = md.pic
	}

	if md.req != "" && md.req != "null" {
		data["request"] = md.req
	}

	if md.res != "" && md.res != "null" {
		data["response"] = md.res
	}

	result, _ := json.Marshal(data)

	return string(result)
}

func (md *metadata) Unwrap() error {
	return GetStandardError(md.err)
}
