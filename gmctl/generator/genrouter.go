package generator

import (
	_ "embed"
	"fmt"
	"github.com/huerni/gmitex/gmctl/parser"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"path/filepath"
	"strings"
)

//go:embed tpl/server/internal/router/router.tpl
var RouterTemplate string

func (g *Generator) GenRouter(ctx DirContext, proto parser.Proto) error {
	dir := ctx.GetRouter()
	service := proto.Service[0]
	serverFileName := fmt.Sprintf("%sRouter", service.Name)
	serverFile := filepath.Join(dir.Filename, serverFileName+".go")
	importPaths := collection.NewSet()
	paths := make([]string, 0)
	auth_paths := make([]string, 0)
	for _, rpc := range service.RPC {
		path := rpc.RPC.Options[0].AggregatedConstants[0].Literal.Source
		paths = append(paths, path)
		if strings.Contains(parser.GetComment(rpc.Doc()), ":auth") {
			auth_paths = append(auth_paths, path)
		}
	}

	pathPrefixes := GetUrlPrefixes(paths)

	for _, path := range pathPrefixes {
		importPath := fmt.Sprintf(`paths = append(paths, "%s")`, path)
		importPaths.AddStr(importPath)
	}

	importAuthPaths := collection.NewSet()
	for _, path := range auth_paths {
		importPath := fmt.Sprintf(`authPaths = append(authPaths, "%s")`, path)
		importAuthPaths.AddStr(importPath)
	}

	return util.With("router").GoFmt(true).Parse(RouterTemplate).SaveTo(map[string]any{
		"serverName":      stringx.From(service.Name).ToCamel(),
		"importPaths":     strings.Join(importPaths.KeysStr(), pathx.NL),
		"importAuthPaths": strings.Join(importAuthPaths.KeysStr(), pathx.NL),
	}, serverFile, true)
}

func longestCommonURLPrefix(path1 string, path2 string) string {
	// 假设第一个URL路径为最长前缀
	prefix := path1
	// 使用strings.Split函数将URL路径拆分为各个部分
	parts1 := strings.Split(prefix, "/")
	parts2 := strings.Split(path2, "/")

	// 找到两个路径中较短的部分数
	minParts := len(parts1)
	if len(parts2) < minParts {
		minParts = len(parts2)
	}

	// 逐个比较部分，直到找到最长前缀或者部分不匹配为止
	i := 0
	for ; i < minParts; i++ {
		if parts1[i] != parts2[i] {
			break
		}
	}

	// 更新prefix为当前的最长前缀
	prefix = strings.Join(parts1[:i], "/")

	return prefix
}

func GetUrlPrefixes(paths []string) []string {
	prefixes := make([]string, 0)

	for id, path := range paths {
		if id == 0 {
			prefixes = append(prefixes, paths[0])
		} else {
			i := 0
			for ; i < len(prefixes); i += 1 {
				tmp := longestCommonURLPrefix(path, prefixes[i])
				if tmp != "" {
					prefixes[i] = tmp
					break
				}
			}
			if i == len(prefixes) {
				prefixes = append(prefixes, path)
			}
		}
	}

	return prefixes
}
