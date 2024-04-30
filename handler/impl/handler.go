package impl

import (
	"context"
	"time"

	ctxHlp "github.com/fenky-sp/custom-error-demo/helper/context"
	funcHlp "github.com/fenky-sp/custom-error-demo/helper/function"
	"github.com/fenky-sp/custom-error-demo/model"
	"github.com/fenky-sp/custom-error-demo/usecase"
)

func (h *handler) HandlerFunc(ctx context.Context, input model.HandlerInput) (output model.HandlerOutput, err error) {
	ctx = ctxHlp.SetContext(ctx, ctxHlp.TraceFunction(funcHlp.GetFunctionName(h.HandlerFunc)))

	_, err = usecase.UsecaseFunc(ctx, model.UsecaseInput{
		RequestTimeUnix: time.Now().Unix(),
	})

	return
}
