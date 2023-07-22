package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gmctl/generator"
	"path/filepath"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create an initial microservice.",
	Long:  `Create an initial microservice.`,
	Args:  cobra.ExactArgs(1),
	RunE:  gmNew,
}

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
	err = generator.GenProto(src)
	if err != nil {
		return err
	}

	ctx := &generator.GmContext{
		Src:    src,
		Output: "/home/test/example",
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
}
