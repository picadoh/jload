[![Build Status](https://travis-ci.org/picadoh/jload.svg?branch=master)](https://travis-ci.org/picadoh/jload) 

JLoad provides the ability to load an agent dynamically into a running VM.

# How it works

Whenever a JVM receives a SIGQUIT signal, if a .attach_pid<pid> file is present,
the JVM will open a server socket connection using Unix Domain Sockets. Through
this socket a client can dynamically load a JAR into the JVM that will run there.

A typical use of this may be loading a monitoring agent into a JVM that will
expose JMX MBeans in a certain format 
(e.g. [Prometheus JMX Exporter Agent](https://github.com/prometheus/jmx_exporter)).

# Building and Testing

To build the binary on the current platform:

    $ go build

To run the tests:

    $ go test

# Docker

To build a docker image:

    $ docker build -t <image_name> .

To run from the [existing Docker image](https://hub.docker.com/r/picadoh/jload/):

    $ docker run --rm picadoh/jload <pid> <agent>

# Cross-Platform Compilation

A Makefile is available to make things easier. Just run the following command to generate binaries for multiple platforms.

    $ make

It will generate an output like the following:

    GOOS=linux GOARCH=amd64 go build -o 'jload-linux-amd64'
    GOOS=windows GOARCH=amd64 go build -o 'jload-windows-amd64.exe'
    GOOS=darwin GOARCH=amd64 go build -o 'jload-darwin-amd64'

# Installing

If you have `go` available in your system, you may just:

    go get github.com/picadoh/jload
    go build github.com/picadoh/jload
    go install github.com/picadoh/jload

Packages for multiple platforms are still not available elsewere so for now it must be built from source.

# Running

    jload <pid> <agent>

*Example*

    $ jload 1234 /path/to/my/myagent.jar

*Example with _pgrep_:*

    $ pgrep -f .*myapp.* | xargs -I % ./jload % /path/to/my/myagent.jar

# How to Contribute

Please take a look at [CONTRIBUTING](CONTRIBUTING.md) for instructions.
