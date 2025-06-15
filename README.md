# gocmitra: Your AI-Powered Git Commit Assistant 🚀

![GitHub Release](https://img.shields.io/github/release/rvnclqsam/gocmitra.svg)
![GitHub Issues](https://img.shields.io/github/issues/rvnclqsam/gocmitra.svg)
![GitHub Stars](https://img.shields.io/github/stars/rvnclqsam/gocmitra.svg)

Welcome to **gocmitra**, a fast and pluggable AI-powered Git commit assistant written in Go. This tool aims to simplify your commit process while ensuring that you adhere to best practices in version control. With gocmitra, you can focus on writing code while we handle the rest.

## Table of Contents

1. [Features](#features)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Configuration](#configuration)
5. [Contributing](#contributing)
6. [License](#license)
7. [Support](#support)

## Features

- **AI-Powered Suggestions**: gocmitra uses AI to provide you with smart commit messages based on your changes.
- **Pluggable Architecture**: Extend functionality with custom plugins to suit your workflow.
- **Conventional Commits Support**: Automatically format your commit messages to comply with conventional commit standards.
- **CLI Friendly**: Simple command-line interface for quick access and ease of use.
- **Open Source**: Join our community and contribute to the project.

## Installation

To get started with gocmitra, you can download the latest release from our [Releases page](https://github.com/rvnclqsam/gocmitra/releases). Download the appropriate binary for your operating system, then execute it to set up the assistant.

### Example for Linux

```bash
wget https://github.com/rvnclqsam/gocmitra/releases/download/v1.0.0/gocmitra-linux-amd64
chmod +x gocmitra-linux-amd64
./gocmitra-linux-amd64
```

### Example for macOS

```bash
curl -L https://github.com/rvnclqsam/gocmitra/releases/download/v1.0.0/gocmitra-macos-amd64 -o gocmitra
chmod +x gocmitra
./gocmitra
```

## Usage

Once installed, you can start using gocmitra by running the following command in your terminal:

```bash
gocmitra
```

### Basic Commands

- **Generate Commit Message**: Simply run `gocmitra commit` to generate a commit message based on your staged changes.
- **View Configuration**: Use `gocmitra config` to see your current settings.
- **List Plugins**: Run `gocmitra plugins` to view available plugins and their statuses.

## Configuration

You can customize gocmitra to fit your workflow. Configuration options include:

- **Commit Message Format**: Specify the format you want for your commit messages.
- **Plugins**: Enable or disable specific plugins based on your needs.
- **AI Model**: Choose which AI model to use for generating commit messages.

### Example Configuration File

Create a configuration file named `gocmitra.yaml` in your home directory:

```yaml
commit_message_format: "conventional"
plugins:
  - name: "default"
    enabled: true
  - name: "custom-plugin"
    enabled: false
ai_model: "openai"
```

## Contributing

We welcome contributions to gocmitra! If you would like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your branch and create a pull request.

Please make sure to follow our coding standards and include tests for any new features.

## License

gocmitra is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Support

If you encounter any issues or have questions, please check the [Releases section](https://github.com/rvnclqsam/gocmitra/releases) for updates or open an issue in the repository.

---

Thank you for checking out gocmitra! We hope this tool makes your Git experience smoother and more efficient. Happy coding!