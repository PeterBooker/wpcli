package main

import (
	"os"

	"github.com/PeterBooker/wpcli/pkg/commands"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app   = kingpin.New("wpcli", "CLI Tool for WordPress.")
	debug = app.Flag("debug", "Enable debug mode.").Bool()
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	commands.RegisterAll(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
