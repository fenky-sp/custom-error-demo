package context

type contextKey string

const (
	KeyFunction contextKey = "func"
	KeyTrace    contextKey = "trace"
)
