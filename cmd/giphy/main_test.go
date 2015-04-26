package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestUsageInstructions(t *testing.T) {
	out, _ := execGo("run", "main.go")

	if !strings.Contains(out, "Commands:") {
		t.Errorf(`expected list of commands`)
	}
}

func execGo(args ...string) (string, error) {
	out, err := exec.Command("go", args...).CombinedOutput()

	return string(out), err
}
