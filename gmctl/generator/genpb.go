package generator

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
	"os/exec"
	"runtime"
)

func (g *Generator) GenPb(ctx DirContext, gctx *GmContext) error {
	arg := fmt.Sprintf("protoc -I %v --go_out %v --go_opt paths=source_relative --go-grpc_out %v --go-grpc_opt paths=source_relative --grpc-gateway_out %v --grpc-gateway_opt paths=source_relative %v",
		ctx.GetProtoGo().Filename, ctx.GetPb().Filename, ctx.GetPb().Filename, ctx.GetPb().Filename, gctx.Src)
	goos := runtime.GOOS
	var cmd *exec.Cmd
	switch goos {
	case vars.OsMac, vars.OsLinux:
		cmd = exec.Command("sh", "-c", arg)
	case vars.OsWindows:
		cmd = exec.Command("cmd.exe", "/c", arg)
	default:
		return fmt.Errorf("unexpected os: %v", goos)
	}
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
