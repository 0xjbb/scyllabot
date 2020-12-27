package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func scyllaHandler(session *discordgo.Session, message *discordgo.MessageCreate, command []string){
	if len(command) == 1{
		session.ChannelMessageSend(message.ChannelID, Usage())
		return
	}

	switch command[1] {
	case "username", "password", "domain", "email":
		if len(command) != 3 && len(command) != 4 {
			session.ChannelMessageSend(message.ChannelID, Usage())
			return
		}
		if len(command) == 4 {
			
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

func execHandler(session *discordgo.Session, message *discordgo.MessageCreate, command []string){

}