package generator

import "gmctl/envcheck"

func (g *Generator) Prepare() error {
	return envcheck.Prepare(true, false, false)
}
