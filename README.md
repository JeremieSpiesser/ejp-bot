# EJP Bot

This simple bot can run in the background and 

## Usage

```
ejp-bot <region> <delay>
```

- region : "nord", "sud", "paca" or "ouest"
- delay : refreshing delay in seconds ( expecting a positive integer )

If delay equals -1, the bot runs in "oneshot mode" (ie : it will not loop and immediately return and run the commands)

## Compilation

You need the standard golang compiler (`sudo apt install go`).
You can then compile the program with :
`go build ./ejp-bot.go`


Optionnal : Static linking

- Windows target : 
`env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./ejp-bot.go`

- For Linux target : 
`env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./ejp-bot.go`


## TODO

- command to be ran when it's EJP or not, today or tomorrow
    - figure out how to get these commands : json config ? arguments ?
