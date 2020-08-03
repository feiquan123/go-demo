package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

// RunCmd is test of command
func RunCmd() {
	fmt.Println("Run command `echo -n \"My first comes from golang\"`: ")
	cmd := exec.Command("echo", "-n", "My first comes from golang")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// useBufferIO := false // 是否使用自定义缓存
	useBufferIO := true // 是否使用自定义缓存
	if useBufferIO {
		var output bytes.Buffer
		for {
			tp := make([]byte, 5)
			n, err := stdout.Read(tp)
			if err != nil {
				if err == io.EOF { // 读取结束
					break
				} else {
					panic(err)
				}
			}
			if n > 0 {
				output.Write(tp[:n])
			}
		}
		fmt.Printf("使用自定义缓存读取：%s\n", output.String())
	} else {
		outputBuff := bufio.NewReader(stdout)
		output, _, err := outputBuff.ReadLine()
		if err != nil {
			panic(err)
		}
		fmt.Printf("使用缓存读取器：%s\n", string(output))
	}
}
