# train-me-maybe

This project watches for free seats on Regiojet trains.

## Features

Incrementally, we will add these features:

### Basic features

- [ ] (one-shot) based on datetime, ID from and ID to, find the given route and check if there is a free seat
- [ ] (one-shot) send a notification to an email if there is a free seat
- [ ] (periodic) watch for the seats with a configurable interval

### Automatic purchase

- [ ] allow the user to enter a regojet login (kreditová jízdenka)
- [ ] authenticate against the regiojet API
- [ ] if there is a free seat and the user has enough credit, buy the seat automatically and notify the user

### Usability features

- [ ] notify the user if they have low credit
- [ ] allow the user to put an upper limit on the price of the automatically-purchased ticket
- [ ] notify the user on their Discord channel
- [ ] allow the user to enter/choose the name of the stations, instead of its IDs (can be some simple terminal UI when starting the app)
- [ ] allow the user to configure the application via a environment variables and/or configuration file (not only via command line arguments)

### Multi-tenancy features

- [ ] deploy the application to a server
- [ ] allow the user to configure the application via a web interface
- [ ] allow multiple users to use the deployed application

## Configuration

We support configuration via:

1. `config.toml` file
2. environment variables prefixed with `REGIOJET_` - can be loaded from a `.env` file

The later options have higher priority.

### General settings

- `POLL_INTERVAL` - how often to check for the free seats (in seconds). Recommended value is at least `60`.

### Email notifications

We need to know how to connect to an SMTP server to send emails. And also to whom. The following environment variables are used:

- REGIOJET_SMTP_SERVER
- REGIOJET_SMTP_PORT
- REGIOJET_SMTP_USERNAME
- REGIOJET_SMTP_PASSWORD
- REGIOJET_SMTP_RECIPIENT

For example for gmail:

- REGIOJET_SMTP_SERVER=smtp.gmail.com
- REGIOJET_SMTP_PORT=587
- REGIOJET_SMTP_USERNAME=<your-email>@gmail.com
- REGIOJET_SMTP_PASSWORD=[you need to generate an app password](https://myaccount.google.com/apppasswords)
- REGIOJET_SMTP_RECIPIENT=<your-email>@gmail.com

To test that the email notification works for you, you can run 

```
go run .\cmd\email-test\main.go
```

