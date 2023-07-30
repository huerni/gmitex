package generator

import (
	_ "embed"
	"github.com/huerni/gmitex/gmctl/util"
	utilx "github.com/zeromicro/go-zero/tools/goctl/util"
	"path/filepath"
	"strings"
)

//go:embed tpl/server/proto/proto.tpl
var protoTemplate string

//go:embed tpl/server/proto/google/api/annotations.tpl
var annoTemplate string

//go:embed tpl/server/proto/google/api/field_behavior.tpl
var fieldTemplate string

//go:embed tpl/server/proto/google/api/http.tpl
var httpTemplate string

//go:embed tpl/server/proto/google/api/httpbody.tpl
var httpbodyTemplate string

//go:embed tpl/server/proto/google/protobuf/api.tpl
var apiTemplate string

//go:embed tpl/server/proto/google/protobuf/descriptor.tpl
var descrTemplate string

func GenProto(out string) error {
	protoFilename := filepath.Base(out)
	serverName := strings.TrimSuffix(protoFilename, filepath.Ext(protoFilename))
	dir := filepath.Dir(out)
	err := util.MkdirIfNotExist(dir)
	if err != nil {
		return err
	}

	err = GenGoogle(dir)
	if err != nil {
		return err
	}

	return utilx.With("proto").Parse(protoTemplate).SaveTo(map[string]any{
		"package":    serverName,
		"serverName": serverName,
	}, out, false)
}

func GenGoogle(baseDir string) error {
	googledir := filepath.Join(baseDir, "google")
	apiDir := filepath.Join(googledir, "api")
	protoDir := filepath.Join(googledir, "protobuf")
	err := util.MkdirIfNotExist(apiDir)
	if err != nil {
		return err
	}
	err = util.MkdirIfNotExist(protoDir)
	if err != nil {
		return err
	}

	err = utilx.With("proto").Parse(annoTemplate).SaveTo(map[string]any{}, filepath.Join(apiDir, "annotations.proto"), false)
	if err != nil {
		return err
	}

	err = utilx.With("proto").Parse(fieldTemplate).SaveTo(map[string]any{}, filepath.Join(apiDir, "field_behavior.proto"), false)
	if err != nil {
		return err
	}

	err = utilx.With("proto").Parse(httpTemplate).SaveTo(map[string]any{}, filepath.Join(apiDir, "http.proto"), false)
	if err != nil {
		return err
	}

	err = utilx.With("proto").Parse(httpbodyTemplate).SaveTo(map[string]any{}, filepath.Join(apiDir, "httpbody.proto"), false)
	if err != nil {
		return err
	}

	err = utilx.With("proto").Parse(apiTemplate).SaveTo(map[string]any{}, filepath.Join(protoDir, "api.proto"), false)
	if err != nil {
		return err
	}

	err = utilx.With("proto").Parse(descrTemplate).SaveTo(map[string]any{}, filepath.Join(protoDir, "descriptor.proto"), false)
	if err != nil {
		return err
	}

	return nil
}
