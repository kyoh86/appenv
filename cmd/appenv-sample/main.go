// +build sample

package main

import (
	"fmt"

	"github.com/kyoh86/appenv"
)

func main() {
	fmt.Printf("A version of the Package %s is %s\n", "appenv", appenv.Version())
}
