# WeltSparen Concierge
🛎  Your Weltsparen Assistant

This is a very experimental automated notification system for investments platforms.

It logins into your WeltSparen account, scrapes the total net value of your investments and sends you an email.

The goal is to have such a task running on a cronjob, so you don't have to.

## Roadmap
- [x] Scrape WeltSparen
  - [x] Login
  - [x] Extract Total Assets Value
- [x] Custom Credentials (injection)
- [ ] Makefile
- [ ] Docker Image (to run everywhere)
- [ ] Setup a custom crontask

## Running Locally

1. First, set your `.env` file, based on the `.env.example`.
1. `make install` to install dependencies
1. `make run` to run the code
1. If eveything goes well, you'll receive an email, e.g.:

```
WeltSparen Update

Investments Total Value: € 192.013,25
```

Note: It won't output anything unless there is an error.
