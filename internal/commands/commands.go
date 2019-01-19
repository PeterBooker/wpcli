package commands

import (
	"github.com/PeterBooker/wpcli/internal/commands/core"
	"gopkg.in/alecthomas/kingpin.v2"
)

// RegisterAll ...
func RegisterAll(app *kingpin.Application) {
	core.Register(app)
}
