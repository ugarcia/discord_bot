package commands

import (
	"fmt"
	"regexp"
	"github.com/bwmarrin/discordgo"
)

// Command Base struct for all commands
type Command struct{
	regex *regexp.Regexp
}

// CommandInterface The interface Commands should implement
type CommandInterface interface{
	Listen(dg *discordgo.Session)
}

func (c *Command) getRegexString() string{
	return ""
}

func (c *Command) matches(s string) (b bool, m []string){
	b = false
	m = make([]string, 0)
	if cmd := c.regex.FindStringSubmatch(s); len(cmd) > 1 {
		b = true
		m = cmd[2:]
	}
	return
}

func (c *Command) doExecute(matches []string, s *discordgo.Session, m *discordgo.MessageCreate){
}

func (c *Command) execute(s *discordgo.Session, m *discordgo.MessageCreate){
	if ok, matches := c.matches(m.Content); ok {
		c.doExecute(matches, s, m)
	} else {
		return
	}
}

// Listen Listener of the command
func (c *Command) Listen(dg *discordgo.Session){
	rst := c.getRegexString()
	if re, err := regexp.Compile(rst); err == nil {
		c.regex = re
	} else {
		fmt.Println("error compiling command regexp,", err, rst)
		return
	}
	dg.AddHandler(c.execute)
}
