# 📁 MC-Sync-Upload

### Description 

This is part of the pack MC-Sync. MC-Sync are two programms that aims to make easier to **play minecraft** (java edition) with your friends **without a real server**.

[See MC-Sync-Download](https://github.com/PiterWeb/MC-Sync-Download)

MC-Sync-Upload has the function of send the world (server) to another person or upload it into the firebase cloud (**free** or paid tier) to make backups.

## Purpouse 

📖 Learn 

 - Create better CLI APPs
 - Interact with NBT files (Minecraft configuration files)
 - Use Firebase Storage
 - Read toml files

✨Create a simple programm to play Minecraft Survival with my friends 👤

## Technologies used 📘

 - Go (Golang)

#### External Packages  📦:

 1. [firebase](firebase.google.com/go/v4) (firebase cloud ☁)
 2. [pterm](github.com/pterm/pterm) (beautiful UI)
 3. [toml](github.com/BurntSushi/toml) (config files ⚙)
 4.  [godotenv](github.com/joho/godotenv) (secrets 🤫)
 5. [fasthttp](github.com/valyala/fasthttp) (serve files through http requests)
 6. [google api](google.golang.org/api) (google cloud options)

## Docs

### Prerequisites 📌

 - [Go 1.18](https://go.dev/) 
 - Ngrok

### Set up 💻

    git clone https://github.com/PiterWeb/MC-Sync-Upload.git
    cd ./MC-Sync-Upload
    go mod tidy

Configure the project with the .env , serviceAccountKey.json and  accounts.toml files.

### Build 🔨

    go build .

An executable for your OS will be created

### Play 🪀

Share the executables and also ngrok with your friends

When you open the world in LAN go to ngrok and type:

    ngrok tcp <port>

And send the ngrok domain name to your friends as an IP for the server
