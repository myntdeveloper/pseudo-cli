# Pseudo CLI (psd)

A lightweight command-line tool to save, manage, and run shell command aliases — written in Go.

## Features

- Save long commands under short memorable names
- Run saved commands instantly from anywhere
- Tag and describe commands for easy organization
- Search and filter by tag
- Persistent local storage (`~/.psd/pseudonyms.json`)
- Colorized output for better readability

## Usage

### Quick Start

1. **Clone this repository**
   ```shell
   git clone https://github.com/myntdeveloper/pseudo-cli.git
   cd pseudo-cli
   ```

2. **Build the binary**
   ```shell
   go build -o psd.exe .
   ```

3. **Move to system PATH** (Windows)
   ```shell
   copy psd.exe C:\Windows\System32\psd.exe
   ```

4. **Save your first command**
   ```shell
   psd save djangorun "python manage.py runserver 8080" -t django -d "Start Django dev server"
   ```

### Commands

| Command         | Alias | Description                        |
|-----------------|---------|------------------------------------|
| `psd save`      | `s`   | Save a new command alias           |
| `psd run`       | `r`   | Run a saved command by name        |
| `psd list`      | `l`   | List all saved commands            |
| `psd show`      | —     | Show details of a command by name  |
| `psd remove`    | —     | Delete a command alias by name     |

### Flags for `save`

| Flag              | Short | Description                   | Default |
|-------------------|-------|-------------------------------|---------|
| `--tag`           | `-t`  | Tag for the command           | —       |
| `--description`   | `-d`  | Description for the command   | —       |

### Example Usage

Save a Go run command with a tag and description:
```shell
psd save run "go run main.go" -t go -d "Run main.go in your directory"
```

Save a Django server command:
```shell
psd save djangorun "python manage.py runserver 8080" -t django -d "Start Django dev server"
```

List all saved commands:
```shell
psd list
```

Filter by tag:
```shell
psd list --tag go
```

Run a saved command:
```shell
psd run djangorun
```

Show details of a command:
```shell
psd show run
```

Delete a command:
```shell
psd remove djangorun
```

## Directory Structure

```
pseudo-cli/
├── cmd/                    # Cobra CLI commands
│   ├── root.go
│   ├── save.go
│   ├── run.go
│   ├── list.go
│   ├── show.go
│   └── remove.go
├── internal/
│   ├── models/             # Data models
│   │   └── pseudonym.go
│   ├── runner/             # Business logic
│   │   └── runner.go
│   └── store/              # JSON storage layer
│       ├── store.go
│       └── errors.go
├── main.go
├── go.mod
└── README.md
```

## Stack

- **[Go](https://golang.org/)** — core language
- **[Cobra](https://github.com/spf13/cobra)** — CLI framework
- **[Viper](https://github.com/spf13/viper)** — configuration management
- **JSON** — local persistent storage (`~/.psd/pseudonyms.json`)

## Prerequisites

- Go 1.21 or newer

---

Built with ❤️ by **mynt**