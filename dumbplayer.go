package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	shuffling = false
	loop      = false
	speed     = "1.0"
)

func main() {
	for i, arg := range os.Args {
		if arg == "-s" {
			shuffling = true
		} else if arg == "-l" {
			loop = true
		} else if arg == "-r" {
			rate, err := strconv.ParseFloat(os.Args[i+1], 64)
			if err != nil {
				fmt.Println("Error: Tigers! Rate must be a float.")
				return
			} else if rate < 0.4 || rate > 4.0 {
				fmt.Println("Error: Tigers! Rate must between 0.4 - 3.0")
				return
			}
			speed = os.Args[i+1]
		} else if arg == "-h" {
			help()
			return
		}
	}
	getLists()
}

func help() {
	fmt.Println("\033[01;38;5;214m\n                   `-:-`\n`/-`           `.:++:-:+/:.            `-/  \n`o++/-`     `-/+/-`     `-/+/-`     `-/+oo  \n`o-`-/+/-.:/+:.     `.`    `.:+/:.-/+:. +o  \n`o-  `:o+/-`     .-/+/+/-`     `-/++-`  +o  \n`o/://:.     `.://:.  `.:+/:.`    `.:+/:+o  \n`oo/.       `/:-` `.:/-.` `-/:`      `-+oo  \n`o/://-`        `://-.-//-        `.:/+:+o  \n`o:  .://:.`     ``     `      `.://:.  /o  \n`o:     `-//:.`             `.///-`     /o  \n`o:       .:+o+/-`       `-/+o+:.       /o  \n`o:       .------.`     `-------.       /o  \n`o:                                     /o  \n`o:              `.......`              /o  \n`o/-:-.  .-:-`   `.:+o+/-`   .-:-` `.::.+o  \n`+oo:..:+/-.        -o.       `.://:../oo+  \n `.:/+o/.           :o-`          .+o+/:.`  \n     `-/+/-.    `-/+/:/+/-`    .://:-`      \n         .:/+::/+:.`   `.:/+:/+/-.          \n	    `-/++-.     .:++:.`             \n                .:/+:-/+/-`                 \n                   `-:.`                    \033[0m")
	fmt.Println("\nUsage: " + os.Args[0] + " [options]\n\t-s\tShuffle mode\n\t-l\tLoop mode\n\t-r\tPlayback speed (0.4 - 3.0)")
}

func getLists() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic("Error: " + err.Error())
	}
looping:
	if shuffling {
		shuffle(files)
	}
	for _, f := range files {
		ext := strings.Split(f.Name(), ".")
		if strings.ToLower(ext[len(ext)-1]) == "mp3" {
			cmd := exec.Command("afplay", f.Name(), "-q", "1", "-r", speed)
			fmt.Println("Playing: " + f.Name())
			cmd.Run()
		}
	}
	if loop {
		goto looping
	}
}

func shuffle(vals []os.FileInfo) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}
