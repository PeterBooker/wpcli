package main

import (
	"os"

	"github.com/PeterBooker/wpcli/internal/commands"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app   = kingpin.New("wpcli", "CLI Tool for WordPress.")
	debug = app.Flag("debug", "Enable debug mode.").Bool()
)

func main() {
	commands.RegisterAll(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
