package generator

import (
	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"gmctl/parser"
	"gmctl/util"
	"path/filepath"
	"strings"
)

const (
	cmd      = "cmd"
	etc      = "etc"
	internal = "internal"
	config   = "config"
	errno    = "errno"
	handler  = "handler"
	service  = "server"
	svc      = "svc"
	pb       = "pb"
	protoGo  = "proto"
)

type (
	DirContext interface {
		GetEtc() Dir
		GetInternal() Dir
		GetConfig() Dir
		GetErrno() Dir
		GetHandler() Dir
		GetService() Dir
		GetSvc() Dir
		GetPb() Dir
		GetCmd() Dir
		GetProtoGo() Dir
		GetServerName() string
	}

	Dir struct {
		Base     string
		Filename string
		Package  string
	}

	defaultDirContext struct {
		inner      map[string]Dir
		serverName string
		ctx        *ctx.ProjectContext
	}
)

func mkdir(ctx *ctx.ProjectContext, proto parser.Proto, gctx *GmContext) (DirContext, error) {

	inner := make(map[string]Dir)
	etcDir := filepath.Join(ctx.WorkDir, "etc")
	internalDir := filepath.Join(ctx.WorkDir, "internal")
	configDir := filepath.Join(internalDir, "config")
	errnoDir := filepath.Join(internalDir, "errno")
	handlerDir := filepath.Join(internalDir, "handler")
	serviceDir := filepath.Join(internalDir, "service")
	svcDir := filepath.Join(internalDir, "svc")
	pbDir := filepath.Join(ctx.WorkDir, "pb")
	cmdDir := filepath.Join(ctx.WorkDir, "cmd")
	protoGoDir := filepath.Join(ctx.WorkDir, "proto")

	inner[cmd] = Dir{
		Filename: cmdDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(cmdDir, ctx.Dir))),
		Base:     filepath.Base(cmdDir),
	}
	inner[protoGo] = Dir{
		Filename: protoGoDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(protoGoDir, ctx.Dir))),
		Base:     filepath.Base(protoGoDir),
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
	inner[config] = Dir{
		Filename: configDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(configDir, ctx.Dir))),
		Base:     filepath.Base(configDir),
	}
	inner[errno] = Dir{
		Filename: errnoDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(errnoDir, ctx.Dir))),
		Base:     filepath.Base(errnoDir),
	}
	inner[handler] = Dir{
		Filename: handlerDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(handlerDir, ctx.Dir))),
		Base:     filepath.Base(handlerDir),
	}
	inner[service] = Dir{
		Filename: serviceDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(serviceDir, ctx.Dir))),
		Base:     filepath.Base(service),
	}
	inner[svc] = Dir{
		Filename: svcDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(svcDir, ctx.Dir))),
		Base:     filepath.Base(svcDir),
	}
	inner[pb] = Dir{
		Filename: pbDir,
		Package:  filepath.ToSlash(filepath.Join(ctx.Path, strings.TrimPrefix(pbDir, ctx.Dir))),
		Base:     filepath.Base(pbDir),
	}

	for _, v := range inner {
		err := util.MkdirIfNotExist(v.Filename)
		if err != nil {
			return nil, err
		}
	}
	serverName := strings.TrimSuffix(proto.Name, filepath.Ext(proto.Name))
	serverName = strings.ReplaceAll(serverName, "-", "")
	return &defaultDirContext{
		ctx:        ctx,
		inner:      inner,
		serverName: ctx.Name,
	}, nil
}

func (d *defaultDirContext) GetEtc() Dir {
	return d.inner[etc]
}

func (d *defaultDirContext) GetInternal() Dir {
	return d.inner[internal]
}

func (d *defaultDirContext) GetConfig() Dir {
	return d.inner[config]
}

func (d *defaultDirContext) GetErrno() Dir {
	return d.inner[errno]
}

func (d *defaultDirContext) GetHandler() Dir {
	return d.inner[handler]
}

func (d *defaultDirContext) GetService() Dir {
	return d.inner[service]
}

func (d *defaultDirContext) GetSvc() Dir {
	return d.inner[svc]
}

func (d *defaultDirContext) GetPb() Dir {
	return d.inner[pb]
}

func (d *defaultDirContext) GetCmd() Dir {
	return d.inner[cmd]
}

func (d *defaultDirContext) GetProtoGo() Dir {
	return d.inner[protoGo]
}

func (d *defaultDirContext) GetServerName() string {
	return d.serverName
}

// Valid returns true if the directory is valid
func (d *Dir) Valid() bool {
	return len(d.Filename) > 0 && len(d.Package) > 0
}
