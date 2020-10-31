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

## Usage

Read documents: [pkg.go.dev](https://pkg.go.dev/github.com/kyoh86/appenv/gen)

## WHY?

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

## Install

```console
$ go get github.com/kyoh86/appenv
```

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
