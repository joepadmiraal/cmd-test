package main

import (
	"fmt"
	"time"

	"github.com/go-cmd/cmd"
)

//Dummy reader which does nothing
type TestReader struct {
}

//Read get's called by go-cmd
func (rt *TestReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func NewTestReader() *TestReader {
	rt := TestReader{}
	return &rt
}

func main() {
	testCmd := cmd.NewCmd("c/seg")
	rt := NewTestReader()
	statusChan := testCmd.StartWithStdin(rt)

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
