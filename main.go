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
	fmt.Printf("DEBUG custom err: %+v\n", err.Error())
	fmt.Println()
	fmt.Printf("DEBUG err: %+v\n", customerror.GetError(err))
}
