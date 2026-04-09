1# qrencode-go

👉 中文版本: [README.md](README.md)

---

## 📖 Table of Contents

* [Introduction](#introduction)
* [Features](#features)
* [Installation](#installation)
* [Usage](#usage)
* [Example](#example)
* [Project Structure](#project-structure)

---

## 📌 Introduction

`qrencode-go` is a minimal Go CLI tool that generates ASCII QR codes directly in the terminal.

Ideal for:

* Headless servers (SSH)
* DevOps workflows
* Quickly sharing links via QR code

---

## ✨ Features

* 🧾 Supports URL and arbitrary text
* 🖥️ Terminal-friendly QR output
* 📏 Compact rendering (half-height)
* 🎨 Inverted QR code for better visibility
* 🔲 Minimal border (no large padding)
* ⚡ Single binary, no runtime dependency
* 🌍 Cross-platform (Linux / macOS / Windows)

---

## 🚀 Installation

### Option 1: Download Release

👉 https://github.com/fscarmen/qrencode-go/releases

Download the appropriate binary:

```bash
qrencode-go-linux-amd64
qrencode-go-darwin-amd64
qrencode-go-windows-amd64.exe
```

---

### Option 2: Build from source

```bash
git clone https://github.com/fscarmen/qrencode-go.git
cd qrencode-go

go build -o qrencode-go .
```

---

## 🧱 Cross Compilation

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o qrencode-go-linux-amd64 .

# macOS
GOOS=darwin GOARCH=amd64 go build -o qrencode-go-darwin-amd64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o qrencode-go-windows-amd64.exe .
```

---

## ⚙️ Usage

```bash
qrencode-go [options] <text>
```

### Options

| Option     | Description                                 |
| ---------- | ------------------------------------------- |
| `-level`   | Error correction level (L/M/Q/H, default L) |
| `-qz`      | Quiet zone size (default 1)                 |
| `-compact` | Compact rendering (default true)            |

---

## 🔥 Recommended

```bash
qrencode-go -level L -qz 1 "https://example.com"
```

---

## 🖼️ Example Output

```
# ./qrencode-go-linux-amd64 https://github.com/fscarmen/qrencode-go
█▀▀▀▀▀▀▀█▀▀▀▀██▀█▀▀▀███▀▀▀▀▀▀▀█
█ █▀▀▀█ ██▀█ ▄ ▀▀▄█▄▄██ █▀▀▀█ █
█ █   █ ███  ▀▄▀█  █▀▄█ █   █ █
█ ▀▀▀▀▀ █▀▄▀▄▀▄ █ █ █ █ ▀▀▀▀▀ █
█▀▀█▀▀█▀█▄ ██ ▄▀▀ ▀▄▀ █▀█████▀█
█ ▀   ▄▀█ ▀▄█▀ ▄ ▀ ▄▄▀▄█  █ ▄██
█▄▄ ▀▄ ▀▀ ▀▀██▀ ▄▀▄█  ▄  █▄██ █
█▄▄▄  ▄▀▀█▀▄█▄█  ▄█▄▀▀▀ ▄ ▄ ▄ █
█ ▄▄ █▀▀█▀▄▀▀ ▄▄ ▄██ ▀▀█▄▄▀▄█▄█
█  ▄█▀▄▀▄█▄▀▄▀▀▀ ▀█▄ ████ ▀▄▄▀█
█  ▄▀ ▄▀▀▄▄▄ █ █▄ ▀ ▄▀▀ ▀▀  ▀██
█▀▀▀▀▀▀▀█▄█▄▀▄ █ ▀▄▄▀ █▀█  ████
█ █▀▀▀█ █▀██   █▀▀▄█  ▀▀▀ ▀█▄▀█
█ █   █ █▄▀█▀▀ ▄  █▀█▀ █▀▀█▀▀ █
█ ▀▀▀▀▀ █  ▄▄██▀  ██▀▀█ ▄▀▄▄█▄█
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
```

---

## 📁 Project Structure

```
qrencode-go/
├── main.go
├── render.go
├── go.mod
└── .github/workflows/release.yml
```

---

## 📄 License

MIT