package shell

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
)

func getShell() string {
	os := runtime.GOOS
	if os == "windows" {
		return "powershell"
	}

	return "bash"
}

var shellToUse = getShell()

func Out(command string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(shellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Start()
	err := cmd.Wait()

	return stdout.String(), stderr.String(), err
}

func OutLive(command string) {
	cmd := exec.Command(shellToUse, "-c", command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	err = cmd.Start()
	loading.Start("Running command...")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(stdout)
	loading.Stop()
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	cmd.Wait()
}

/*Execute a command and return the output*/
func ExecuteCommand(cmd string, print ...bool) string {
	var ignore bool
	out, errout, err := Out(cmd)

	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)

		events.App.Emit("cleanup")

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
