# End-to-End Guide: Using GoC Mitra in Your Dev Workflow

This guide walks you through how to set up and use **GoC Mitra** in a real-world Git environment, with examples for each supported model provider.

---

## 📋 Prerequisites

Before using GoC Mitra, ensure the following tools are installed:

- ✅ Go 1.21+  
- ✅ Git  
- ✅ (Optional) [Ollama](https://ollama.com) — for local LLaMA 3 support  
- ✅ (Optional) `golangci-lint` — for contributing or linting

---

## 🚀 1. Install the CLI

You can build GoC Mitra from source:

```bash
git clone https://github.com/Cre4T3Tiv3/gocmitra.git
cd gocmitra
go build -o cmd/gocmitra/gocmitra ./cmd/gocmitra
````

Then make it globally available:

```bash
sudo ln -s $(pwd)/cmd/gocmitra/gocmitra /usr/local/bin/gocmitra
```

> 💡 Confirm it's working with `gocmitra --version`

---

## ⚙️ 2. Create Your Config

Create a `.gocmitra.json` file in your project root or home directory.

Example for OpenAI:

```json
{
  "model": "gpt-4o",
  "endpoint": "https://api.openai.com/v1/chat/completions",
  "tone": "neutral",
  "style": "conventional",
  "instructions": "You are an assistant that writes concise Git commit messages."
}
```

You can also use prebuilt profiles from the `profiles/` directory:

* `profiles/.gocmitra-openai.json`
* `profiles/.gocmitra-claude3.json`
* `profiles/.gocmitra-llama3.json`

---

## 🔐 3. Set API Keys

For cloud LLM providers, export your API keys:

```bash
export OPENAI_API_KEY=sk-xxxx
export CLAUDE_API_KEY=sk-ant-xxxx
```

> 🧠 Ollama does not require an API key if running locally.

---

## 💻 4. Generate a Commit Message

To generate a commit message based on Git diffs:

```bash
gocmitra diff
```

For staged changes only:

```bash
gocmitra diff --staged
```

To override the model for just this run:

```bash
gocmitra diff --model llama3
```

📌 **Note:** The `--model` string must match the model name used in your config or by the LLM provider.

---

## 🔧 5. Set a Model Profile (Persistent)

You can permanently switch the default model:

```bash
gocmitra config set-model gpt-4o
```

This will generate or update `.gocmitra.json` accordingly.

---

## 🔁 6. Enable Shell Completion (Optional)

Autocompletion helps speed up CLI usage. Generate a script for your shell:

```bash
# Bash
gocmitra completion bash > ~/.gocmitra_completion && source ~/.gocmitra_completion

# Zsh
gocmitra completion zsh > "${fpath[1]}/_gocmitra"

# Fish
gocmitra completion fish | source
```

> 💡 Run `gocmitra completion --help` to explore more options (e.g. PowerShell)

---

## 🧪 7. Example Output

```bash
$ gocmitra diff --staged
[✓] Git diff collected
[✓] Loaded config
[✓] Using model: gpt-4o @ OpenAI
[✓] Sending prompt to model...

feat: add LLM integration for commit message generation
```

---

## 🧠 Supported Models

| Provider  | CLI Name   | Notes                     |
| --------- | ---------- | ------------------------- |
| OpenAI    | `gpt-4o`   | Requires OPENAI\_API\_KEY           |
| Anthropic | `claude-3` | Requires CLAUDE\_API\_KEY |
| Ollama    | `llama3`   | No API key required       |

---

## 🔁 Advanced Tips

* Keep config files under version control for team consistency
* Use CI to enforce commit message style (`style: "conventional"`)
* Automate commits via:

```bash
gocmitra diff | git commit -F -
```

---

## 🧩 Troubleshooting

| Symptom                        | Fix                                                    |
| ------------------------------ | ------------------------------------------------------ |
| `gocmitra` command not found   | Ensure it’s in your `$PATH`                            |
| Model errors / empty output    | Check your API key, model name, or local Ollama status |
| Confused by staged vs unstaged | Use `--staged` to limit to what you’ve added           |

---

## ✅ Done!

You’re ready to use GoC Mitra in your workflow. Try it after each code change and see how it improves your commit hygiene.

---