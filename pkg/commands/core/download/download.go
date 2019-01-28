package download

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Command holds the relevant flags.
type Command struct {
	Path        string
	Locale      string
	Version     string
	SkipContent bool
	Force       bool
}

func (d *Command) run(c *kingpin.ParseContext) error {
	fmt.Printf("version=%s\n", d.Version)
	return nil
}

// Register the command.
func Register(app *kingpin.CmdClause) {
	cmd := &Command{}
	download := app.Command("download", "Download WordPress.").Action(cmd.run)
	download.Flag("version", "Version to download.").StringVar(&cmd.Version)
}
