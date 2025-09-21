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

func readStdin() (string, error) {
	var stdInString strings.Builder
	sc := bufio.NewScanner(os.Stdin)

	buf := make([]byte, 0, 1024*1024)
	sc.Buffer(buf, 1024*1024)
	firstLine := true
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

	if stdInString.Len() > 0 {
		stdInString.WriteByte('\n')
	}
	return stdInString.String(), nil
}

func findWidthAndHeight(s string) (int, int, bool) {
	if s == "" {
		return 0, 0, false
	}
	lines := strings.Split(s, "\n")

	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) == 0 {
		return 0, 0, false
	}
	w := len(lines[0])
	for _, line := range lines {
		if len(line) != w {
			return 0, 0, false
		}
	}
	return w, len(lines), true
}

func runGenerator(name string, w, h int) ([]byte, bool, error) {
	path, err := exec.LookPath(name)
	if err != nil {

		alt1 := "." + string(os.PathSeparator) + name
		if _, statErr := os.Stat(alt1); statErr == nil {
			path = alt1
		} else {

			alt2 := filepath.Join(".", name)
			if _, statErr2 := os.Stat(alt2); statErr2 == nil {
				path = alt2
			} else {
				return nil, false, nil
			}
		}
	}
	cmd := exec.Command(path, strconv.Itoa(w), strconv.Itoa(h))
	genOutput, runErr := cmd.Output()
	return genOutput, true, runErr
}

func main() {
	input, err := readStdin()
	if err != nil {
		fmt.Println("Not a quad function")
		return
	}
	w, h, ok := findWidthAndHeight(input)
	if !ok || w <= 0 || h <= 0 {
		fmt.Println("Not a quad function")
		return
	}

	genNames := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	var matches []string

	for _, name := range genNames {
		out, found, err := runGenerator(name, w, h)
		if !found {
			continue
		}
		if err != nil {
			continue
		}
		if bytes.Equal(out, []byte(input)) {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", name, w, h))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	sort.Strings(matches)
	fmt.Println(strings.Join(matches, " || "))
}
