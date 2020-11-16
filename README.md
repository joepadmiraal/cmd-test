# cmd-test

Demonstrates stdin behaviour of StartWithStdin().

First compile the C code.

    cd c
    gcc seg.c -o seg
    cd ..

After that you can compile and run the Go code.

    go build
    ./cmd-test

The result will be that the Go code keeps blocking on `state := <-statusChan`.
