package runtime

import (
	"runtime"
)

type Invoker struct {
	PC   uintptr
	File string
	Line int
	OK   bool
	Fn   *runtime.Func
}

// GetInvokerInformation
//
// Needs parameter:
// - skip, is the number of stack frames to ascend, with 0 identifying the caller of GetInvokerInformation.
//
// Returns information about function invocations on the calling goroutine's stack:
// - file path.
// - line number.
// - ok, to inform whether it is possible to recover the information or not.
// - a *Func describing the function that contains the given program counter address, or else nil.
func GetInvokerInformation(skip int) (output Invoker) {
	skip += 1 // skip is set +1 to skip `GetInvokerInformation`

	output.PC, output.File, output.Line, output.OK = runtime.Caller(skip)
	if !output.OK {
		return output
	}

	output.Fn = runtime.FuncForPC(output.PC)

	return output
}
