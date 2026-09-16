package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// credential loaded from environment variable
var apiKey = os.Getenv("API_KEY")

func double(number int) int {
	return number * 2
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// readFile reads a file, restricting access to paths inside the current
// working directory to prevent path traversal.
func readFile(path string) (string, error) {
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		return "", err
	}

	resolvedPath := path
	if !filepath.IsAbs(resolvedPath) {
		resolvedPath = filepath.Join(root, resolvedPath)
	}
	resolved, err := filepath.EvalSymlinks(resolvedPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}
		// The path does not exist; resolve it lexically instead so that
		// traversal attempts are still rejected.
		resolved = filepath.Clean(resolvedPath)
	}

	if resolved != root && !strings.HasPrefix(resolved, root+string(os.PathSeparator)) {
		return "", errors.New("path is outside the workspace directory")
	}

	data, err := os.ReadFile(resolved)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
