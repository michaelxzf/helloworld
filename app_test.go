package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDoublePositive(t *testing.T) {
	if got := double(2); got != 4 {
		t.Errorf("double(2) = %d, want 4", got)
	}
}

func TestDoubleZero(t *testing.T) {
	if got := double(0); got != 0 {
		t.Errorf("double(0) = %d, want 0", got)
	}
}

func TestDoubleNegative(t *testing.T) {
	if got := double(-3); got != -6 {
		t.Errorf("double(-3) = %d, want -6", got)
	}
}

func TestDividePositive(t *testing.T) {
	got, err := divide(6, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 2 {
		t.Errorf("divide(6, 3) = %v, want 2", got)
	}
}

func TestDivideNegative(t *testing.T) {
	got, err := divide(-6, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != -2 {
		t.Errorf("divide(-6, 3) = %v, want -2", got)
	}
}

func TestDivideFloatResult(t *testing.T) {
	got, err := divide(1, 4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 0.25 {
		t.Errorf("divide(1, 4) = %v, want 0.25", got)
	}
}

func TestDivideByZeroReturnsError(t *testing.T) {
	if _, err := divide(6, 0); err == nil {
		t.Error("divide(6, 0) should return an error")
	}
}

// workspaceDir creates a temporary directory inside the current working
// directory, which is the root readFile restricts access to.
func workspaceDir(t *testing.T) string {
	t.Helper()
	dir, err := os.MkdirTemp(".", ".tmp_test_")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })
	return dir
}

func TestReadFileReturnsContents(t *testing.T) {
	dir := workspaceDir(t)
	path := filepath.Join(dir, "example.txt")
	if err := os.WriteFile(path, []byte("hello world"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := readFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "hello world" {
		t.Errorf("readFile = %q, want %q", got, "hello world")
	}
}

func TestReadFileEmptyFile(t *testing.T) {
	dir := workspaceDir(t)
	path := filepath.Join(dir, "empty.txt")
	if err := os.WriteFile(path, []byte(""), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := readFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "" {
		t.Errorf("readFile = %q, want empty string", got)
	}
}

func TestReadFilePreservesMultilineContents(t *testing.T) {
	dir := workspaceDir(t)
	path := filepath.Join(dir, "lines.txt")
	if err := os.WriteFile(path, []byte("first\nsecond\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := readFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "first\nsecond\n" {
		t.Errorf("readFile = %q, want %q", got, "first\nsecond\n")
	}
}

func TestReadFileAcceptsRelativePathInWorkspace(t *testing.T) {
	got, err := readFile("main.go")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasPrefix(got, "package main") {
		t.Error(`readFile("main.go") should start with "package main"`)
	}
}

func TestReadFileRejectsPathTraversal(t *testing.T) {
	if _, err := readFile(filepath.Join("..", "secret.txt")); err == nil {
		t.Error("readFile should reject paths outside the workspace directory")
	}
}

func TestReadFileRejectsAbsolutePathOutsideWorkspace(t *testing.T) {
	if _, err := readFile("/etc/passwd"); err == nil {
		t.Error("readFile should reject absolute paths outside the workspace directory")
	}
}

func TestReadFileRejectsSymlinkEscape(t *testing.T) {
	dir := workspaceDir(t)
	link := filepath.Join(dir, "escape_link")
	if err := os.Symlink("/etc/passwd", link); err != nil {
		t.Skipf("cannot create symlink: %v", err)
	}
	if _, err := readFile(link); err == nil {
		t.Error("readFile should reject symlinks that escape the workspace directory")
	}
}
