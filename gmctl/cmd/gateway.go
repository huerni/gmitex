package cmd

import (
	"fmt"
	"github.com/huerni/gmitex/gmctl/generator/gw"
	"github.com/huerni/gmitex/gmctl/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// gmctl gateway --projectname <> --output <>
// gmctl gateway <projectname>
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Create gateway.",
	Long:  `Create gateway`,
	Args:  cobra.ExactArgs(0),
	RunE:  gmGateway,
}

var (
	GwOut       string
	ProjectName string
)

func gmGateway(_ *cobra.Command, args []string) error {
	projectName := ProjectName
	outputDir, err := filepath.Abs(Out)
	if err != nil {
		return err
	}
	ext := filepath.Ext(projectName)
	if len(ext) > 0 {
		return fmt.Errorf("unexpected ext: %s", ext)
	}
	gatewayoutput := fmt.Sprintf("%v/gateway", outputDir)
	err = util.MkdirIfNotExist(gatewayoutput)
	if err != nil {
		return err
	}

	ctx := &gw.GwContext{
		ProjectName: projectName,
		Output:      gatewayoutput,
	}
	g := gw.NewGenerator()
	err = g.Generate(ctx)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(gatewayCmd)
	// 绑定flag  -name  -output
	pwdDir, _ := os.Getwd()
	gatewayCmd.Flags().StringVarP(&GwOut, "out", "o", pwdDir, "gateway output dir")
	gatewayCmd.Flags().StringVarP(&ProjectName, "project", "p", "gmitex", "belongs to the project")
}
