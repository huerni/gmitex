package envcheck

import (
	"fmt"
	"github.com/huerni/gmitex/gmctl/envcheck/protocgengrpcgateway"
	"github.com/huerni/gmitex/gmctl/util"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/tools/goctl/pkg/protoc"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/protocgengo"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/protocgengogrpc"
)

type bin struct {
	name   string
	exists bool
	get    func(cacheDir string) (string, error)
}

var bins = []bin{
	{
		name:   "protoc",
		exists: protoc.Exists(),
		get:    protoc.Install,
	},
	{
		name:   "protoc-gen-go",
		exists: protocgengo.Exists(),
		get:    protocgengo.Install,
	},
	{
		name:   "protoc-gen-go-grpc",
		exists: protocgengogrpc.Exists(),
		get:    protocgengogrpc.Install,
	},
	// TODO: 添加 protoc-gen-grpc-gateway
	{
		name:   "protoc-gen-grpc-gateway",
		exists: protocgengrpcgateway.Exists(),
		get:    protocgengrpcgateway.Install,
	},
}

func Prepare(install, force, verbose bool) error {
	pending := true
	fmt.Println("[gmctl-env]: preparing to check env")
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("%+v\n", p)
			return
		}
		if pending {
			fmt.Println("\n[gmctl-env]: congratulations! your gmctl environment is ready!")
		} else {
			fmt.Println(`[gmctl-env]: check env finish, some dependencies is not found in PATH.`)
		}
	}()
	for _, e := range bins {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("")
		fmt.Printf("[gmctl-env]: looking up %v\n", e.name)
		if e.exists {
			fmt.Printf("[gmctl-env]: %v is installed\n", e.name)
			continue
		}
		fmt.Printf("[gmctl-env]: %v is not found in PATH\n", e.name)
		if install {
			install := func() {
				fmt.Printf("[gmctl-env]: preparing to install %v\n", e.name)
				// 更换cache为 .gmctl/cache
				home, err := os.UserHomeDir()
				if err != nil {
					panic(err)
				}
				gmctlCache := filepath.Join(home, ".gmctl/cache")
				err = util.MkdirIfNotExist(gmctlCache)
				if err != nil {
					panic(err)
				}
				path, err := e.get(gmctlCache)
				//path, err := e.get(env.Get(env.GoctlCache))
				if err != nil {
					fmt.Printf("[gmctl-env]: an error interrupted the installation: %+v\n", err)
					pending = false
				} else {
					fmt.Printf("[gmctl-env]: %v is already installed in %v\n", e.name, path)
				}
			}
			if force {
				install()
				continue
			}
			fmt.Printf("[gmctl-env]: do you want to install %q [y: YES, n: No]\n", e.name)
			for {
				var in string
				fmt.Scanln(&in)
				var brk bool
				switch {
				case strings.EqualFold(in, "y"):
					install()
					brk = true
				case strings.EqualFold(in, "n"):
					pending = false
					fmt.Printf("[gmctl-env]: %q installation is ignored", e.name)
					brk = true
				default:
					fmt.Printf("[gmctl-env]: invalid input, input 'y' for yes, 'n' for no")
				}
				if brk {
					break
				}
			}
		} else {
			pending = false
		}
	}
	return nil
}
