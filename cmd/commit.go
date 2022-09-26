package cmd

import (
	"log"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var feat, fix, docs, refactor, test, build, body string
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
	commitCmd.PersistentFlags().StringVarP(&body, "body", "m", "", "Commit body message")
	commitCmd.PersistentFlags().StringVarP(&feat, "feature", "F", "{false}", "Add a new feature")
	commitCmd.PersistentFlags().StringVarP(&fix, "fix", "f", "{false}", "Fix an existing issue")
	commitCmd.PersistentFlags().StringVarP(&docs, "docs", "d", "{false}", "Add documentation")
	commitCmd.PersistentFlags().StringVarP(&refactor, "refactor", "r", "{false}", "Refactor an existing issue")
	commitCmd.PersistentFlags().StringVarP(&test, "test", "t", "{false}", "Add tests")
	commitCmd.PersistentFlags().StringVarP(&build, "build", "b", "{false}", "Build the project")
	rootCmd.AddCommand(commitCmd)
}

func getScope(skipFatal bool) string {
	if feat != "{false}" {
		ctype = "feat"
		return feat
	} else if fix != "{false}" {
		ctype = "fix"
		return fix
	} else if docs != "{false}" {
		ctype = "docs"
		return docs
	} else if refactor != "{false}" {
		ctype = "refactor"
		return refactor
	} else if test != "{false}" {
		ctype = "test"
		return test
	} else if build != "{false}" {
		ctype = "build"
		return build
	} else {
		if !skipFatal {
			log.Fatal("No commit type specified")
			return ""
		} else {
			return ""
		}
	}
}
