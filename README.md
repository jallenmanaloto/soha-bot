# Project soha-bot

Soha-bot is a discord bot that scrapes manhwa websites and send an alert to subscribed server for any new chapters.

Soha can look for a manhwa title. He can also watch titles for you and alert you for any new chapters he find.

For the full list of commands, you can type `!soha tricks`.

## Getting Started

1. Install [Go](https://go.dev/) through their official website.
2. Clone the repository to obtain a copy on your local machine.
3. Change directory: `cd soha-bot`.
4. Install dependencies `go mod download`.

## MakeFile

build the application
```bash
make build
```

run the application
```bash
make run
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```
