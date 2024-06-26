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
