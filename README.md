# appenv

Application options manager for golang

[![Go Report Card](https://goreportcard.com/badge/github.com/kyoh86/appenv)](https://goreportcard.com/report/github.com/kyoh86/appenv)
[![Coverage Status](https://img.shields.io/codecov/c/github/kyoh86/appenv.svg)](https://codecov.io/gh/kyoh86/appenv)

## What's this?

In developing applications in golang, we want to manage many options for it.
And they often are in a file, envars or the keyring.

It is too heavy and bores developers to manage them.
So appenv generates codes to do it.

- Load/save function
- Configuration accessor (get, set or unset them with application)

## How to use it?

- Define options for your app.
- Generate codes.
- Call them from your app!

### Define options

You can define options with just creating structs that
implement `appenv/types.Value` that
have `Value()`, `Default()`, `MarshalText` and `UnmarshalText`.

If you want to define a primitive (like `string`) option,
you can embed `types.XXXPropertyBase` like below.

Sample:

```go
type Token struct {
	types.StringPropertyBase
}
```

### Generate codes

`appenv` does NOT provide any tools like `xxx-generate`.
Creating a main package, calling it, you can get the code.

To get the code, you may create new `appenv/gen.Generator`
instance and call its `Do` function with some options:

- Generating package path (full-name)
- Output directory
- Option properties

Option properties are built by `appenv/gen.Prop` from
`Value` s that you defined in above.
`gen.Prop` receives `Store` options that specify 
where the option will be stored to or loaded from.

Now `appenv` supports some Stores like below.

- YAML
- Keyring
- Envar (environment variable)

Each option can store to / be loaded from multiple `Store`
like `YAML`, `Keyring` and `Envar`.

```go
// +build generate

package main

import (
	"log"

	"github.com/kyoh86/appenv/gen"
	"github.com/kyoh86/appenv/env"
)

//go:generate go run -tags generate ./main.go

func main() {
	g := &gen.Generator{}

	if err := g.Do(
		"github.com/kyoh86/appenv/env",
		"../",
		gen.Prop(new(env.Token), gen.YAML(), gen.Keyring(), gen.Envar()),
	); err != nil {
		log.Fatal(err)
	}
}
```

#### WHY?

Why appenv does NOT provide any generator:

Even if it is done, there's not much diference in usage.

i.e. You may create the generation shell script (or Makefile or ...).

```sh
appenv-gen \
  -package github.com/kyoh86/appenv/env \
  -outdir ../ \
  -prop github.com/kyoh86/appenv/env.Token -store keyring -store envar
```

I think that we can read and maintin go code easier than shell script.

#### Note

- Generation code should be tagged like `// +build generate`.
The tag may prevent the generator from being unintendedly built in your `$GOBIN`.
- You can write the go:generate directive to call it from `go generate`.

### Use them

Generated codes can be used like below.

#### To get value

```go
import (
	// Import generated package
	"github.com/kyoh86/appenv/env"
)

const (
  configFile     = "~/.config/appenv/config.yaml"
  keyringService = "appenv.kyoh86.dev" // any unique service name
  envarPrefix    = "APPENV_" 
)

file, _ := os.Open(configFile)
access, _ := env.GetAccess(file, keyringService, envarPrefix)
println(access.Token())
```

#### To change configure

```go
import (
	// Import generated package
	"github.com/kyoh86/appenv/env"
)

const (
  configFile     = "~/.config/appenv/config.yaml"
  keyringService = "appenv.kyoh86.dev" // any unique service name
  envarPrefix    = "APPENV_" 
)

file, _ := os.Open(configFile)
config, _ := env.GetConfig(file, keyringService, envarPrefix)
githubToken, _ := config.Property("github.token")
githubToken.Set("foobar")
config.SaveFile(file, keyringService) // NOTE: Envar will not exported.
```

## Install

```
go get github.com/kyoh86/appenv
```

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
