package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"adventofcode2023/internal/day1"
)

func fetchInput(day int) (string, error) {
    cacheDir := "inputs"
    cacheFile := filepath.Join(cacheDir, fmt.Sprintf("day%d.txt", day))

    if content, err := os.ReadFile(cacheFile); err == nil {
        return string(content), nil
    }

    sessionCookie := os.Getenv("AOC_SESSION")
    if sessionCookie == "" {
        return "", fmt.Errorf("AOC_SESSION environment variable not set")
    }

    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", fmt.Errorf("failed to create request: %w", err)
    }

    req.AddCookie(&http.Cookie{
        Name:  "session",
        Value: sessionCookie,
    })

    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("failed to fetch input: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("failed to read response body: %w", err)
    }

    if err := os.MkdirAll(cacheDir, 0755); err != nil {
        return "", fmt.Errorf("failed to create cache directory: %w", err)
    }

    if err := os.WriteFile(cacheFile, body, 0644); err != nil {
        return "", fmt.Errorf("failed to cache input: %w", err)
    }

    return string(body), nil
}

func runPart(day, part int, fn func(string) (int, error), input string) error {
    start := time.Now()
    result, err := fn(input)
    if err != nil {
        return fmt.Errorf("day %d part %d failed: %v", day, part, err)
    }
    duration := time.Since(start)
    fmt.Printf("Day %d Part %d Result: %d (took %v)\n", day, part, result, duration)
    return nil
}

func runDay(day int) error {
    start := time.Now()
    input, err := fetchInput(day)
    if err != nil {
        return fmt.Errorf("failed to fetch input for day %d: %v", day, err)
    }

    switch day {
    case 1:
        if err := runPart(1, 1, day1.PartOne, input); err != nil {
            return err
        }
        if err := runPart(1, 2, day1.PartTwo, input); err != nil {
            return err
        }
    default:
        return fmt.Errorf("day %d not implemented yet", day)
    }

    fmt.Printf("Total time: %v\n", time.Since(start))
    return nil
}

func main() {
    day := 1
    if len(os.Args) > 1 {
        var err error
        day, err = strconv.Atoi(os.Args[1])
        if err != nil {
            log.Fatalf("Invalid day specified: %v", err)
        }
    }

    if err := runDay(day); err != nil {
        log.Fatalf("Error running day %d: %v", day, err)
    }
}