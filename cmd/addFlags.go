package cmd

import "github.com/spf13/cobra"

var feat, fix, docs, refactor, build, chore, ci, perf, revert, style, test, merge, ticket, baseBranch string
var message []string

type AddFlags struct{}

func (x AddFlags) Commit(_cmd *cobra.Command) {
	_cmd.PersistentFlags().StringArrayVarP(&message, "message", "m", []string{}, "Commit message")
	_cmd.PersistentFlags().StringVarP(&feat, "feat", "F", "", "A new feature ✨")
	_cmd.PersistentFlags().StringVarP(&fix, "fix", "f", "", "A bug fix 🐛")
	_cmd.PersistentFlags().StringVarP(&docs, "docs", "d", "", "Documentation only changes 📚")
	_cmd.PersistentFlags().StringVarP(&style, "style", "", "", "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc) 💎")
	_cmd.PersistentFlags().StringVarP(&refactor, "refactor", "r", "", "A code change that neither fixes a bug nor adds a feature 📦")
	_cmd.PersistentFlags().StringVarP(&perf, "perf", "p", "", "A code change that improves performance 🚀")
	_cmd.PersistentFlags().StringVarP(&test, "test", "", "", "Adding missing tests or correcting existing tests 🧪")
	_cmd.PersistentFlags().StringVarP(&build, "build", "b", "", "Changes that affect the build system or external dependencies 🛠")
	_cmd.PersistentFlags().StringVarP(&ci, "ci", "", "", "Changes to our CI configuration files and scripts ⚙️")
	_cmd.PersistentFlags().StringVarP(&chore, "chore", "c", "", "Other changes that don't modify src or test files ♻️")
	_cmd.PersistentFlags().StringVarP(&revert, "revert", "", "", "Reverts a previous commit 🗑")
	_cmd.PersistentFlags().BoolVarP(&addAll, "all", "a", false, "Add all files changed to commit")
}

func (x AddFlags) Push(_cmd *cobra.Command) {
	x.Commit(_cmd)
	_cmd.PersistentFlags().StringVarP(&merge, "merge", "M", "", "Merge a branch")
}
