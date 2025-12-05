# Installer - Quick Installation Tool for GitHub Release Binaries

An intelligent HTTP service for quickly installing pre-compiled binaries from GitHub Releases. It automatically detects your system platform (OS and architecture), selects the appropriate asset, downloads, extracts, and optionally moves it to your system PATH.

[![GoDev](https://img.shields.io/static/v1?label=godoc&message=reference&color=00add8)](https://pkg.go.dev/github.com/jpillora/installer)
[![CI](https://github.com/jpillora/installer/workflows/CI/badge.svg)](https://github.com/jpillora/installer/actions?workflow=CI)

## Features

- 🚀 **One-Command Installation**: Install GitHub Release binaries directly via `curl`
- 🔍 **Smart Search**: Automatically searches the web for repositories when not found
- 🎯 **Platform Detection**: Auto-detects OS (Linux, macOS, BSD, etc.) and architecture (amd64, arm64, 386, etc.)
- 📦 **Multiple Format Support**: Supports `.zip`, `.tar.gz`, `.tar.bz2`, `.tar.xz`, `.gz`, `.bz2` and more
- 🎭 **Multi-Binary Support**: Install multiple binaries from a single release
- ☁️ **Cloudflare Workers**: Deploy as Cloudflare Workers for global edge acceleration
- 🔒 **Private Repository Support**: Access private repos via `GITHUB_TOKEN`

## Quick Start

### Basic Usage

```bash
# Install a repository (to current directory)
curl https://your-domain.workers.dev/user/repo | bash

# Install to /usr/local/bin/ (use the ! suffix)
curl https://your-domain.workers.dev/user/repo! | bash

# Install a specific version
curl https://your-domain.workers.dev/user/repo@v1.2.3! | bash

# Use repository name only (will attempt search)
curl https://your-domain.workers.dev/micro! | bash
```

### Install Multiple Binaries

```bash
# Install both client and server from one release
curl https://your-domain.workers.dev/xtaci/kcptun!?mp=client,server | bash
```

## Usage

### URL Format

```
https://your-domain.workers.dev/<user>/<repo>@<release>!?<parameters>
```

### Path Parameters

- `user` - GitHub username (optional, uses default if omitted)
- `repo` - Repository name (required)
- `release` - Release version (optional, defaults to `latest`)
- `!` - Install to `/usr/local/bin/` (optional, without it installs to current directory)

### Query Parameters

| Parameter | Description | Example |
|-----------|-------------|---------|
| `mp` | Install multiple programs (comma-separated) | `?mp=client,server` |
| `as` | Rename the binary | `?as=rg` |
| `os` | Override OS detection | `?os=linux` |
| `arch` | Override architecture detection | `?arch=arm64` |
| `type` | Output format (`script`, `json`, `text`) | `?type=json` |
| `select` | Select specific asset variant | `?select=musl` |
| `insecure` | Skip SSL verification | `?insecure=1` |

## Examples

### Install micro Editor

```bash
curl https://your-domain.workers.dev/zyedidia/micro! | bash
micro --version
```

### Install Caddy Web Server

```bash
curl https://your-domain.workers.dev/caddyserver/caddy@v2.7.6! | bash
caddy version
```

### Install and Rename ripgrep

```bash
curl https://your-domain.workers.dev/BurntSushi/ripgrep!?as=rg | bash
rg --version
```

### Install Multiple Binaries (kcptun)

```bash
curl https://your-domain.workers.dev/xtaci/kcptun!?mp=client,server | bash
mv /usr/local/bin/client /usr/local/bin/kcptun-client
mv /usr/local/bin/server /usr/local/bin/kcptun-server
```

### Install shadowsocks-rust

```bash
curl https://your-domain.workers.dev/shadowsocks/shadowsocks-rust!?mp=sslocal,ssserver | bash
```

### Use in Docker

```dockerfile
FROM ubuntu:22.04
RUN apt-get update && apt-get install -y curl
RUN curl https://your-domain.workers.dev/zyedidia/micro! | bash
CMD ["/bin/bash"]
```

### Debug Installation

```bash
# View script without executing
curl https://your-domain.workers.dev/jpillora/serve

# Run with debug output
DEBUG=1 curl https://your-domain.workers.dev/jpillora/serve! | bash

# Get JSON asset information
curl "https://your-domain.workers.dev/jpillora/serve?type=json"
```

## Deployment

### Cloudflare Workers (Recommended)

#### Prerequisites

```bash
# Install Wrangler CLI
npm install -g wrangler

# Install TinyGo
brew install tinygo  # macOS

# Install quicktemplate
go install github.com/valyala/quicktemplate/qtc@latest
```

#### Configure

Edit `wrangler.jsonc`:

```jsonc
{
    "vars": {
        "DEFAULT_USER": "your-github-username",
        "GITHUB_TOKEN": "ghp_xxxxx"  // Optional
    }
}
```

#### Deploy

```bash
make build
wrangler deploy
```

### Fly.io

```bash
fly auth login
fly deploy
```

### Run Locally

```bash
# Install dependencies
go mod download

# Run
go run main.go

# Or with custom config
PORT=8080 USER=myusername go run main.go
```

#### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Listen port | `3000` |
| `USER` / `DEFAULT_USER` | Default GitHub user | `cxjava` |
| `GITHUB_TOKEN` | GitHub API token (optional) | - |
| `FORCE_USER` | Lock to specific user (optional) | - |
| `FORCE_REPO` | Lock to specific repo (optional) | - |

## Supported Platforms

### Operating Systems
Linux, macOS, FreeBSD, OpenBSD, NetBSD, DragonFly BSD, Android, Solaris

### Architectures
amd64, arm64, 386, arm, loong64, ppc64, ppc64le, riscv64, mips, mips64, s390x, wasm

### Compression Formats
`.zip`, `.tar.gz`, `.tgz`, `.tar.bz2`, `.tar.xz`, `.txz`, `.gz`, `.bz2`, `.bin`

## Tips

### GitHub API Rate Limits

GitHub API limits:
- Unauthenticated: 60 requests/hour
- Authenticated: 5000 requests/hour

Get a token at https://github.com/settings/tokens and set:

```bash
export GITHUB_TOKEN=ghp_xxxxxxxxxxxx
```

### Apple Silicon Compatibility

The script automatically detects ARM64 versions. If unavailable, falls back to amd64 (runs via Rosetta 2).

### Security Warning

⚠️ Piping scripts from the internet to bash carries security risks. Always inspect scripts first or deploy your own instance.

```bash
# Inspect before executing
curl https://your-domain.workers.dev/user/repo
```

## License

MIT License - Copyright © 2020 Jaime Pillora &lt;dev@jpillora.com&gt;
