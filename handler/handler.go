package handler

import (
	"context"

	"github.com/fenky-sp/custom-error-demo/model"
)

type Handler interface {
	HandlerFunc(ctx context.Context, input model.HandlerInput) (output model.HandlerOutput, err error)
}
