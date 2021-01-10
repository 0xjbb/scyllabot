package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/0xjbb/scyllago"
	"github.com/bwmarrin/discordgo"
)

type ScyllaCfg struct{
	session *discordgo.Session
	message *discordgo.MessageCreate
	size int
	start int
	maxSize int
}

// $scylla -username Joe Blogs -password test -size 5 -start 0
func ScyllaNew(session *discordgo.Session, message *discordgo.MessageCreate, size int, start int, maxSize int)  ScyllaCfg{
	return ScyllaCfg{
		session: session,
		message: message,
		size: size,
		start: start,
		maxSize: maxSize,
	}
}

// maybe break this function up into a few smaller funcs
func (sc *ScyllaCfg) Handle(command []string){
	sFlag := flag.NewFlagSet("Scylla", flag.ContinueOnError)
	username := sFlag.String("user", "", "Username you wish to search")
	password := sFlag.String("passw", "", "Password you wish to search")
	email := sFlag.String("email", "", "Email you wish to search")
	domain := sFlag.String("url", "", "Domain you wish to search")
	ip := sFlag.String("ip", "", "IP address you wish to search")
	passhash := sFlag.String("passh", "", "Password hash you wish to search")
	size := sFlag.Int("size", sc.size, "Number of results to return (max 10)")
	start := sFlag.Int("start", sc.start, "Result starting position.")

	sFlag.Usage = sc.usage(sFlag)

	err := sFlag.Parse(command)

	if err != nil {
		fmt.Println(err)
	}

	qVars := make(map[string]string, 6)

	qVars["username"] = *username
	qVars["password"] = *password
	qVars["email"] = *email
	qVars["domain"] = *domain
	qVars["ip"] = *ip
	qVars["passhash"] = *passhash

	query := ""

	for key,  val := range qVars{
		if val == ""{ // Skip any that don't have a string
			continue
		}

		if query == ""{
			query = fmt.Sprintf("%s:%s", key, val)
			continue
		}

		query = fmt.Sprintf(" & %s:%s", key, val)
	}

	if query == ""{
		// print usuage I guess.
		return
	}

	fmt.Println("Query sent: ", query)

	result, err := scyllago.Query(query, *size, *start)

	if err != nil {
		fmt.Println(err) // do this better
		// send message bask to user.
		return
	}

	for _,  val := range result{
		fmt.Println(val)// send to channel.
	}
}

// send usage to channel instead of stdout/err
func (sc *ScyllaCfg) usage(fs *flag.FlagSet) func(){
	buffer := new(bytes.Buffer)
	fs.SetOutput(buffer)

	return func() {
		fs.PrintDefaults()
		//sc.session.ChannelMessageSend(sc.message.ChannelID, fmt.Sprintf("``` %s ```", buffer.String()) )
		sc.SendEmbed("Usage: ", fmt.Sprintf("``` %s ```", buffer.String()))
	}
}

func (sc *ScyllaCfg) SendEmbed(name string, value string){
	messageEmbed := discordgo.MessageEmbed{
		Title: "__ScyllaBot__",
		Fields: func() []*discordgo.MessageEmbedField {
			var embedFields []*discordgo.MessageEmbedField

			currentEmbed := discordgo.MessageEmbedField{
				Name:  name,
				Value: value,
			}

			embedFields = append(embedFields, &currentEmbed)

			return embedFields
		}(),
	}
	sc.session.ChannelMessageSendEmbed(sc.message.ChannelID, &messageEmbed)
}