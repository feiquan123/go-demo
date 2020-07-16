package bash

import (
	"testing"
)

func TestCommand(t *testing.T) {
	cmd := "ifconfig | grep l"
	output, err := Cmd(cmd)
	if err != nil {
		t.Error(err)
	}

	t.Log(output.String())
}
