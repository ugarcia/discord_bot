package commands

import (
	"github.com/bwmarrin/discordgo"
)

// Global variables
var (
	cmds []CommandInterface = []CommandInterface{
		&SimpleCommand{},
	}
)

// Global Listener
func Listen(dg *discordgo.Session) {
	for i, cmd := range(cmds) {
		cmd.Listen(dg)
	}
}
