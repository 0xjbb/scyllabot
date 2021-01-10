package main

/*

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
}*/
