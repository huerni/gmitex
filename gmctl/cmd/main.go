package main

import (
	"gmctl/generator"
)

func main() {
	ctx := &generator.GmContext{
		Src:      "user.proto",
		Output:   "/home/test/example",
		GoModule: "gmitex",
	}
	g := generator.NewGenerator()
	err := g.Generate(ctx)
	if err != nil {
		panic(err)
	}
}
