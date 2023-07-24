package cmd

import (
	"os"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/project"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/cli"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ctype string
var addAll bool

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes to Git",
	Run: func(cmd *cobra.Command, args []string) {
		if !cli.CheckConfig() {
			color.Red("\nSorry, you need to initialize the project first.")
			os.Exit(1)
		}
		project.ReadConfigFile(true)

		msg := ""
		body := ""

		if len(message) > 0 {
			msg = message[0]
		}

		if len(message) > 1 {
			body = message[1]
		}

		config.AppData.Commit.AddAll = addAll
		config.AppData.Commit.Message = msg
		config.AppData.Commit.Scope = getScope(true)
		config.AppData.Commit.Type = ctype
		config.AppData.Commit.Body = body

		pkg_action.Commit(true)
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
