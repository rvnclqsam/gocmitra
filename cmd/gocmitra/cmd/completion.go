package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate the autocompletion script for the specified shell",
	Long: `To load completions:

Bash:

  $ source <(gocmitra completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ gocmitra completion bash > /etc/bash_completion.d/gocmitra
  # macOS:
  $ gocmitra completion bash > /usr/local/etc/bash_completion.d/gocmitra

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it. Add the following to your ~/.zshrc:
  #   autoload -U compinit; compinit

  $ gocmitra completion zsh > "${fpath[1]}/_gocmitra"

Fish:

  $ gocmitra completion fish | source
  $ gocmitra completion fish > ~/.config/fish/completions/gocmitra.fish

PowerShell:

  PS> gocmitra completion powershell | Out-String | Invoke-Expression
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(_ *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			if err := rootCmd.GenBashCompletion(os.Stdout); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating bash completion: %v\n", err)
			}
		case "zsh":
			if err := rootCmd.GenZshCompletion(os.Stdout); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating zsh completion: %v\n", err)
			}
		case "fish":
			if err := rootCmd.GenFishCompletion(os.Stdout, true); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating fish completion: %v\n", err)
			}
		case "powershell":
			if err := rootCmd.GenPowerShellCompletionWithDesc(os.Stdout); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating PowerShell completion: %v\n", err)
			}

		}
	},
}
