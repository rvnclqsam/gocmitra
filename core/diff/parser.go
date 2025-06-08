package diff

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
)

// FileDiff contains statistics about changes in a file.
type FileDiff struct {
	File      string
	Additions int
	Deletions int
}

// Parse reads a unified git diff and returns simple statistics for each file.
func Parse(diffText string) []FileDiff {
	scanner := bufio.NewScanner(strings.NewReader(diffText))
	var diffs []FileDiff
	var current *FileDiff

	logger.Info("Parsing git diff content")

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "diff --git") {
			parts := strings.Split(line, " ")
			if len(parts) >= 4 {
				file := strings.TrimPrefix(parts[3], "b/")
				diffs = append(diffs, FileDiff{File: file})
				current = &diffs[len(diffs)-1]
			} else {
				logger.Warn(fmt.Sprintf("Malformed diff header: %s", line))
			}
			continue
		}

		if current == nil {
			continue
		}

		if strings.HasPrefix(line, "+++") || strings.HasPrefix(line, "---") {
			continue
		}

		if strings.HasPrefix(line, "+") {
			current.Additions++
		} else if strings.HasPrefix(line, "-") {
			current.Deletions++
		}
	}

	return diffs
}
