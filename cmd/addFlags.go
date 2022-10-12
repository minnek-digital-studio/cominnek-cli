package cmd

import "github.com/spf13/cobra"

var feat, fix, docs, refactor, build, body, merge, ticket, baseBranch string

type AddFlags struct{}

func (x AddFlags) Commit(_cmd *cobra.Command) {
	_cmd.PersistentFlags().StringVarP(&body, "body", "m", "", "Commit body message")
	_cmd.PersistentFlags().StringVarP(&feat, "feat", "F", "{false}", "Add a new feature")
	_cmd.PersistentFlags().StringVarP(&fix, "fix", "f", "{false}", "Fix an existing issue")
	_cmd.PersistentFlags().StringVarP(&docs, "docs", "d", "{false}", "Add documentation")
	_cmd.PersistentFlags().StringVarP(&refactor, "refactor", "r", "{false}", "Refactor an existing issue")
	_cmd.PersistentFlags().StringVarP(&build, "build", "b", "{false}", "Build the project")
}

func (x AddFlags) Push(_cmd *cobra.Command) {
	x.Commit(_cmd)
	_cmd.PersistentFlags().StringVarP(&merge, "merge", "M", "", "Merge a branch")
}
