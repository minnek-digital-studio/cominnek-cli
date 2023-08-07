package ask

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
)

func One(p survey.Prompt, response interface{}, opts ...survey.AskOpt) {
	err := survey.AskOne(p, response, opts...)

	if err != nil {
		if err == terminal.InterruptErr {
			events.App.Emit("cleanup")
			println("Aborted ✅")
			os.Exit(0)
		}
	}
}
