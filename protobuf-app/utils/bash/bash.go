package bash

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

// Cmd is run command,then get output
func Cmd(command string) (output *bytes.Buffer, err error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd")
	default:
		cmd = exec.Command("sh")
	}
	input := new(bytes.Buffer)
	output = new(bytes.Buffer)
	outerr := new(bytes.Buffer)

	cmd.Stdin = input
	cmd.Stdout = output
	cmd.Stderr = outerr

	input.WriteString(command)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%v , %v", err, outerr.String())
	}
	return
}
