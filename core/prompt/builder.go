package prompt

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/Cre4T3Tiv3/gocmitra/core/config"
	"github.com/Cre4T3Tiv3/gocmitra/core/diff"
	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
)

type templateData struct {
	Diffs []diff.FileDiff
}

func Build(diffs []diff.FileDiff, cfg config.Config) string {
	if cfg.PromptTemplate != "" {
		tmpl, err := template.New("prompt").Parse(cfg.PromptTemplate)
		if err != nil {
			logger.Warn(fmt.Sprintf("Failed to parse custom prompt template: %v", err))
		} else {
			var buf bytes.Buffer
			if err := tmpl.Execute(&buf, templateData{Diffs: diffs}); err != nil {
				logger.Warn(fmt.Sprintf("Failed to execute prompt template: %v", err))
			} else {
				return buf.String()
			}
		}
	}

	var b strings.Builder
	if cfg.Instructions != "" {
		fmt.Fprintf(&b, "%s\n", cfg.Instructions)
	} else {
		fmt.Fprintf(&b, "Generate a commit message in %s style", cfg.Style)
		if cfg.Tone != "" {
			fmt.Fprintf(&b, " using a %s tone", cfg.Tone)
		}
		fmt.Fprintln(&b, " summarizing these changes:")
	}

	for _, d := range diffs {
		fmt.Fprintf(&b, "- %s: +%d -%d\n", d.File, d.Additions, d.Deletions)
	}

	return b.String()
}
