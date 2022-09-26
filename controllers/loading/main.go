package loading

import (
	"time"

	"github.com/briandowns/spinner"
)

var loader = spinner.New(spinner.CharSets[23], 100*time.Millisecond)

func Start(prefix string) {
	loader.Start()
	loader.Prefix = prefix
}

func Update(prefix string) {
	loader.Prefix = prefix
}

func Stop() {
	loader.Stop()
}