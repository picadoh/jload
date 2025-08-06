[![CI](https://github.com/picadoh/jload/actions/workflows/ci.yml/badge.svg)](https://github.com/picadoh/jload/actions/workflows/ci.yml)

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

```shell
go build
```

To run the tests:

```shell
go test -v ./...
```

# Installing

If you have `go` available in your system, you may just:

```
go install github.com/picadoh/jload@latest
```

# Running

```shell
jload <pid> <agent>
```

*Example*

```shell
jload 1234 /path/to/my/myagent.jar
```

*Example with _pgrep_:*

```shell
pgrep -f .*myapp.* | xargs -I % ./jload % /path/to/my/myagent.jar
```

*Or run the provided example:*

Note: Requires Java 21+ and Go 1.23+.

```shell
./example/run.sh
```

# How to Contribute

Please take a look at [CONTRIBUTING](CONTRIBUTING.md) for instructions.
