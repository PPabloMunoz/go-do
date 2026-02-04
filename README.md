# 🌟 go-do

> A minimal, keyboard-driven todo list for your terminal ✨

A sleek command-line todo manager built with Go and
[Bubble Tea](https://github.com/charmbracelet/bubbletea).
Stay organized without leaving your terminal.

![go-do in action](https://via.placeholder.com/800x400?text=go-do+TUI+Screenshot)

---

## ✨ Features

- ⚡ **Lightning fast** - Built in Go, starts instantly
- 🎯 **Keyboard only** - Full control without touching your mouse
- 💾 **Persistent** - Your todos are saved automatically to `data.json`
- 🔄 **Undo support** - Accidentally deleted something? Get it back!
- 🎨 **Beautiful UI** - Clean, colorful terminal interface
- 🔒 **Safe** - Data is saved on exit (no manual saving needed)

---

## 🎮 Controls

| Key | Action |
|-----|--------|
| `↑` / `k` | Move cursor up |
| `↓` / `j` | Move cursor down |
| `space` | Toggle todo completion |
| `a` | Add new todo |
| `enter` | Confirm / Submit |
| `d` | Delete todo |
| `u` | Undo deletion |
| `esc` | Cancel |
| `q` | Quit & save |

---

## 🚀 Quick Start

### Build from source

```bash
git clone https://github.com/ppablomunoz/go-do
cd go-do
go build -o go-do .
```

### Run it

```bash
./go-do
```

---

## 📦 Requirements

- **Go** 1.25 or higher
- **Terminal** with true color support (24-bit colors)
- **macOS / Linux / Windows** (via WSL or terminal)

---

## 📝 License

MIT License - feel free to use and modify.

---

**Made with ❤️ using Go + Bubble Tea**
