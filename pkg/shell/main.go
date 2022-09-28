package shell

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/Minnek-Digital-Studio/cominnek/pkg"
)

func getShell() string {
	os := runtime.GOOS
	if os == "windows" {
		return "powershell"
	}
	
	return "bash"
}

func Out(command string) (error, string, string) {
	var shellToUse = getShell()
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(shellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Start()
	err := cmd.Wait()

	return err, stdout.String(), stderr.String()
}

/*Execute a command and return the output*/
func ExecuteCommand(cmd string, print ...bool) string {
	var ignore bool
	err, out, errout := Out(cmd)

	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)

		pkg.App.Emit("cleanup")

		log.Fatal("")
	}

	if len(print) > 0 && !print[0] {
		ignore = true
	}

	if !ignore {
		fmt.Println(out)
		fmt.Println(errout)
	}

	return out
}
