package protocgengrpcgateway

import (
	"github.com/zeromicro/go-zero/tools/goctl/vars"
	"os/exec"
	"runtime"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/pkg/goctl"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/golang"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
)

const (
	Name                    = "protoc-gen-grpc-gateway"
	binProtocGenGrpcGateway = "protoc-gen-grpc-gateway"
	url                     = "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest"
)

func Install(cacheDir string) (string, error) {
	return goctl.Install(cacheDir, Name, func(dest string) (string, error) {
		err := golang.Install(url)
		return dest, err
	})
}

func Exists() bool {
	_, err := lookUpProtocGenGrpcGateway()
	return err == nil
}

// Version is used to get the version of the protoc-gen-go-grpc plugin.
func Version() (string, error) {
	path, err := lookUpProtocGenGrpcGateway()
	if err != nil {
		return "", err
	}
	version, err := execx.Run(path+" --version", "")
	if err != nil {
		return "", err
	}
	fields := strings.Fields(version)
	if len(fields) > 1 {
		return fields[1], nil
	}
	return "", nil
}

func lookUpProtocGenGrpcGateway() (string, error) {
	suffix := getExeSuffix()
	xProtocGenGoGrpc := binProtocGenGrpcGateway + suffix
	return LookPath(xProtocGenGoGrpc)
}

func getExeSuffix() string {
	if runtime.GOOS == vars.OsWindows {
		return ".exe"
	}
	return ""
}

func LookPath(xBin string) (string, error) {
	suffix := getExeSuffix()
	if len(suffix) > 0 && !strings.HasSuffix(xBin, suffix) {
		xBin = xBin + suffix
	}

	bin, err := exec.LookPath(xBin)
	if err != nil {
		return "", err
	}
	return bin, nil
}
