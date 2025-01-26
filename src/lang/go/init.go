package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <year> <day> <part>\n", os.Args[0])
		os.Exit(1)
	}

	year, day, part := os.Args[1], os.Args[2], os.Args[3]

	// Download input if needed
	inputPath, err := DownloadInput(year, day)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading input: %v\n", err)
		os.Exit(1)
	}

	// Open input file
	input, err := os.Open(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input: %v\n", err)
		os.Exit(1)
	}
	defer input.Close()

	// Construct solution path
	solutionPath := filepath.Join("src", "year", year, day, fmt.Sprintf("part%s.go", part))

	// Run the solution with input piped to stdin
	cmd := exec.Command("go", "run", solutionPath)
	cmd.Stdin = input
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running solution: %v\n", err)
		os.Exit(1)
	}
}

func DownloadInput(year, day string) (string, error) {
	// Read session cookie from file
	sessionBytes, err := os.ReadFile("session.cookie")
	if err != nil {
		return "", fmt.Errorf("failed to read session.cookie file: %v", err)
	}
	sessionCookie := strings.TrimSpace(string(sessionBytes))
	if sessionCookie == "" {
		return "", fmt.Errorf("session.cookie file is empty")
	}

	inputPath := filepath.Join("src", "year", year, day, "input.txt")
	if _, err := os.Stat(inputPath); err == nil {
		return inputPath, nil // File already exists
	}

	// Ensure directory exists
	dir := filepath.Dir(inputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	// Download input
	var dayNum int
	must(fmt.Sscanf(day, "%02d", &dayNum))
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%d/input", year, dayNum)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download input: %s", resp.Status)
	}

	file, err := os.Create(inputPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return inputPath, err
}

func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
