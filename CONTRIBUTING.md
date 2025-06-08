# Contributing to GoC Mitra

👋 Thanks for considering a contribution! GoC Mitra is an open-source AI-powered CLI tool that helps you generate smart Git commit messages. This guide will help you get started contributing to the project.

---

## 📦 Project Structure

```

gocmitra/
├── .editorconfig               # Editor config rules
├── .gocmitra.json              # User or project config (generated)
├── .gitignore                  # Git exclusions
├── .golangci.yml               # Linter configuration
├── CHANGELOG.md                # Release changelog
├── CONTRIBUTING.md             # Contribution guidelines
├── LICENSE                     # Project license
├── README.md                   # Project overview and usage
├── gocmitra.zip                # Distribution archive
├── go.mod / go.sum             # Go modules
│
├── assets/                     # Static assets (e.g., preview images)
│   └── gocmitra_v0.1.0.jpg
│
├── cmd/                        # CLI entry point and command handlers
│   └── gocmitra/
│       ├── cmd/                # Subcommands: diff, config, completion, root
│       │   ├── completion.go
│       │   ├── config.go
│       │   ├── diff.go
│       │   ├── root.go
│       ├── gocmitra            # Compiled binary (ignored in repo)
│       └── main.go             # CLI launcher
│
├── core/                       # Business logic and internal services
│   ├── config/                 # Configuration loader and validator
│   │   └── config.go
│   ├── diff/                   # Git diff parser
│   │   └── parser.go
│   ├── llm/                    # LLM integration and provider clients
│   │   └── providers/
│   │       ├── claude.go
│   │       ├── client.go
│   │       ├── ollama.go
│   │       ├── openai.go
│   │       └── util.go
│   ├── logger/                 # Logging utility
│   │   └── logger.go
│   ├── prompt/                 # Prompt construction
│   │   └── builder.go
│   └── util/                   # General-purpose utilities
│       └── redact.go
│
├── docs/                       # Markdown guides
│   └── E2E-Guide.md
│
├── profiles/                   # Sample model configurations
│   ├── .gocmitra-claude3.json
│   ├── .gocmitra-llama3.json
│   └── .gocmitra-openai.json
│
└── .github/                    # GitHub-specific configuration
    └── workflows/
        └── ci.yml              # GitHub Actions CI workflow
````

---

## 🚀 Getting Started

* [🔄 Installation Guide](./README.md#installation)

---

## 🧪 Local Dev Tips

### Try different models

```bash
./gocmitra diff --model gpt-4o
./gocmitra diff --model claude-3-5-sonnet-20241022
./gocmitra diff --model llama3
```

> Tip: Use `--config ./profiles/.gocmitra-<model>.json` to test specific settings.

### Shell completion

```bash
# Bash
./gocmitra completion bash > ~/.gocmitra_completion && source ~/.gocmitra_completion

# Zsh
./gocmitra completion zsh > "${fpath[1]}/_gocmitra"

# Fish
./gocmitra completion fish | source
```

> Also works with PowerShell.

### Set environment variables

```bash
export OPENAI_API_KEY=sk-...
export CLAUDE_API_KEY=sk-ant-...
```

---

## ✅ Code Standards

* Use `gofmt` and `go vet`
* Keep standard output clean (only the commit message)
* Use `stderr` for logs, diagnostics, or errors
* Comment your logic clearly
* Prefer structured logging (via `logger` package)

---

## 📂 Profiles

Predefined model configs are in the `profiles/` directory:

* `.gocmitra-openai.json`
* `.gocmitra-claude3.json`
* `.gocmitra-llama3.json`

---

## 🧪 Testing

🧪 **Currently**: Manual testing via CLI
🛠️ **Coming Soon**: Unit and integration tests with mock LLM endpoints

---

## 💡 Suggestions

Want to contribute ideas like:

* New LLM providers?
* Shell UX enhancements?
* Plugin architecture?

Open an issue or [start a discussion](https://github.com/Cre4T3Tiv3/gocmitra/discussions)!

---

## 🤝 How to Contribute

1. Fork the repository
2. Create a new feature branch
3. Commit changes using semantic commit messages (`feat:`, `fix:`, `docs:`)
4. Push and submit a pull request

Please follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages.

---

## 📜 License

MIT © 2025 [@Cre4T3Tiv3](https://github.com/Cre4T3Tiv3)

---
