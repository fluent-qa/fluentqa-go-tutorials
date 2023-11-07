package std

import (
	"os"
	"os/exec"
)

func OsCommandExecDemo() {
	cmd := exec.Command("echo", "hello world")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
