# Insults bot

A stupidly silly and unnecessary bot with no practical use besides insulting people.
Created so that some idiot could learn some stuff or something.

# Development

The program assumes that on of two environment variables are set; TOKEN or TOKEN_FILE. If TOKEN is set then that will be used as the bot token. If TOKEN_FILE is set then the file it points to will be read and used as the bot token.

## Building

This project can be built in two ways.

1. Manually with `go` -> `go build -o bin/bot`
2. With docker -> `docker build -t insults-bot .`

How to run the project depends on how you built it.

1. Manually -> `./bin/bot`
2. With docker -> `docker run --rm insults-bot`
