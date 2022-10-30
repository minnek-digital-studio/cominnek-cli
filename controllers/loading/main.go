package loading

import (
	"time"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/logger"
	"github.com/briandowns/spinner"
)

var loader = spinner.New(spinner.CharSets[23], 100*time.Millisecond)
var t1 time.Time

func Start(prefix string) {
	if config.Public.Logs {
		t1 = time.Now()
		logger.Warning(t1.Format("15:04:05"), prefix)
		return
	}
	loader.Start()
	loader.Prefix = prefix
}

func Update(prefix string) {
	if config.Public.Logs {
		return
	}
	loader.Prefix = prefix
}

func Stop() {
	if config.Public.Logs {
		diff := time.Since(t1)
		logger.Success(time.Now().Format("15:04:05"), "Finished in", diff.String())
		return
	}
	loader.Stop()
}
