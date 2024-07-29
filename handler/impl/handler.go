package impl

import (
	"context"
	"errors"
	"time"

	"github.com/fenky-sp/custom-error-demo/constant"
	customerror "github.com/fenky-sp/custom-error-demo/custom-error"
	ctxHlp "github.com/fenky-sp/custom-error-demo/helper/context"
	"github.com/fenky-sp/custom-error-demo/model"
	"github.com/fenky-sp/custom-error-demo/usecase"
)

func (h *handler) HandlerFunc(ctx context.Context, input model.HandlerInput) (output model.HandlerOutput, err error) {
	ctx = ctxHlp.SetContext(ctx, ctxHlp.DefaultTraceFunctionOption())

	_, err = usecase.UsecaseFunc(ctx, model.UsecaseInput{
		Phone:           input.Phone,
		RequestTimeUnix: time.Now().Unix(),
	})
	err = errors.Join(constant.HandlerErr1,
		customerror.WrapError(ctx,
			err,
			constant.ServiceMesocarp,
			customerror.ErrorTypeValidation,
			customerror.OptionalParameter{
				Request: input,
			},
		))

	return
}
