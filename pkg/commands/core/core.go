package core

import (
	"github.com/PeterBooker/wpcli/pkg/commands/core/download"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Command ...
type Command struct {
}

func (d *Command) run(c *kingpin.ParseContext) error {
	return nil
}

// Register ...
func Register(app *kingpin.Application) {
	cmd := &Command{}
	core := app.Command("core", "Core WordPress Commands.").Action(cmd.run)

	download.Register(core)
}
