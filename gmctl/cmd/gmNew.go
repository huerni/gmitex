package cmd

import (
	"fmt"
	"github.com/huerni/gmitex/gmctl/generator"
	"github.com/huerni/gmitex/gmctl/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// gmctl new --name <> --output <>
// gmctl new <name>
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create an initial microservice.",
	Long:  `Create an initial microservice.`,
	Args:  cobra.ExactArgs(1),
	RunE:  gmNew,
}

var (
	Out string
	//ProjectName string
)

func gmNew(_ *cobra.Command, args []string) error {
	servername := args[0]
	ext := filepath.Ext(servername)
	if len(ext) > 0 {
		return fmt.Errorf("unexpected ext: %s", ext)
	}

	protoName := fmt.Sprintf("proto/%v.proto", servername)
	filename := filepath.Join(".", servername, protoName)
	src, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	outputDir, err := filepath.Abs(Out)
	if err != nil {
		return err
	}
	outputDir = filepath.Join(outputDir, servername)

	err = util.MkdirIfNotExist(outputDir)
	if err != nil {
		return err
	}

	err = generator.GenProto(src)
	if err != nil {
		return err
	}

	ctx := &generator.GmContext{
		Op:     "new",
		Src:    src,
		Output: outputDir,
	}
	g := generator.NewGenerator()
	err = g.Generate(ctx)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)
	// 绑定flag  -name  -output
	pwdDir, _ := os.Getwd()
	newCmd.Flags().StringVarP(&Out, "out", "o", pwdDir, "server output dir")
	newCmd.Flags().StringVarP(&ProjectName, "project", "p", "gmitex", "belongs to the project")
}
