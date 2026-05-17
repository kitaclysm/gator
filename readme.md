# Gator CLI
Gator is a CLI application for fetching and browsing RSS feeds.

## Requirements
To run Gator, you need:
- [Go](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation
Install the CLI with:
```bash
go install github.com/YOUR_GITHUB_USERNAME/YOUR_REPO_NAME@latest
```

## Configuration
Create a file named .gatorconfig.json in your home directory:
```
{
    "db_url": "postgres://postgres:password@localhost:5432/gator?sslmode=disable",
    "current_user_name": ""
}
```
Example locations:
- __Linux/macOS:__ ~/.gatorconfig.json
- __Windows:__ your home directory

## Running Gator
After installing, run commands with:
`gator COMMAND`
### Commands
- `register <name>` register a new user
- `login <name>` log in as an existing user
- `users` list all users
- `addfeed <name> <url>` add a new feed
- `feeds` list all feeds
- `follow <url>` follow a feed
- `following` list followed feeds
- `unfollow <url>` unfollow a feed
- `agg <time_between_reqs>` continuously scrape feeds
- `browse <limit>` browse posts
- `reset` reset the database
