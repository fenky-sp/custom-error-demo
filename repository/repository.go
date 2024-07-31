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

	// opt 1
	err = customerror.WrapError(ctx,
		errors.Join(constant.RepositoryErr1, constant.RepositoryErr2),
		constant.ServiceMesocarp,
		customerror.ErrorTypeDB,
		customerror.OptionalParameter{
			Request: input,
		},
	)

	// // opt 2 - fluent interface design pattern
	// err = customerror.
	// 	Create(ctx,
	// 		errors.Join(constant.RepositoryErr1, constant.RepositoryErr2),
	// 		constant.ServiceMesocarp,
	// 		customerror.ErrorTypeDB,
	// 	).
	// 	WithRequest(input).
	// 	WithResponse("dummy response")

	return
}
