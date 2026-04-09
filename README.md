# qrencode-go

👉 English version: [README_en.md](README_en.md)

---

## 📖 目录

* [项目介绍](#项目介绍)
* [功能特性](#功能特性)
* [安装](#安装)
* [使用方法](#使用方法)
* [示例](#示例)
* [项目结构](#项目结构)

---

## 📌 项目介绍

`qrencode-go` 是一个极简的 Go 命令行工具，用于在终端直接生成 ASCII 二维码。

适用于：

* 无图形界面的服务器（SSH）
* DevOps / 运维场景
* 快速分享链接（扫码访问）

---

## ✨ 功能特性

* 🧾 支持 URL / 任意文本
* 🖥️ 终端直接显示二维码
* 📏 半高压缩（更紧凑）
* 🎨 二维码本体黑白反转（更清晰）
* 🔲 可控边框（避免大白边）
* ⚡ 单文件二进制，无需依赖
* 🌍 支持 Linux / macOS / Windows

---

## 🚀 安装

### 方法一：下载 Release

前往：

👉 https://github.com/你的用户名/qrencode-go/releases

下载对应系统版本：

```bash
qrencode-go-linux-amd64
qrencode-go-darwin-amd64
qrencode-go-windows-amd64.exe
```

---

### 方法二：源码构建

```bash
git clone https://github.com/你的用户名/qrencode-go.git
cd qrencode-go

go build -o qrencode-go .
```

---

## 🧱 交叉编译（多平台）

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o qrencode-go-linux-amd64 .

# macOS
GOOS=darwin GOARCH=amd64 go build -o qrencode-go-darwin-amd64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o qrencode-go-windows-amd64.exe .
```

---

## ⚙️ 使用方法

```bash
qrencode-go [options] <text>
```

### 参数说明

| 参数         | 说明                       |
| ---------- | ------------------------ |
| `-level`   | 纠错等级：L / M / Q / H（默认 L） |
| `-qz`      | 边框大小（默认 1）               |
| `-compact` | 半高压缩输出（默认 true）          |

---

## 🔥 推荐用法（最小尺寸）

```bash
qrencode-go -level L -qz 1 "https://example.com"
```

---

## 🖼️ 示例

```
██████████████████████████
██ ▄▄▄▄▄ ██ ▀█ ▄▄▄▄▄ ██
██ █   █ █▀▄█ █   █ ██
██ ▀▀▀▀▀ ██▄█ ▀▀▀▀▀ ██
████████████████████████
```

---

## 📁 项目结构

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
