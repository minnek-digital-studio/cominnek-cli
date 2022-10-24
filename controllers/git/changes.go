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

func cleanRaw(data string) string {
	// just keep route
	data = strings.ReplaceAll(data, "A", "")
	data = strings.ReplaceAll(data, "M", "")
	data = strings.ReplaceAll(data, "D", "")
	data = strings.ReplaceAll(data, "U", "")
	data = strings.ReplaceAll(data, " ", "")

	return data
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

	if out == "" {
		return []string{}, []string{}
	}

	out = strings.TrimSpace(out)
	out = strings.ReplaceAll(out, "\n ", "\n")
	out = strings.ReplaceAll(out, "\n", ",")
	out = strings.TrimSuffix(out, ",")

	styled = strings.Split(cleanChangesData(out), ",")
	raw = strings.Split(cleanRaw(out), ",")
	return styled, raw
}
