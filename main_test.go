package main

import (
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

// ULIDの正規表現（26文字、Crockford Base32）
var ulidRegexp = regexp.MustCompile(`^[0-9A-HJKMNP-TV-Z]{26}$`)

func TestNoArg(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ulid := strings.TrimSpace(string(out))
	if !ulidRegexp.MatchString(ulid) {
		t.Errorf("output is not valid ULID: %q", ulid)
	}
}

func TestPrefix(t *testing.T) {
	prefix := "PRE_"
	cmd := exec.Command("go", "run", "main.go", prefix)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result := strings.TrimSpace(string(out))
	if !strings.HasPrefix(result, prefix) {
		t.Errorf("output does not start with prefix: %q", result)
	}
	ulid := strings.TrimPrefix(result, prefix)
	if !ulidRegexp.MatchString(ulid) {
		t.Errorf("ULID part is not valid: %q", ulid)
	}
}

func TestSuffix(t *testing.T) {
	suffix := "_SUF"
	cmd := exec.Command("go", "run", "main.go", "-s", suffix)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result := strings.TrimSpace(string(out))
	if !strings.HasSuffix(result, suffix) {
		t.Errorf("output does not end with suffix: %q", result)
	}
	ulid := strings.TrimSuffix(result, suffix)
	if !ulidRegexp.MatchString(ulid) {
		t.Errorf("ULID part is not valid: %q", ulid)
	}
}

func TestHelp(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "-h")
	out, _ := cmd.CombinedOutput()
	if !strings.Contains(string(out), "Usage:") {
		t.Errorf("help output missing Usage: %q", string(out))
	}
}
