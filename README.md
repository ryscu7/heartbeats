# 💖 Heartbeats

A tiny Go-based reminder app that sends periodic heartbeats - toast notifications on Windows. Originally made for my girlfriend, but you're free to use, modify or build your own version of it.

---

## ✨ What It Does

- Sends a sweet toast notification every 10 minutes ( this can't be modified yet but i'm working on it )
- Shows a morning message at 11 AM and a night message at 11 PM
- Runs quietly in the background
- Small resource footprint ( ~1% CPU and ~1MB RAM Usage)

## 📥 Quick Install ( Windows )

You can install it directly using Powershell:

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force
iex (Invoke-WebRequest -UseBasicParsing "https://raw.githubusercontent.com/ryscu7/heartbeats/refs/heads/main/install.ps1").Content
```

This script:

- Downloads the latest `heartbeats.exe` from GitHub Releases.
- Adds it to the Startup Folder so it runs when your PC starts

> [!NOTE]
> To see the startup folder, Press `Windows Key + R` and type `shell:startup` and press `Enter`.
