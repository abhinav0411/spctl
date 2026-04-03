# spctl

A terminal-based Spotify client built with Go, Bubbletea, and Lipgloss.

## Features

- Search songs and play them directly from the terminal
- Browse and play your Spotify playlists
- View tracks inside any playlist
- Player controls — pause, resume, skip next/prev
- Keyboard-driven navigation

## Prerequisites

- Go 1.21+
- A Spotify Premium account
- A registered Spotify app (for API credentials)

## Setup

### 1. Create a Spotify App

1. Go to [Spotify Developer Dashboard](https://developer.spotify.com/dashboard)
2. Create a new app
3. Set the redirect URI to `http://localhost:8080/callback`
4. Copy your `Client ID`

### 2. Configure Environment

Create a `.env` file in the project root:
```env
CLIENT_ID=your_spotify_client_id
REDIRECT_URL=http://localhost:8080/callback
```

### 3. Build and Run
```bash
git clone https://github.com/abhinav0411/spctl
cd spctl
go build ./cmd/spctl/
./spctl
```

## Keybindings

| Key | Action |
|-----|--------|
| `/` | Search songs |
| `enter` | Play selected track |
| `tab` | Switch between results and playlists |
| `v` | View playlist tracks |
| `p` | Play playlist |
| `space` | Pause / Resume |
| `n` | Next track |
| `b` | Previous track |
| `q` | Quit |

## Tech Stack

- [Bubbletea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) — Terminal styling
- [Spotify Web API](https://developer.spotify.com/documentation/web-api)

## Notes

- Spotify Premium is required for playback control via the API
- Make sure you have an active Spotify session open on a device before launching spctl
