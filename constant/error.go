package constant

import "errors"

var (
	RepositoryErr1 = errors.New("expected repository error 1")
	RepositoryErr2 = errors.New("expected repository error 2")
	RepositoryErr3 = errors.New("expected repository error 3")

	UsecaseErr1 = errors.New("expected usecase error 1")

	HandlerErr1 = errors.New("expected handler error 1")
)
