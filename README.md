# train-me-maybe

This project watches for free seats on Regiojet trains.

Incrementally, we will add these features:

## Basic features

- [ ] (one-shot) based on datetime, ID from and ID to, find the given route and check if there is a free seat
- [ ] (one-shot) send a notification to an email if there is a free seat
- [ ] (periodic) watch for the seats with a configurable interval

## Automatic purchase

- [ ] allow the user to enter a regojet login (kreditová jízdenka)
- [ ] authenticate against the regiojet API
- [ ] if there is a free seat and the user has enough credit, buy the seat automatically and notify the user

## Usability features

- [ ] notify the user if they have low credit
- [ ] allow the user to put an upper limit on the price of the automatically-purchased ticket
- [ ] notify the user on their Discord channel
- [ ] allow the user to enter/choose the name of the stations, instead of its IDs (can be some simple terminal UI when starting the app)
- [ ] allow the user to configure the application via a environment variables and/or configuration file (not only via command line arguments)

## Multi-tenancy features

- [ ] deploy the application to a server
- [ ] allow the user to configure the application via a web interface
- [ ] allow multiple users to use the deployed application
