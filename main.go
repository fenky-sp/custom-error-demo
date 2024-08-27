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
