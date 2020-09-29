package errors

import (
	"errors"
	"fmt"
	"strings"

	pkgErrors "github.com/pkg/errors"
)

// Export origin func of errors
var (
	Unwrap       = errors.Unwrap
	Is           = errors.Is
	As           = errors.As
	Wrap         = pkgErrors.Wrap
	Wrapf        = pkgErrors.Wrapf
	New          = pkgErrors.New
	Cause        = pkgErrors.Cause
	WithMessage  = pkgErrors.WithMessage
	WithMessagef = pkgErrors.WithMessagef
	WithStack    = pkgErrors.WithStack
)

// StackTrace returns stack frames
func StackTrace(e error) []string {
	stacktrace := fmt.Sprintf("%+v\n", e)
	output := strings.Split(stacktrace, "\n")
	return output[:len(output)-1]
}
