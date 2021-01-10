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
	sFlag := flag.NewFlagSet("scylla", flag.ContinueOnError)
	username := sFlag.String("user", "", "Username you wish to search")
	password := sFlag.String("password", "", "Password you wish to search")
	email := sFlag.String("email", "", "Email you wish to search")
	domain := sFlag.String("url", "", "Domain you wish to search")
	ip := sFlag.String("ip", "", "IP address you wish to search")
	passhash := sFlag.String("passh", "", "Password hash you wish to search")
	size := sFlag.Int("size", sc.size, "Number of results to return (max 10)")
	start := sFlag.Int("start", sc.start, "Result starting position.")

	sFlag.Usage = sc.usage(sFlag)

	err := sFlag.Parse(command)

	// find out why the flag library doesn't already do this
	if *username == "" && *password == "" && *email == "" && *ip == "" && *domain == "" && *passhash == ""{
		sFlag.Usage()
		return
	}

	if err != nil {
		fmt.Println("Parse error: ", err)
		return
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
		sc.usage(sFlag)
		return
	}

	fmt.Println("Query sent: ", query)

	result, err := scyllago.Query(query, *size, *start)

	if err != nil {
		fmt.Println("ScyllaGo Error: ",err) // do this better
		// send message bask to user.
		return
	}

	if len(result) == 0{
		sc.SendEmbed("Error", "No results found!")
		return
	}

	// @TODO convert to function
	messageEmbed := discordgo.MessageEmbed{
		Title: "__ScyllaBot__",
		Fields: func() []*discordgo.MessageEmbedField {
			var embedFields []*discordgo.MessageEmbedField
			for _, values := range result {
				currentEmbed := discordgo.MessageEmbedField{
					Name:  "------------------------------------",
					Value: fmt.Sprintf("IP: %s\nUsername: %s\nPassword: %s\nPasshash: %s\nEmail: %s\nDomain: %s",
						values.Fields.Ip,
						values.Fields.Username,
						values.Fields.Password,
						values.Fields.Passhash,
						values.Fields.Email,
						values.Fields.Domain,
						),
				}

				embedFields = append(embedFields, &currentEmbed)
			}
			return embedFields
		}(),
	}
	sc.session.ChannelMessageSendEmbed(sc.message.ChannelID, &messageEmbed)


}

// send usage to channel instead of stdout/err
func (sc *ScyllaCfg) usage(fs *flag.FlagSet) func(){
	buffer := new(bytes.Buffer)
	fs.SetOutput(buffer)

	return func() {
		fs.PrintDefaults()
		sc.SendEmbed("Usage: ", fmt.Sprintf("``` %s ```", buffer.String()))
	}
}
// @todo rewrite.
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
