package customerror

import (
	"context"
)

type metadataSetter func(*metadata)

func contextMetadataSetter(input context.Context) metadataSetter {
	return func(md *metadata) {
		md.ctx = input
	}
}

func errorTypeMetadataSetter(input string) metadataSetter {
	return func(md *metadata) {
		md.errType = input
	}
}

func picMetadataSetter(input string) metadataSetter {
	return func(md *metadata) {
		md.pic = input
	}
}

func requestMetadataSetter(input any) metadataSetter {
	return func(md *metadata) {
		md.req = convertContextualErrorDataToString(input)
	}
}

func responseMetadataSetter(input any) metadataSetter {
	return func(md *metadata) {
		md.res = convertContextualErrorDataToString(input)
	}
}

func (md *metadata) WithRequest(request any) CustomError {
	md.setMetadata(
		requestMetadataSetter(request),
	)
	return md
}

func (md *metadata) WithResponse(response any) CustomError {
	md.setMetadata(
		responseMetadataSetter(response),
	)
	return md
}
