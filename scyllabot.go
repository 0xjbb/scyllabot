package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	Token  string
	Prefix string = "$"
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	discord, err := discordgo.New("Bot " + Token)

	if err != nil {
		log.Fatal("error creating discord session, ", err)
	}

	discord.AddHandler(messageHandler)

	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = discord.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = discord.Close()
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// ignore messages from the bot itself
	if message.Author.ID == session.State.User.ID {
		return
	}
	// ignore other bots.
	if message.Author.Bot {
		return
	}
	// only see messages with out prefix.
	if !strings.HasPrefix(message.Content, Prefix) {
		return
	}
	//793522452952514620 == chanid for PE, other is my testing server
	if message.ChannelID != "793522452952514620" && message.ChannelID != "792923886185611285"{
		return
	}

	command := strings.Split(message.Content, " ")

	switch command[0] {
	case Prefix + "scylla":
		scyllaHandler(session, message, command)
		break
	case Prefix + "0day_is_gay":
		session.ChannelMessageSend(message.ChannelID, "``` I know, so is Briskets. :kekw: ```")
		break
	case Prefix + "writeup":
		if message.Author.ID == "197322386092195840" {//szymez id=197322386092195840
			session.ChannelMessageSend(message.ChannelID, "``` Hello Sxymex, which box would you like the writeup for today? ```")
		}else{
			session.ChannelMessageSend(message.ChannelID, "``` Sorry, only sxymex can run this command!! ```")
		}
	case Prefix + "exec":
		if message.Author.ID == "309688166929924096" {
			execHandler(session,message,command)
		}else{
			session.ChannelMessageSend(message.ChannelID, "Will you fuck off mate, only jB can use this command!")
		}
		break
	case Prefix + "":
		break
	}
}