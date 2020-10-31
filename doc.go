// Package appenv build option handlers to manage application options.
//
// Usage index
//
// Using this library, follow each step below.
//
// 1. Define options.
//
// 2. Create a generator.
//
// 3. Call that generator (`go run <main>`).
//
// Then you can access or configure options with generated function (GetAccessor, GetConfig or GetAppenv).
// Each step is explained below.
//
// Define options
//
// You can define options with just creating structs that
// implement `appenv/types.Value` that
// have `Value()`, `Default()`, `MarshalText` and `UnmarshalText`.
//
// If you want to define a primitive (like `string`) option,
// you can embed `types.XXXValue` like below.
//
// Sample:
//
//   type Token struct {
//   	types.StringValue
//   }
//
// Create a generator
//
// `appenv` does NOT provide any tools like `xxx-generate`.
// Creating a generator, calling it, you can get the handlers.
// To generate, you may call `appenv.Generate` function with options.
//
// Options are built by `appenv.Opt` from `Value`s that you defined in above.
// `appenv.Opt` receives `store` options that specify where the option will be stored to or loaded from.
// Now `appenv` supports stores: YAML file, keyring or environment variables.
//
// Each option can store to / be loaded from multiple `store`.
//
// Generation code should be tagged like `// +build generate`.
// The tag may prevent the generator from being unintendedly built.
// You can write the go:generate directive to call it from `go generate`.
//
//   //+build generate
//
//   //go:generate go run -tags generate .
//
//   func main() {
//   	appenv.Generate(...)
//   }
//
// See: https://pkg.go.dev/github.com/kyoh86/appenv#example-Generate
//
// Access options with generated function
//
// See: Example (GetAccess)
//
// Configure options with generated function
//
// See: Example (GetConfig)
package appenv
