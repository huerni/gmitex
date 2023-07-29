package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gmctl/generator/gw"
	"gmctl/util"
	"path/filepath"
)

// gmctl gateway --projectname <> --output <>
// gmctl gateway <projectname>
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Create gateway.",
	Long:  `Create gateway`,
	Args:  cobra.ExactArgs(2),
	RunE:  gmGateway,
}

func gmGateway(_ *cobra.Command, args []string) error {
	projectName := args[0]
	outputDir := args[1]
	ext := filepath.Ext(projectName)
	if len(ext) > 0 {
		return fmt.Errorf("unexpected ext: %s", ext)
	}
	gatewayoutput := fmt.Sprintf("%v/gateway", outputDir)
	err := util.MkdirIfNotExist(gatewayoutput)
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
}
