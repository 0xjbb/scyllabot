package main

import (
	"flag"
	"fmt"
	"github.com/0xjbb/scyllago"
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
	discord.Close()
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

	command := strings.Split(message.Content, " ")
	fmt.Println(command[0])

	if command[0] != Prefix+"scylla" && command[0] != Prefix+"0dayisgay" {
		return
	}

	// Cheeky little easter egg for 0day and Briskets <3
	if command[0] == Prefix+"0dayisgay" {
		session.ChannelMessageSend(message.ChannelID, "``` I know, so is briskets! ```")
		return
	}

	if len(command) == 1 {
		session.ChannelMessageSend(message.ChannelID, "``` Print usage.```")
		return
	}

	switch command[1] {
	case "username", "password", "domain", "email":
		if len(command) != 3 || len(command) != 4 ){
			session.ChannelMessageSend(message.ChannelID, "``` Print usage.```")
			return
		}

		query := fmt.Sprintf("%s:%s", command[1], command[2])
		result, err := getQueryData(query)

		if err != nil {
			fmt.Println(err) // do this better
			return
		}

		for _, values := range result {
			test := buildMessage(values.Fields.Username, values.Fields.Password, values.Fields.Email, values.Fields.Domain)
			session.ChannelMessageSend(message.ChannelID, test)
		}
		break
	default:
		return
	}

}

// just gets data from ze api
func getQueryData(query string) ([]scyllago.Result, error) {
	size := 2
	start := 0

	r, err := scyllago.Query(query, size, start)

	if err != nil {
		fmt.Println(err) // do this better
		return nil, err
	}

	return r, nil
}

func buildMessage(u string, p string, e string, d string) string {
	return fmt.Sprintf("```Username : %s \nPassword : %s \nEmail : %s \nDomain : %s```", u, p, e, d)
}

func Usage() string{
	return ""
}