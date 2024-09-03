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
	}

	// check contextual error with errors.Is()
	fmt.Printf("\n\n\n")
	fmt.Println("DEBUG errors.Is(err, errors.New(\"expected repository error 2\")):", errors.Is(err, errors.New("expected repository error 2"))) // false
	fmt.Println("DEBUG errors.Is(err, constant.UsecaseErr1):", errors.Is(err, constant.RepositoryErr2))                                          // true

	// err = customerror.Create(ctx, errors.New("temp"), constant.ServiceMesocarp, customerror.ErrorTypeValidation).
	// 	WithRequest(map[int]struct {
	// 		Test1 string  `ctxerr:"pii"`
	// 		Test2 int64   `ctxerr:"pii"`
	// 		Test3 uint64  `ctxerr:"pii"`
	// 		Test4 float64 `ctxerr:"pii"`
	// 	}{
	// 		// Test1: "1",
	// 		// Test2: 2,
	// 		// Test3: 3,
	// 		// Test4: 3.3,

	// 		// {
	// 		// 	Test1: "1",
	// 		// 	Test2: 2,
	// 		// 	Test3: 3,
	// 		// 	Test4: 3.3,
	// 		// },
	// 		// {
	// 		// 	Test1: "4",
	// 		// 	Test2: 5,
	// 		// 	Test3: 6,
	// 		// 	Test4: 6.6,
	// 		// },
	// 		// {
	// 		// 	Test1: "7",
	// 		// 	Test2: 8,
	// 		// 	Test3: 9,
	// 		// 	Test4: 9.9,
	// 		// },

	// 		1: {
	// 			Test1: "1",
	// 			Test2: 2,
	// 			Test3: 3,
	// 			Test4: 3.3,
	// 		},
	// 		2: {
	// 			Test1: "4",
	// 			Test2: 5,
	// 			Test3: 6,
	// 			Test4: 6.6,
	// 		},
	// 		3: {
	// 			Test1: "7",
	// 			Test2: 8,
	// 			Test3: 9,
	// 			Test4: 9.9,
	// 		},
	// 	}).
	// 	WithResponse(model.HandlerInput{})
	// fmt.Printf("\n\n\n")
	// fmt.Println("DEBUG err:", err)
}
