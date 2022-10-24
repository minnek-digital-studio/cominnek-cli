package git_controller

import (
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func style(data string) string {
	datas := strings.Split(data, ",")
	newData := ""

	for i, _data := range datas {
		_data = crossOutDeleted(_data)
		coma := ","

		if i == len(datas)-1 {
			coma = ""
		}

		newData += _data + coma
	}

	return newData
}

func crossOutDeleted(data string) string {
	crossOut := color.New(color.CrossedOut).Add(color.Bold).Add(color.FgRed).SprintFunc()
	datas := strings.Split(data, "D ")

	for i, data := range datas {
		if i == 0 {
			continue
		}

		datas[i] = crossOut(data)
	}

	return strings.Join(datas, "D ")
}

func cleanChangesData(data string) string {
	data = strings.ReplaceAll(data, "?? ", color.HiGreenString("U "))
	data = strings.ReplaceAll(data, "D ", color.HiRedString("D "))
	data = strings.ReplaceAll(data, "M ", color.HiYellowString("M "))
	data = strings.ReplaceAll(data, "A ", color.HiBlueString("A "))

	data = strings.TrimSpace(data)

	return data
}

func cleanRaw(data string) string {
	data = strings.ReplaceAll(data, "A ", "")
	data = strings.ReplaceAll(data, "M ", "")
	data = strings.ReplaceAll(data, "D ", "")
	data = strings.ReplaceAll(data, "U ", "")
	data = strings.ReplaceAll(data, "?? ", "")
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

	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range raw {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	raw = list

	return
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

	raw = strings.Split(cleanRaw(out), ",")

	out = style(out)
	styled = strings.Split(cleanChangesData(out), ",")
	return styled, raw
}
