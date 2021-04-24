# DiscordBot

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>

## About The Project

There are many great Discord bots out there, but I wanted to develop my own. I wanted the freedom of being able to easily add my own commands and tailor it to my own needs. After thinking a bit about which language to use, I decided on developing it in Go. I decided on Go after my newfound love for the programming language after the cloud technologies course I had during my bachelor in programming. Of course this Discord bot may not cover all of your own needs, but you're more than welcome to fork this repo and change it how you like.

### Built With

This discord bot is built with:
* [DiscordGo](https://github.com/bwmarrin/discordgo)

## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

In order to run this project you will need the following:
* Go

### Installation

1. Go to the [Discord developer portal](https://discord.com/developers)
2. Create a new application
3. Add a bot user to the application
4. Get the token for the bot
5. Clone the repo
   ```sh
   git clone https://github.com/Dahl99/DiscordBot.git
   ```
3. Install [DiscordGo](https://github.com/bwmarrin/discordgo)
   ```sh
   go get github.com/bwmarrin/discordgo
   ```
4. Create the file `token.go` in the cmd/ folder and set the token
   ```Go
   const authToken string = "ENTER YOUR TOKEN";
   ```
   
## Usage

To run the discord bot from root directory:
```sh
go run cmd/*.go
```

To run the discord bot from cmd/ directory:
```sh
go run *.go
```

## Contact

Project Link: [https://github.com/Dahl99/DiscordBot](https://github.com/Dahl99/DiscordBot)

## Acknowledgements

* [DiscordGo](https://github.com/bwmarrin/discordgo)
