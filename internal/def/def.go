package def

import (
	"github.com/kyoh86/appenv/types"
)

// Token is a string option
type Token struct {
	types.StringPropertyBase
}

// HostName is a string option with default value
type HostName struct {
	types.StringPropertyBase
}

// Default is the default value for host-name option
func (*HostName) Default() interface{} {
	return "kyoh86.dev"
}

// DryRun is a boolean option
type DryRun struct {
	types.BoolPropertyBase
}
