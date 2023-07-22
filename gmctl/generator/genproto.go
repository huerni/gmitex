package generator

import (
	_ "embed"
	utilx "github.com/zeromicro/go-zero/tools/goctl/util"
	"gmctl/util"
	"path/filepath"
	"strings"
)

//go:embed tpl/server/proto/proto.tpl
var protoTemplate string

func GenProto(out string) error {
	protoFilename := filepath.Base(out)
	serverName := strings.TrimSuffix(protoFilename, filepath.Ext(protoFilename))
	dir := filepath.Dir(out)
	err := util.MkdirIfNotExist(dir)
	if err != nil {
		return err
	}
	return utilx.With("proto").Parse(protoTemplate).SaveTo(map[string]any{
		"package":    serverName,
		"serverName": serverName,
	}, out, false)
}
