package generator

import (
	"errors"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
	"os"
	"path/filepath"
)

func (g *Generator) GenGoMod(gctx *GmContext, workDir string) error {
	hasGoMod, err := HasGoMod(workDir)
	if err != nil {
		return err
	}
	if hasGoMod == false {
		name := gctx.GoModule
		if name == "" {
			name = filepath.Base(workDir)
			gctx.GoModule = name
		}
		_, err = execx.Run("go mod init "+name, workDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func HasGoMod(workDir string) (bool, error) {
	if len(workDir) == 0 {
		return false, errors.New("the work directory is not found")
	}
	if _, err := os.Stat(workDir); err != nil {
		return false, err
	}

	data, err := execx.Run("go list -m -f '{{.GoMod}}'", workDir)
	if err != nil || len(data) == 0 {
		return false, nil
	}

	return true, nil
}
