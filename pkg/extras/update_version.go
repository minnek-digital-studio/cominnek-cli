package extras

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
)

func UpdateVersion(version string) {
	msg := "update version to " + version

	git.PushWithOutTicket(msg, "", "build", "")
}
