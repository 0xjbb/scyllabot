package main

import (
	"encoding/json"
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
	ConfigFileName  string
	config Config
)

type Config struct{
	Channel string `json:"channelID"`
	Token string `json:"token"`
	Prefix string `json:"prefix"`
	Size int `json:"size"`
	Start int `json:"start"`
	MaxSize int `json:"maxSize"`
}

func init() {
	flag.StringVar(&ConfigFileName, "c", "config.json", "Config file")
	flag.Parse()
}
/*
	@TODO fix the error handling messages, change to log and use the log library to write to a file.
*/
func main() {
	// @todo Move this to a func.
	configFile, err := os.Open(ConfigFileName)

	if err != nil{
		fmt.Println(err)
		return
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)

	if err != nil{
		fmt.Println(err)
	}

	discord, err := discordgo.New("Bot " + config.Token)

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
	if !strings.HasPrefix(message.Content, config.Prefix) {
		return
	}

	// @TODO change to multiple channelIDs
	if message.ChannelID != config.Channel{
		return
	}

	command := strings.Split(message.Content, " ")

	switch command[0] {
	case config.Prefix + "scylla":
		scylla := ScyllaNew(session, message, config.Size, config.Start, config.MaxSize)
		scylla.Handle(command[1:])
		break
	default:
		break
	}
}

func ParseConfig(){}