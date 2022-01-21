package main

import (
	"io"
	"log"
	"os/exec"
)

type TestWriter struct{}

func (rt *TestWriter) Write(p []byte) (n int, err error) {
	log.Println(string(p))
	n = len(p)
	return
}

func main() {
	oscmd := exec.Command("../c/seg", "")
	tw := &TestWriter{}

	stdin, err := oscmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()

	oscmd.Stdout = tw
	oscmd.Start()

	io.WriteString(stdin, "Test line from GO")

	// Block waiting for command to exit, be stopped, or be killed.
	// This does unblock
	oscmd.Wait()
	log.Println("Finished")
}
