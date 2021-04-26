# Insults bot

A stupidly silly and unnecessary bot with no practical use besides insulting people.
Created so that some idiot could learn some stuff or something.

# Development

The program assumes that on of two environment variables are set; TOKEN or TOKEN_FILE. If TOKEN is set then that will be used as the bot token. If TOKEN_FILE is set then the file it points to will be read and used as the bot token. Keep in mind how insecure it is to keep this env var set when running untrusted programs from the same shell, and only leave it set for the minimum time possible.

## Building

This project can be built in two ways.

### Manually

To build:
```bash
go build -o bin/bot
```

To run:
```bash
TOKEN=<...> ./bin/bot
```

Or build and run in one with:
```bash
TOKEN=<...> go run
```

### With Docker

To build an image:
```bash
docker build -t insults-bot .
```

To run a container based on built image
```bash
docker run --rm -it -e TOKEN=<...> insults-bot
```

# Deployment

There are a few considerations to take into account when considering how to deploy the bot. Mainly how it's built, how to run it and keep it running, and how to configure access to the secret token.

## Heroku

There are two main methods for deploying the app to Heroku. The default Heroku way, using the heroku-20 stack. Or by deploying the app as a container using the container stack. The default way is configured by the `Procfile`, and the container way is configured by `heroko.yml` and the `Dockerfile`.

Deploy the default way:
```bash
# Login to heroku if you aren't already
heroku login
# Create the heroku app
heroku create
# Set the secret bot token
heroku config:set TOKEN=<...>
# Deploy to heroku
git push heroku main
```

To configure heroku to deploy the app as a container, run the following commands just before first deploying to heroku:
```bash
# Set the stack to container as opposed to heroku-20
heroku stack:set container
# Redeploy
git push heroku main
```

## GCP via Container optimized OS

> TODO

## Kubernetes?

> TODO

## Manually on OpenStack or other IaaS solution

> TODO
