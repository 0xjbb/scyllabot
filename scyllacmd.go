package main

import (
	"flag"
	"fmt"
	"github.com/0xjbb/scyllago"
	"github.com/bwmarrin/discordgo"
)
var (
	query string = ""
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

	err := sFlag.Parse(command)

	if err != nil {
		fmt.Println(err)
	}

	test := make(map[string]string, 6)

	test["username"] = *username
	test["password"] = *password
	test["email"] = *email
	test["domain"] = *domain
	test["ip"] = *ip
	test["passhash"] = *passhash


	for key,  val := range test{
		if val == ""{ // Skip any that don't have a string
			continue
		}

		if query == ""{
			query = fmt.Sprintf("%s:%s", key, val)
			continue
		}

		query = fmt.Sprintf(" & %s:%s", key, val)
	}

	result, err := scyllago.Query(query, *size, *start)

	if err != nil {
		fmt.Println(err) // do this better
		// send message bask to user.
		return
	}

	for _,  val := range result{
		fmt.Println(val)
	}
}
