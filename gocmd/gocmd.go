package main

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/go-cmd/cmd"
)

func main() {
	testCmd := cmd.NewCmd("../c/seg")
	reader, writer := io.Pipe()
	go func() {
		_, _ = io.Copy(writer, bytes.NewBufferString("hello from the other side\n"))
		// close immediately
		_ = writer.Close()
	}()
	statusChan := testCmd.StartWithStdin(reader)

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {
			status := testCmd.Status()
			fmt.Println(status.Stdout)
			fmt.Println(status.Stderr)
		}
	}()

	// Block waiting for command to exit, be stopped, or be killed.
	// This does not unblock
	state := <-statusChan

	fmt.Println("statusChan final state")
	fmt.Println(state.Stdout)
	fmt.Println(state.Stderr)
}
