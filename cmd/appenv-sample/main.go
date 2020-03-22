// +build generate

package main

import (
	"log"

	"github.com/kyoh86/appenv/env"
	"github.com/kyoh86/appenv/gen"
)

//go:generate go run -tags generate ./main.go

func main() {
	g := &gen.Generator{
		BuildTag: "sample",
	}

	if err := g.Do(
		"github.com/kyoh86/appenv/env",
		"../../env",
		gen.Prop(new(env.GithubHost), gen.YAML(), gen.Envar()),
		gen.Prop(new(env.GithubUser), gen.YAML(), gen.Envar()),
		gen.Prop(new(env.Roots), gen.YAML(), gen.Envar()),
		gen.Prop(new(env.Hooks), gen.YAML(), gen.Envar()),
	); err != nil {
		log.Fatalln(err)
	}
}
