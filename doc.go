// Package appenv build a option handlers to manage application options.
//
// Usage index
//
// Using this library, follow each step below.
//
// 1. Define options.
//
// 2. Create a main package to generate.
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
// you can embed `types.XXXPropertyBase` like below.
//
// Sample:
//
//   type Token struct {
//   	types.StringPropertyBase
//   }
//
// Create a main package to generate
//
// See: Example (Generate).
//
// `appenv` does NOT provide any tools like `xxx-generate`.
// Creating a main package, calling it, you can get the code.
// To generate, you may call `appenv.Generate` function with option properties.
//
// Option properties are built by `appenv.Prop` from `Value`s that you defined in above.
// `appenv.Prop` receives `Store` options that specify where the option will be stored to or loaded from.
// Now `appenv` supports Stores: YAML file, keyring or environment variables
//
// Each option can store to / be loaded from multiple `Store`.
//
// Generation code should be tagged like `// +build generate`.
// The tag may prevent the generator from being unintendedly built in your `$GOBIN`.
// You can write the go:generate directive to call it from `go generate`.
//
//   //go:generate go run -tags generate .
//
// Access options with generated function
//
// See: Example (GetAccess)
//
// Configure options with generated function
//
// See: Example (GetConfig)
package appenv
