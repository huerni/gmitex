package gw

import (
	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"gmctl/util"
	"path/filepath"
	"strings"
)

const (
	cmd      = "cmd"
	etc      = "etc"
	internal = "internal"
	app      = "app"
	config   = "config"
)

type (
	DirContext interface {
		GetEtc() Dir
		GetInternal() Dir
		GetApp() Dir
		GetConfig() Dir
		GetCmd() Dir
		GetProjectName() string
	}

	Dir struct {
		Base     string
		Filename string
		Package  string
	}

	defaultDirContext struct {
		inner       map[string]Dir
		projectName string
		ctx         *ctx.ProjectContext
	}
)

func mkdir(ctx *ctx.ProjectContext, gctx *GwContext) (DirContext, error) {

	inner := make(map[string]Dir)
	etcDir := filepath.Join(ctx.WorkDir, "etc")
	internalDir := filepath.Join(ctx.WorkDir, "internal")
	appDir := filepath.Join(internalDir, "app")
	configDir := filepath.Join(internalDir, "config")
	cmdDir := filepath.Join(ctx.WorkDir, "cmd")

	inner[cmd] = Dir{
		Filename: cmdDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(cmdDir, ctx.Dir))),
		Base:     filepath.Base(cmdDir),
	}
	inner[etc] = Dir{
		Filename: etcDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(etcDir, ctx.Dir))),
		Base:     filepath.Base(etcDir),
	}
	inner[internal] = Dir{
		Filename: internalDir,
		Package: filepath.ToSlash(filepath.Join(ctx.Path,
			strings.TrimPrefix(internalDir, ctx.Dir))),
		Base: filepath.Base(internalDir),
	}
	inner[app] = Dir{
		Filename: appDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(appDir, ctx.Dir))),
		Base:     filepath.Base(appDir),
	}
	inner[config] = Dir{
		Filename: configDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(configDir, ctx.Dir))),
		Base:     filepath.Base(configDir),
	}

	for _, v := range inner {
		err := util.MkdirIfNotExist(v.Filename)
		if err != nil {
			return nil, err
		}
	}

	return &defaultDirContext{
		ctx:         ctx,
		inner:       inner,
		projectName: gctx.ProjectName,
	}, nil
}

func (d *defaultDirContext) GetEtc() Dir {
	return d.inner[etc]
}

func (d *defaultDirContext) GetInternal() Dir {
	return d.inner[internal]
}

func (d *defaultDirContext) GetApp() Dir {
	return d.inner[app]
}

func (d *defaultDirContext) GetConfig() Dir {
	return d.inner[config]
}

func (d *defaultDirContext) GetCmd() Dir {
	return d.inner[cmd]
}

func (d *defaultDirContext) GetProjectName() string {
	return d.projectName
}

// Valid returns true if the directory is valid
func (d *Dir) Valid() bool {
	return len(d.Filename) > 0 && len(d.Package) > 0
}
