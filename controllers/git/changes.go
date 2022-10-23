package git_controller

import (
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func cleanChangesData(data string) string {
	untraked := strings.ReplaceAll(data, "??", color.HiGreenString("U"))
	deleted := strings.ReplaceAll(untraked, "D", color.HiRedString("D"))
	modified := strings.ReplaceAll(deleted, "M", color.HiYellowString("M"))
	added := strings.ReplaceAll(modified, "A", color.HiBlueString("A"))
	trim := strings.TrimSpace(added)

	return trim
}

func ListUnstageChanges() (raw []string) {
	out, _, err := shell.Out("git ls-files --others --modified --deleted --exclude-standard")

	if err != nil {
		panic(err)
	}

	out = strings.ReplaceAll(out, "\n", ",")
	out = strings.TrimSpace(out)
	out = strings.TrimSuffix(out, ",")

	if out == "" {
		return []string{}
	}

	raw = strings.Split(out, ",")

	return raw
}

func ListChanges() (styled []string, raw []string) {
	out, _, err := shell.Out("git status -s -uall")

	if err != nil {
		panic(err)
	}

	out = strings.TrimSpace(out)
	out = strings.ReplaceAll(out, "\n ", "\n")

	styled = strings.Split(cleanChangesData(out), "\n")
	raw = strings.Split(out, "\n")

	return styled, raw
}
