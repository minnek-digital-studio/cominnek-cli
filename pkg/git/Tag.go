package git

import (
	"fmt"
	"os"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func AddTag(tag string) {
	cmd := fmt.Sprintf("git tag %s", tag)
	out, errout, err := shell.Out(cmd)

	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		events.App.Emit("cleanup")

		os.Exit(1)
	}

	fmt.Println("Added tag", tag)
}

func PublishTag(tag string) {
	cmd := fmt.Sprintf("git push origin %s", tag)
	out, errout, err := shell.Out(cmd)

	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		events.App.Emit("cleanup")

		os.Exit(1)
	}

	fmt.Println("Published tag", tag)
}
