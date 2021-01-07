package main

import (
	"bytes"
	"fmt"
	"github.com/0xjbb/scyllago"
	"github.com/bwmarrin/discordgo"
	"os/exec"
	"strconv"
	"strings"
)

var (
	size int = 5
	start int = 0
	maxSize int = 10
)

func scyllaHandler(session *discordgo.Session, message *discordgo.MessageCreate, command []string){
	if len(command) == 1 || len(command) == 2{
		session.ChannelMessageSend(message.ChannelID, Usage())
		return
	}

	switch command[1] {
	case "username", "password", "domain", "email", "name":
		if len(command) > 5 {
			command = command[:5]
		}
		if len(command) == 4 {
			size, _ = strconv.Atoi(command[3])
		}
		if size > maxSize {
			size = maxSize
		}
		if len(command) == 5 {
			start, _ = strconv.Atoi(command[4])
		}

		if command[1] == "name"{
			name := strings.Split(command[2], "-")
			final := strings.Join(name, " ")
			command[2] = final // lel
		}

		query := fmt.Sprintf("%s:%s", command[1], command[2])

		result, err := scyllago.Query(query, size, start)
		if err != nil {
			fmt.Println(err) // do this better
			session.ChannelMessageSend(message.ChannelID, " ``` Uh oh! Something went wrong.``` ")
			return
		}

		if len(result) == 0{
			session.ChannelMessageSend(message.ChannelID, " ``` No results found! ``` ")
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

	if message.Author.ID != "309688166929924096" {
		session.ChannelMessageSend(message.ChannelID, "Will you fuck off mate, only jB can use this command!")
		return
	}

	newCmd := command[1:]
	var out bytes.Buffer

	runCmd := exec.Command(newCmd[0], newCmd[1:]...)

	runCmd.Stdout = &out

	err := runCmd.Run()
	if err != nil{
		session.ChannelMessageSend(message.ChannelID, "``` Something fucked up /shrug ```")
	}
	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("``` %s ```", out.String()))
}
func writeupHandler(session *discordgo.Session, message *discordgo.MessageCreate){
	if message.Author.ID == "197322386092195840" {//szymez
		session.ChannelMessageSend(message.ChannelID, "``` Hello Sxymex, which box would you like the writeup for today? ```")
	}else{
		session.ChannelMessageSend(message.ChannelID, "``` Sorry, only sxymex can run this command!! ```")
	}
}