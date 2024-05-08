# custom-error-demo

## output
```
DEBUG custom err (1): {"error":"expected repository error 1\nexpected repository error 2","func":"RepositoryFunc","lines":["/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/repository/repository.go:17"],"pic":"mesocarp","request":"{RequestTimeUnix:1715137702}","response":"\u003cnil\u003e","trace":"fenky-sp/custom-error-demo/handler/impl/handler.go-(HandlerFunc)#github.com/fenky-sp/custom-error-demo/usecase/usecase.go-(UsecaseFunc)#github.com/fenky-sp/custom-error-demo/repository/repository.go-(RepositoryFunc)","type":"db"}

DEBUG custom err (2): {"error":"expected repository error 1\nexpected repository error 2","func":"RepositoryFunc","lines":["/Users/fenky/go/src/github.com/fenky-sp/custom-error-demo/repository/repository.go:17"],"pic":"mesocarp","request":"{RequestTimeUnix:1715137702}","response":"\u003cnil\u003e","trace":"fenky-sp/custom-error-demo/handler/impl/handler.go-(HandlerFunc)#github.com/fenky-sp/custom-error-demo/usecase/usecase.go-(UsecaseFunc)#github.com/fenky-sp/custom-error-demo/repository/repository.go-(RepositoryFunc)","type":"db"}

DEBUG just err: expected repository error 1
```
