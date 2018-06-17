package main

import (
	"fmt"
	"os"
	"testing"
)

func TestJvm(t *testing.T) {
	expectedPid := 1234
	expectedCwd := fmt.Sprintf("/proc/%d/cwd", 1234)
	if _, err := os.Stat(expectedCwd); os.IsNotExist(err) {
		expectedCwd = os.TempDir()
	}
	expectedAttachSocketAddr := os.TempDir() + string(os.PathSeparator) + ".java_pid1234"

	subject := NewJvmProcess(1234)

	if pid := subject.Pid(); pid != expectedPid {
		t.Fatalf("Unexpected PID:\nGot:\t\t%d\nExpected:\t%d\n", pid, expectedPid)
	}

	if cwd := subject.Cwd(); cwd != expectedCwd {
		t.Fatalf("Unexpected cwd:\nGot:\t\t%s\nExpected:\t%s\n", cwd, expectedCwd)
	}

	if attachSocketAddr := subject.AttachSocketAddr(); attachSocketAddr != expectedAttachSocketAddr {
		t.Fatalf("Unexpected cwd:\nGot:\t\t%s\nExpected:\t%s\n", attachSocketAddr, expectedAttachSocketAddr)
	}
}
