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

func ErrorTypeMetadataSetter(input string) metadataSetter {
	return func(md *metadata) {
		md.errType = input
	}
}

func PicMetadataSetter(input string) metadataSetter {
	return func(md *metadata) {
		md.pic = input
	}
}

func RequestMetadataSetter(input any) metadataSetter {
	return func(md *metadata) {
		md.req = convertContextualErrorDataToString(input)
	}
}

func ResponseMetadataSetter(input any) metadataSetter {
	return func(md *metadata) {
		md.res = convertContextualErrorDataToString(input)
	}
}

func (md *metadata) WithErrorType(errorType string) CustomError {
	md.setMetadata(
		ErrorTypeMetadataSetter(errorType),
	)
	return md
}

func (md *metadata) WithPIC(pic string) CustomError {
	md.setMetadata(
		PicMetadataSetter(pic),
	)
	return md
}

func (md *metadata) WithRequest(request any) CustomError {
	md.setMetadata(
		RequestMetadataSetter(request),
	)
	return md
}

func (md *metadata) WithResponse(response any) CustomError {
	md.setMetadata(
		ResponseMetadataSetter(response),
	)
	return md
}
