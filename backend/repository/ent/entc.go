//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
)

func main() {
	if err := entc.Generate("./schema", &gen.Config{}, entc.Extensions(entviz.Extension{})); err != nil {
		log.Fatal("failed to generate ER figure", err)
	}
}
