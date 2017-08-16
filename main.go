package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
	CommandRegex *regexp.Regexp
	OverlayRegex *regexp.Regexp
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

	var re *regexp.Regexp
	var err error

	if re, err = regexp.Compile(" *!([a-z]+) *"); err == nil {
		CommandRegex = re
	} else {
		fmt.Println("error compiling command regexp,", err)
		return
	}

	if re, err = regexp.Compile(".*overlay.*"); err == nil {
		OverlayRegex = re
	} else {
		fmt.Println("error compiling overlay regexp,", err)
		return
	}
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// FIXME: Delete me, testing!
	dg.AddHandler(somethingCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if OverlayRegex.MatchString(m.Content) {
		s.ChannelMessageSend(m.ChannelID, "Shut up your fucking mouth and learn to read, damn asshole!")
	}

	if cmd := CommandRegex.FindStringSubmatch(m.Content); len(cmd) > 1 {

		res := ""

		switch cmd[1] {
			case "yo":
				res = "ya"
			case "help":
				res = "Commands:\n!help\n!money\n!start\n!payouts\n!support\n!mods\n!waiting\n!intro"
			case "money":
				res = "How much money to make -> https://loots.com/en/account/how-loots-works"
			case "start":
				res = "How to get started -> https://loots.com/en/account/how-loots-works"
			case "payouts":
				res = "Next payout: [DATE] - will run for 3 days, finished on [DATE+3]"
			case "support":
				res = "Support form: -> https://loots.com/en/account/support"
			case "mods":
				res = "mod names, languages, timezone\nand some more to 'auto-post' every x hours (in addition to use it as a 'manual' command)"
			case "waiting":
				res = "Current waiting list status: ['confirmed' entries], current waiting time: 14 days, join now [https://loots.com/en/auth/waiting]"
			case "intro":
				res = "How loots works -> https://loots.com/en/account/how-loots-works"
			default:
				res = "Unrecognized command, what the heck are you typing?"
				fmt.Println("Unrecognized Command: ", cmd[1])
		}

		s.ChannelMessageSend(m.ChannelID, res)
	}	
}

func somethingCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Paaaaaaaaaang!")
	}
}