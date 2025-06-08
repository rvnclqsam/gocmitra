<p align="center">
  <a href="https://github.com/Cre4T3Tiv3/gocmitra" target="_blank">
    <img src="https://raw.githubusercontent.com/Cre4T3Tiv3/gocmitra/main/docs/assets/gocmitra_v0.1.0.jpg" alt="GoC Mitra social preview" width="640"/>
  </a>
</p>

<p align="center"><em>
🧠 An AI-powered Git commit assistant that analyzes your code changes and generates smart, conventional commits — right from your terminal.
</em></p>

<p align="center">
  <a href="https://github.com/Cre4T3Tiv3/gocmitra/actions/workflows/ci.yml?query=branch%3Amain" target="_blank">
    <img src="https://github.com/Cre4T3Tiv3/gocmitra/actions/workflows/ci.yml/badge.svg?branch=main" alt="CI">
  </a>
  <a href="https://www.go.dev/dl/" target="_blank">
    <img src="https://img.shields.io/badge/Go-1.21+-blue" alt="Go Version">
  </a>
  <a href="https://github.com/Cre4T3Tiv3/gocmitra/tags" target="_blank">
    <img src="https://img.shields.io/github/v/tag/Cre4T3Tiv3/gocmitra" alt="Latest Tag">
  </a>
  <a href="https://github.com/Cre4T3Tiv3/gocmitra/blob/main/LICENSE" target="_blank">
    <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License: MIT">
  </a>
  <a href="https://github.com/Cre4T3Tiv3/gocmitra/stargazers" target="_blank">
    <img src="https://img.shields.io/github/stars/Cre4T3Tiv3/gocmitra?style=social" alt="GitHub Stars">
  </a>
  <a href="#contributing" target="_blank">
    <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg" alt="Contributions welcome">
  </a>
</p>

---

## 🚀 About

**GoC Mitra** is your Git commit assistant – a developer-friendly CLI tool powered by AI that turns your Git diffs into clear, conventional, single-line commit messages. Built in Go and designed for developers who live in the terminal.

Supports:
- ✅ OpenAI (GPT-4o, GPT-3.5, etc.)
- ✅ Claude (3, 3.5)
- ✅ Ollama (LLaMA 3 and other local models)

---

## 🔧 Installation

> 💡 Requires Go 1.21 or higher.

```bash
git clone https://github.com/Cre4T3Tiv3/gocmitra.git
cd gocmitra
go build -o cmd/gocmitra/gocmitra ./cmd/gocmitra
````

Add to your `$PATH` or symlink to a local bin:

```bash
sudo ln -s $(pwd)/cmd/gocmitra/gocmitra /usr/local/bin/gocmitra
```

> ✅ This places the compiled binary in `cmd/gocmitra/`, keeping the root of your repo clean.

---

## 🛠️ Configuration

GoC Mitra looks for a `.gocmitra.json` config file in your project or home directory.

This file is **automatically generated** when you select a model using:

```bash
gocmitra config set-model gpt-4o
```

The command copies a predefined config from `profiles/.gocmitra-<model>.json` into your local `.gocmitra.json`.

You can also pass a custom path using the `--config` flag.

### Example config:

```json
{
  "model": "gpt-4o",
  "endpoint": "https://api.openai.com/v1/chat/completions",
  "tone": "neutral",
  "style": "conventional",
  "instructions": "You are an assistant that writes commit messages..."
}
```

Prebuilt model configs are available in the `profiles/` directory.

---

## ✨ Usage

Use the `diff` command to generate a commit message from your Git changes.

### 🔍 Examples

```bash
# Generate a commit message from all unstaged and staged changes
gocmitra diff
```

```bash
# Generate a commit message from only staged changes
gocmitra diff --staged
```

```bash
# Use a custom config file
gocmitra diff --config ~/.gocmitra.json
```

```bash
# Override the model temporarily
gocmitra diff --model claude-3-5-sonnet-20241022
```

> ℹ️ **Note:** The `--staged` flag corresponds to `git diff --staged`, which only includes changes you've staged with `git add`.

To change the default model and regenerate the config file:

```bash
gocmitra config set-model gpt-4o
```

---

## 🔁 Shell Completion

GoC Mitra supports shell autocompletion for Bash, Zsh, Fish, and PowerShell.

### Bash

```bash
gocmitra completion bash > /etc/bash_completion.d/gocmitra
# or for user-level
gocmitra completion bash > ~/.gocmitra_completion && source ~/.gocmitra_completion
```

### Zsh

```bash
gocmitra completion zsh > "${fpath[1]}/_gocmitra"
```

### Fish

```bash
gocmitra completion fish | source
```

> 💡 Run `gocmitra completion --help` to view all shell options.

---

## 🧠 Supported LLMs

| Provider  | Model Name                   | Notes                     |
| --------- | ---------------------------- | ------------------------- |
| OpenAI    | `gpt-4o`                     | Fast, reliable            |
| Anthropic | `claude-3-5-sonnet-20241022` | Requires `CLAUDE_API_KEY` |
| Ollama    | `llama3`                     | Local, privacy-friendly   |

Environment variables required:

* `OPENAI_API_KEY`
* `CLAUDE_API_KEY`

---

## 📚 Docs

* [🔄 E2E Guide](docs/E2E-Guide.md) – How to integrate GoC Mitra in your dev flow
* [🛠️ Contributing](CONTRIBUTING.md) – Setup, coding standards, and dev notes

---

## 🧑‍💻 Contributing

We welcome contributions! Please read [CONTRIBUTING.md](CONTRIBUTING.md) before submitting PRs or opening issues.

---

## 🧾 License

MIT © 2025 [@Cre4T3Tiv3](https://github.com/Cre4T3Tiv3)

---

## 📣 Public Beta

GoC Mitra is now in **public beta**. It's stable, actively maintained, and ready for feedback. Try it in your dev workflow and help us improve it.

---
