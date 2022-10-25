package ask

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

func One(p survey.Prompt, responce interface{}, opts ...survey.AskOpt) {
	err := survey.AskOne(p, responce, opts...)

	if err != nil {
		if err == terminal.InterruptErr {
			println("Aborted ✅")
			os.Exit(0)
		}
	}
}
