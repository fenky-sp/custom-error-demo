package main

import (
	"context"
	"fmt"

	handlerImpl "github.com/fenky-sp/custom-error-demo/handler/impl"
	"github.com/fenky-sp/custom-error-demo/model"
)

func main() {
	handler := handlerImpl.New()
	_, err := handler.HandlerFunc(context.Background(), model.HandlerInput{})
	fmt.Printf("DEBUG err: %+v\n", err.Error())
}
