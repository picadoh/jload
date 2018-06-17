package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <pid> <agent>\n", os.Args[0])
		os.Exit(1)
	}

	pid := os.Args[1]
	agent := os.Args[2]

	pidnum, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid <pid> value: %s\n", pid)
		os.Exit(1)
	}

	jvm := NewJvmProcess(pidnum)
	agentLoader := NewAgentLoader(jvm)

	if agentLoader.LoadAgent(agent) != nil {
		os.Exit(1)
	}
}
