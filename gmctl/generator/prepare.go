package generator

import "github.com/huerni/gmitex/gmctl/envcheck"

func (g *Generator) Prepare() error {
	return envcheck.Prepare(true, false, false)
}
