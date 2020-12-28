package main

import (
	"fmt"
)

// just gets data from ze api

func buildMessage(u string, p string, e string, d string) string {
	// change to embed maybe.
	return fmt.Sprintf("```Username : %s \nPassword : %s \nEmail : %s \nDomain : %s```", u, p, e, d)
}

func Usage() string{
	// change to embed or smth.
	return "```\tUsage:\n $scylla <option> <search query>\n\tOptions:\n username\npassword\ndomain\nemail\n\tExample:\n$scylla username fred\n```"
}