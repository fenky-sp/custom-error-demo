# custom-error-demo

## Description
This repository contains demo for contextual error implementation

## Utility
There are two utilities used in this demo
- Context modifier: modify context to store the trace of process
- Error wrapper: custom error interface that can store metadata of: context, pic, error type, request, response and so on

## Flow
Flow of process: `main.go` > `handler` > `usecase` > `repository` then back to the beginning

## Error Initiation
Error is thrown and wrapped in `repository` layer
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

	// option 1
	err = customerror.WrapError(ctx,
		errors.Join(constant.RepositoryErr1, constant.RepositoryErr2),
		constant.ServiceMesocarp,
		customerror.ErrorTypeDB,
		customerror.OptionalParameter{
			Request: input,
		},
	)

	// // option 2 - fluent interface design pattern
	// err = customerror.
	// 	Create(ctx,
	// 		errors.Join(constant.RepositoryErr1, constant.RepositoryErr2),
	// 		constant.ServiceMesocarp,
	// 		customerror.ErrorTypeDB,
	// 	).
	// 	WithRequest(input).
	// 	WithResponse("dummy response")

	return
}
```

### Options
There are two ways to init a custom error. We need to determine which one to use before proceeding to implementation.

#### Option 1
Custom error is initiated with mandatory and optional metadata at once.

I prefer this option because I have a concern where Engineers would forget to add metadata like request and response if we proceed with option 2

#### Option 2
Custom error is initiated with mandatory metadata only. Optional metadata could be added with method chaining, adopting fluent interface design pattern.

## Output
Error is logged at `main.go`
```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/fenky-sp/custom-error-demo/constant"
	customerror "github.com/fenky-sp/custom-error-demo/custom-error"
	handlerImpl "github.com/fenky-sp/custom-error-demo/handler/impl"
	"github.com/fenky-sp/custom-error-demo/model"
)

func main() {
	handler := handlerImpl.New()
	_, err := handler.HandlerFunc(context.Background(), model.HandlerInput{
		Phone: "01234567890",
	})
	err = errors.Join(errors.New("expected main error 1"), err)
	if err != nil {
		fmt.Printf("DEBUG custom err (1): %+v\n\n", err.Error())

		fmt.Printf("DEBUG custom err (2): %+v\n\n", err)

		fmt.Printf("DEBUG standard err: %+v\n\n", customerror.GetStandardError(err))

		if customerror.Is(err, constant.RepositoryErr2) {
			fmt.Println("DEBUG error is identified")
		} else {
			fmt.Println("DEBUG error is not identified")
		}
	}
}
```

### Log
```JSON
DEBUG custom err (1): expected main error 1
expected handler error 1
{"error":"expected usecase error 1\n{\"error\":\"expected repository error 1\\nexpected repository error 2\",\"func\":\"RepositoryFunc\",\"lines\":[\"/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/repository/repository.go:18\"],\"pic\":\"mesocarp\",\"request\":\"{\\\"PhoneNo\\\":\\\"***\\\",\\\"RequestTimeUnix\\\":1722425795}\",\"trace\":\"/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/handler/impl/handler.go-(HandlerFunc)#/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/usecase/usecase.go-(UsecaseFunc)#/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/repository/repository.go-(RepositoryFunc)\",\"type\":\"db\"}","func":"HandlerFunc","lines":["/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/handler/impl/handler.go:23"],"pic":"mesocarp","request":"{\"Phone\":\"***\"}","trace":"/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/handler/impl/handler.go-(HandlerFunc)","type":"validation"}

DEBUG custom err (2): expected main error 1
expected handler error 1
{"error":"expected usecase error 1\n{\"error\":\"expected repository error 1\\nexpected repository error 2\",\"func\":\"RepositoryFunc\",\"lines\":[\"/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/repository/repository.go:18\"],\"pic\":\"mesocarp\",\"request\":\"{\\\"PhoneNo\\\":\\\"***\\\",\\\"RequestTimeUnix\\\":1722425795}\",\"trace\":\"/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/handler/impl/handler.go-(HandlerFunc)#/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/usecase/usecase.go-(UsecaseFunc)#/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/repository/repository.go-(RepositoryFunc)\",\"type\":\"db\"}","func":"HandlerFunc","lines":["/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/handler/impl/handler.go:23"],"pic":"mesocarp","request":"{\"Phone\":\"***\"}","trace":"/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/handler/impl/handler.go-(HandlerFunc)","type":"validation"}

DEBUG standard err: expected main error 1
expected handler error 1
expected usecase error 1
expected repository error 1
expected repository error 2

DEBUG error is identified
```
