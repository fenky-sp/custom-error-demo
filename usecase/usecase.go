package usecase

import (
	"context"
	"errors"

	"github.com/fenky-sp/custom-error-demo/constant"
	ctxHlp "github.com/fenky-sp/custom-error-demo/helper/context"
	funcHlp "github.com/fenky-sp/custom-error-demo/helper/function"
	"github.com/fenky-sp/custom-error-demo/model"
	"github.com/fenky-sp/custom-error-demo/repository"
)

func UsecaseFunc(ctx context.Context, input model.UsecaseInput) (output model.UsecaseOutput, err error) {
	ctx = ctxHlp.SetContext(ctx, ctxHlp.TraceFunction(funcHlp.GetFunctionName(UsecaseFunc)))

	_, err = repository.RepositoryFunc(ctx, model.RepositoryInput{
		PhoneNo:         input.Phone,
		RequestTimeUnix: input.RequestTimeUnix,
	})
	err = errors.Join(constant.UsecaseErr1, err)

	return
}
