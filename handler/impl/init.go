package impl

import (
	h "github.com/fenky-sp/custom-error-demo/handler"
)

type handler struct{}

func New() h.Handler {
	return &handler{}
}
