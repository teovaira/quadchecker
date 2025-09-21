package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// readStdin reads all text from standard input and returns it as a string.
// It preserves line breaks and ensures the final string ends with a newline if needed.
func readStdin() (string, error) {
	var stdInString strings.Builder
	sc := bufio.NewScanner(os.Stdin)

	// Increase scanner buffer size to handle large inputs
	buf := make([]byte, 0, 1024*1024)
	sc.Buffer(buf, 1024*1024)
	firstLine := true

	// Read input line by line
	for sc.Scan() {
		if !firstLine {
			stdInString.WriteByte('\n')
		}
		stdInString.WriteString(sc.Text())
		firstLine = false
	}
	if err := sc.Err(); err != nil {
		return "", err
	}

	// Ensure output ends with a newline if not empty
	if stdInString.Len() > 0 {
		stdInString.WriteByte('\n')
	}
	return stdInString.String(), nil
}

// findWidthAndHeight checks if the given string forms a valid rectangle.
// It returns width, height, and a boolean indicating if the check passed.
func findWidthAndHeight(s string) (int, int, bool) {
	if s == "" {
		return 0, 0, false
	}
	lines := strings.Split(s, "\n")

	// Remove empty last line if present
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) == 0 {
		return 0, 0, false
	}

	// All lines must have the same width
	w := len(lines[0])
	for _, line := range lines {
		if len(line) != w {
			return 0, 0, false
		}
	}

	return w, len(lines), true
}

// runGenerator executes one of the quad generators with the given width and height.
// It returns the generated output, whether the generator was found, and any run error.
func runGenerator(name string, w, h int) ([]byte, bool, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		// If not found in PATH, try local directory
		alt1 := "." + string(os.PathSeparator) + name
		if _, statErr := os.Stat(alt1); statErr == nil {
			path = alt1
		} else {
			alt2 := filepath.Join(".", name)
			if _, statErr2 := os.Stat(alt2); statErr2 == nil {
				path = alt2
			} else {
				// Generator not found
				return nil, false, nil
			}
		}
	}

	// Run the generator with w and h as arguments
	cmd := exec.Command(path, strconv.Itoa(w), strconv.Itoa(h))
	genOutput, runErr := cmd.Output()
	return genOutput, true, runErr
}

func main() {
	// Read the input shape
	input, err := readStdin()
	if err != nil {
		fmt.Println("Not a quad function")
		return
	}

	// Find dimensions of the input
	w, h, ok := findWidthAndHeight(input)
	if !ok || w <= 0 || h <= 0 {
		fmt.Println("Not a quad function")
		return
	}

	// List of possible quad generators
	genNames := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	var matchedGenerators []string

	// Compare input against each generator
	for _, name := range genNames {
		genOutput, found, err := runGenerator(name, w, h)
		if !found || err != nil {
			continue
		}
		// If the output matches the input, record the match
		if bytes.Equal(genOutput, []byte(input)) {
			matchedGenerators = append(matchedGenerators, fmt.Sprintf("[%s] [%d] [%d]", name, w, h))
		}
	}

	// Print results
	if len(matchedGenerators) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	sort.Strings(matchedGenerators)
	fmt.Println(strings.Join(matchedGenerators, " || "))
}
