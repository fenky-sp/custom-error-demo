package model

type HandlerInput struct {
	Phone string `ctxerr:"pii"`
}
type HandlerOutput struct{}

type UsecaseInput struct {
	Phone           string
	RequestTimeUnix int64
}
type UsecaseOutput struct{}

type RepositoryInput struct {
	PhoneNo         string
	RequestTimeUnix int64
}
type RepositoryOutput struct{}
