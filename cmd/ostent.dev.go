// +build !bin

package cmd

import (
	"go/build"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	"github.com/ostrost/ostent/cmd/cmdcobra"
)

// OstentPkgPath defined for looking up the package directory.
const OstentPkgPath = "github.com/ostrost/ostent"

var (
	// GenDocCmd represents the gendoc subcommand
	GenDocCmd = &cobra.Command{
		Use:   "gendoc",
		Short: "Generate ostent commands docs",
		RunE:  GenDocRunE,
	}
	// GenDocDir is the flag value.
	GenDocDir string
)

func init() {
	pkg, err := build.Import(OstentPkgPath, "", build.FindOnly)
	if err != nil {
		log.Fatal(err)
	}
	GenDocDir = filepath.Join(pkg.Dir, "doc")
	OstentCmd.PersistentFlags().StringVar(&cmdcobra.ProfileHeapOutput, "profile-heap", "",
		"Profiling heap output `filename`")
	OstentCmd.PersistentFlags().StringVar(&cmdcobra.ProfileCPUOutput, "profile-cpu", "",
		"Profiling CPU output `filename`")
	cmdcobra.PersistentPreRuns.Add(cmdcobra.ProfileHeapRun)
	cmdcobra.PersistentPreRuns.Add(cmdcobra.ProfileCPURun)
	GenDocCmd.Flags().StringVar(&GenDocDir, "directory", GenDocDir,
		"Output `directory` for saving docs")
	OstentCmd.AddCommand(GenDocCmd)
}

func GenDocRunE(*cobra.Command, []string) error {
	OstentCmd.DisableAutoGenTag = true
	if cmd, _, err := OstentCmd.Find([]string{"gendoc"}); err == nil {
		// err is gone
		OstentCmd.RemoveCommand(cmd)
	}
	if err := doc.GenMarkdownTree(OstentCmd, GenDocDir); err != nil {
		return err
	}
	return nil
}
