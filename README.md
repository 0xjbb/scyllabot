# ScyllaBot

A Simple discord bot to retrieve data from Scylla using scyllago and discordgo

## Getting Started

### Installation

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

Installing the required modules

```bash
go get github.com/0xjbb/scyllago
go get github.com/bwmarrin/discordgo
```

Build the project.

```bash
go build
```

Running with your token
```bash
./scyllabot -t <your discord token here>
```

### Bugs etc

There's a good chance of bugs, I will actively patch as I find them but if you find one, create an issue :)