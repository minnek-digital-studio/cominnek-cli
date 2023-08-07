package pkg_action

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
)

type ReleaseData struct {
	Commit  string
	Message string
	Hash    string
	Type    string
	Scope   string
}

func getData(str string) ReleaseData {
	commit := getCommit(str)
	message := getCommitMsg(commit)
	hash := getHash(str)
	cType := getCommitType(commit)
	scope := getScope(commit)

	return ReleaseData{
		Commit:  commit,
		Message: message,
		Hash:    hash,
		Type:    cType,
		Scope:   scope,
	}
}

func getScope(str string) string {
	return extractBasedOnReg(str, `\((.*?)\)`)
}

func getCommit(str string) string {
	msg := extractBasedOnReg(str, `message\[(.*?)\]`)
	trimmed := strings.TrimSpace(msg)
	removeDoubleSpace := strings.ReplaceAll(trimmed, "  ", " ")
	return removeDoubleSpace
}

func getCommitMsg(str string) string {
	return extractBasedOnReg(str, `^(?:\w+\((?:\w+)\): )?(.+)$`)
}

func getCommitType(str string) string {
	return extractBasedOnReg(str, `^(\w+)(\(|:|$)`)
}

func getHash(str string) string {
	return extractBasedOnReg(str, `hash\[(.*?)\]`)
}

func extractBasedOnReg(str string, reg string) string {
	re := regexp.MustCompile(reg)
	match := re.FindStringSubmatch(str)

	if len(match) >= 2 {
		return match[1]
	} else {
		return ""
	}
}

func isInArray(str string, allowedTypes []string) bool {
	for _, t := range allowedTypes {
		if t == str {
			return true
		}
	}
	return false
}

func updateVersion(v string, t string) string {
	versionArr := strings.Split(v, ".")

	if len(versionArr) != 3 {
		println("Invalid version")
		os.Exit(1)
	}

	if t == "major" {
		versionArr[0] = increaseVersion(versionArr[0])
		versionArr[1] = "0"
		versionArr[2] = "0"
	}

	if t == "minor" {
		versionArr[1] = increaseVersion(versionArr[1])
		versionArr[2] = "0"
	}

	if t == "patch" {
		versionArr[2] = increaseVersion(versionArr[2])
	}

	newVersion := strings.Join(versionArr, ".")
	return newVersion
}

func increaseVersion(v string) string {
	old_version, err := strconv.Atoi(v)

	if err != nil {
		println("Invalid version")
		os.Exit(1)
	}

	new_version := old_version + 1
	return strconv.Itoa(new_version)
}

func getReleaseFileData() string {
	fileName := "CHANGELOG.md"

	if !files.CheckExist(fileName) {

		files.Create([]byte{}, fileName)
		return ""
	}

	data := files.Read(fileName)

	return string(data)
}

func saveChangelog(changelog string) {
	fileName := "CHANGELOG.md"

	if !files.CheckExist(fileName) {
		files.Create([]byte{}, fileName)
	}

	files.Create([]byte(changelog), fileName)
}
