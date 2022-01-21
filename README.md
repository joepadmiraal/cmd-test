# cmd-test
Test repository to demonstrate blocking behavior with exec methods while using stdin.

First compile the C code.

    cd c
    gcc seg.c -o seg
    cd ..

After that you can compile and run the two Go code examples.

## go-cmd
Demonstrates blocking behaviour of StartWithStdin() from [go-cmd](https://github.com/go-cmd/cmd).

    cd gocmd
    go build
    ./gocmd

The result will be that the Go code keeps blocking on `state := <-statusChan`.

## exec
Demonstrates blocking behaviour of StdinPipe from the [exec](https://pkg.go.dev/os/exec) standard library package.

    cd exec
    go build
    ./exec

The result will be that the Go code will unblock as soon as the `seg` executable finishes.