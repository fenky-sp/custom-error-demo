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

	err = customerror.WrapError(ctx,
		errors.Join(errors.New("expected repository error 1"), errors.New("expected repository error 2")),
		constant.ServiceMesocarp,
		customerror.ErrorTypeDB,
		customerror.OptionalParameter{
			Request: input,
		},
	)

	return
}
