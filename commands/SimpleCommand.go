package commands

import (
	"strings"
	"github.com/bwmarrin/discordgo"
)

// SimpleCommand A simple command
type SimpleCommand Command

func (c *SimpleCommand) getRegexString() string{
	return " *!([a-z]+) *"
}

func (c *SimpleCommand) doExecute(matches []string, s *discordgo.Session, m *discordgo.MessageCreate){
	s.ChannelMessageSend(m.ChannelID, "Yayyyyy!!!! " + strings.Join(matches, "::"))
}

// Listen Listener of the command
func (c *SimpleCommand) Listen(dg *discordgo.Session){
	c.Command.Listen(dg)
}
