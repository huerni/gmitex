package cmd

import (
	"github.com/huerni/gmitex/gmctl/generator"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

// gmctl update <name>
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the microservice.",
	Long:  `Update the microservice.`,
	Args:  cobra.ExactArgs(0),
	RunE:  gmUpdate,
}

var (
	ProtoSrc string
)

func gmUpdate(_ *cobra.Command, args []string) error {

	pwdDir, _ := os.Getwd()
	outputDir := filepath.Join(pwdDir, "pb")
	fileExt := filepath.Ext(ProtoSrc)
	if fileExt == ".proto" {
		ctx := &generator.GmContext{
			Op:     "update",
			Src:    ProtoSrc,
			Output: outputDir,
		}
		g := generator.NewGenerator()
		err := g.Generate(ctx)
		if err != nil {
			return err
		}
	} else {
		protoDir := filepath.Join(pwdDir, "proto")
		rd, err := ioutil.ReadDir(protoDir)
		if err != nil {
			return err
		}
		for _, fi := range rd {
			if !fi.IsDir() {
				fullDir := protoDir + "/" + fi.Name()
				ctx := &generator.GmContext{
					Op:     "update",
					Src:    fullDir,
					Output: pwdDir,
				}
				g := generator.NewGenerator()
				err := g.Generate(ctx)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(updateCmd)
	// 绑定flag  -name  -output
	pwdDir, _ := os.Getwd()
	newCmd.Flags().StringVarP(&ProtoSrc, "src", "s", filepath.Join(pwdDir, "proto"), "proto src")
}
