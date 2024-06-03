# custom-error-demo

## description
this repository contains demo for contextual error implementation

## utility
there are two utilities used in this demo
- context modifier: modify context to store the trace of process
- error wrapper: custom error interface that can store metadata of: context, pic, error type, request, response and so on

## flow
flow of process: `main.go` > `handler` > `usecase` > `repository`

## error initiation
error is thrown and wrapped in `repository` layer
```go
package repository

import (
	"context"
	"errors"

	"github.com/fenky-sp/custom-error-demo/constant"
	customerror "github.com/fenky-sp/custom-error-demo/custom-error"
	ctxHlp "github.com/fenky-sp/custom-error-demo/helper/context"
	funcHlp "github.com/fenky-sp/custom-error-demo/helper/function"
	"github.com/fenky-sp/custom-error-demo/model"
)

func RepositoryFunc(ctx context.Context, input model.RepositoryInput) (output model.RepositoryOutput, err error) {
	ctx = ctxHlp.SetContext(ctx, ctxHlp.TraceFunction(funcHlp.GetFunctionName(RepositoryFunc)))

	// opt 1
	err = customerror.WrapError(ctx,
		errors.Join(errors.New("expected repository error 1"), errors.New("expected repository error 2")),
		constant.ServiceMesocarp,
		customerror.ErrorTypeDB,
		customerror.OptionalParameter{
			Request: input,
		},
	)

	// // opt 2
	// err = customerror.
	// 	Init(ctx,
	// 		errors.Join(errors.New("expected repository error 1"), errors.New("expected repository error 2")),
	// 		constant.ServiceMesocarp,
	// 		customerror.ErrorTypeDB,
	// 	).
	// 	WithOptions(
	// 		customerror.WithRequest(input),
	// 	)

	// // opt 3 - fluent interface design pattern
	// err = customerror.
	// 	Init(ctx,
	// 		errors.Join(errors.New("expected repository error 1"), errors.New("expected repository error 2")),
	// 		constant.ServiceMesocarp,
	// 		customerror.ErrorTypeDB,
	// 	).
	// 	WithRequest(input).
	// 	WithResponse("dummy response").
	// 	Create()

	return
}
```

### options
there are three options for custom error initiation. we need to determine which one before proceeding to implementation.

#### option 1
custom error is initiated with mandatory and optional metadata at once.

I prefer this option because I have a concern engineers would forget to add metadata like request and response if we proceed with option 2 and 3

#### option 2
custom error is initiated with mandatory metadata only. optional metadata is added in bulk with method `WithOptions()`.

#### option 3
custom error is initiated with mandatory metadata only. optional metadata could be added with method chaining, adopting fluent interface design pattern. `Create()` function have to be called at the end in order to finish the custom error creation.

## output
error is logged at `main.go`
```go
package main

import (
	"context"
	"fmt"

	customerror "github.com/fenky-sp/custom-error-demo/custom-error"
	handlerImpl "github.com/fenky-sp/custom-error-demo/handler/impl"
	"github.com/fenky-sp/custom-error-demo/model"
)

func main() {
	handler := handlerImpl.New()
	_, err := handler.HandlerFunc(context.Background(), model.HandlerInput{})
	if err != nil {
		fmt.Printf("DEBUG custom err (1): %+v\n", err.Error())
		fmt.Println()
		fmt.Printf("DEBUG custom err (2): %+v\n", err)
		fmt.Println()
		fmt.Printf("DEBUG just err: %+v\n", customerror.GetError(err))
	}
}
```

### log
```JSON
DEBUG custom err (1): {"error":"expected repository error 1\nexpected repository error 2","func":"RepositoryFunc","lines":["/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/repository/repository.go:18"],"pic":"mesocarp","request":"{RequestTimeUnix:1717396654}","response":"\u003cnil\u003e","trace":"fenky-sp/custom-error-demo/handler/impl/handler.go-(HandlerFunc)#github.com/fenky-sp/custom-error-demo/usecase/usecase.go-(UsecaseFunc)#github.com/fenky-sp/custom-error-demo/repository/repository.go-(RepositoryFunc)","type":"db"}

DEBUG custom err (2): {"error":"expected repository error 1\nexpected repository error 2","func":"RepositoryFunc","lines":["/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/repository/repository.go:18"],"pic":"mesocarp","request":"{RequestTimeUnix:1717396654}","response":"\u003cnil\u003e","trace":"fenky-sp/custom-error-demo/handler/impl/handler.go-(HandlerFunc)#github.com/fenky-sp/custom-error-demo/usecase/usecase.go-(UsecaseFunc)#github.com/fenky-sp/custom-error-demo/repository/repository.go-(RepositoryFunc)","type":"db"}

DEBUG just err: expected repository error 1
expected repository error 2
```
