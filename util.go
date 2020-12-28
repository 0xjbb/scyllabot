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
	// @TODO fix this crap
	return "```Usage:\n\t$scylla <option> <search query> *<size> *<start>\nOptions:\n\tusername\n\tpassword\n\tdomain\n\temail\nExample:\n\t$scylla username fred\nNote:\n\t size and start are optional, defaults are 5,0\nMax size is 20. use the start position if you need more.```"
}