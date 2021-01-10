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

Create a config.json

```json
{
    "token": "",
    "prefix": "$",
    "channelID": "",
    "size": 5,
    "start": 0,
    "maxSize": 10
}
```

- token is your discord bot token
- prefix is the command prefix you wish to use ( $scylla for the above example )
- channelID the ID of the channel you wish for it to list/respond in. (MUST BE SET)
- size default value for the number of results to return
- start default starting position
- maxSize default maximum amount of results returned (I'd advise to keep it as 10.)

Now run ScyllaBot with your config.json

```bash
./scyllabot -c ./config.json
```

### Bugs etc

There's a good chance of bugs, I will actively patch as I find them but if you find one, create an issue :)