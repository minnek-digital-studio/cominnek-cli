package cmd

import (
	"log"
	"os"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ctype string

var commitCmd = &cobra.Command{
	Use:   "commit <message>",
	Short: "Commit changes to Git",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		msg := args[0]

		if msg == "" {
			log.Fatal("No commit message specified")
		}

		git.Commit(msg, body, ctype, getScope(false))
	},
}

func init() {
	AddFlags{}.Commit(commitCmd)
	rootCmd.AddCommand(commitCmd)
}

func getScope(skipFatal bool) string {
	if feat != "" {
		ctype = "feat"
		return feat
	}

	if fix != "" {
		ctype = "fix"
		return fix
	}

	if docs != "" {
		ctype = "docs"
		return docs
	}

	if refactor != "" {
		ctype = "refactor"
		return refactor
	}

	if build != "" {
		ctype = "build"
		return build
	}

	if style != "" {
		ctype = "style"
		return style
	}

	if chore != "" {
		ctype = "chore"
		return chore
	}

	if ci != "" {
		ctype = "ci"
		return ci
	}

	if perf != "" {
		ctype = "perf"
		return perf
	}

	if revert != "" {
		ctype = "revert"
		return revert
	}

	if test != "" {
		ctype = "test"
		return test
	}

	if !skipFatal {
		println(color.HiRedString("No commit type specified"))
		println("Use --feat, --fix, --docs, --refactor, --build, --style, --chore, --ci, --perf, --revert, --test")
		os.Exit(1)
		return ""
	} else {
		return ""
	}

}
