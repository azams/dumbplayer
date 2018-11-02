# DumbPlayer

![alt text](https://github.com/azams/dumbplayer/raw/master/dumbplayer.png "DumbPlayer logo")

DumbPlayer is a mp3 player in MacOS terminal using afplay. DumbPlayer will play all mp3 files in current directory.
# Installation Guide
  - Install brew -> `/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`
  - Install afplay -> `brew install afplay`
  - Install golang -> `https://golang.org/dl/` or `brew install golang`
  - Clone repo -> `git clone https://github.com/azams/dumbplayer;cd dumbplayer`
  - Run without building -> `go run dumbplayer.go`
  - Build binary -> `go build dumbplayer.go`

# Usage
Usage: ./dumbplayer [options]
  - -s	Shuffle mode
  - -l	Loop mode
  - -r	Playback speed (0.4 - 3.0)

# License
MIT

**Free Software, Hell Yeah!**
