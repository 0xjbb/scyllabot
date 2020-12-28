package main

import (
	"fmt"
	"github.com/0xjbb/scyllago"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func scyllaHandler(session *discordgo.Session, message *discordgo.MessageCreate, command []string){
	if len(command) == 1 || len(command) == 2{
		session.ChannelMessageSend(message.ChannelID, Usage())
		return
	}

	if len(command) > 5 {
		command = command[:5]
	}

	switch command[1] {
	case "username", "password", "domain", "email":
		start := 0
		size := 5

		if len(command) == 4 {
			size, _ = strconv.Atoi(command[3])
		}

		if len(command) == 5 {
			start, _ = strconv.Atoi(command[3])
		}
		query := fmt.Sprintf("%s:%s", command[1], command[2])
		result, err := scyllago.Query(query, size, start)

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

func execHandler(session *discordgo.Session, message *discordgo.MessageCreate, command []string){

}