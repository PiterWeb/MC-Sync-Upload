# π MC-Sync-Upload

### Description 

This is part of the pack MC-Sync. MC-Sync are two programms that aims to make easier to **play minecraft** (java edition) with your friends **without a real server**.

[See MC-Sync-Download](https://github.com/PiterWeb/MC-Sync-Download)

MC-Sync-Upload has the function of send the world (server) to another person or upload it into the firebase cloud (**free** or paid tier) to make backups.

## Purpouse 

π Learn 

 - Create better CLI APPs
 - Interact with NBT files (Minecraft configuration files)
 - Use Firebase Storage
 - Read toml files

β¨Create a simple programm to play Minecraft Survival with my friends π€

## Technologies used π

 - Go (Golang)

#### External Packages  π¦:

 1. [firebase](firebase.google.com/go/v4) (firebase cloud β)
 2. [pterm](github.com/pterm/pterm) (beautiful UI)
 3. [toml](github.com/BurntSushi/toml) (config files β)
 4.  [godotenv](github.com/joho/godotenv) (secrets π€«)
 5. [fasthttp](github.com/valyala/fasthttp) (serve files through http requests)
 6. [google api](google.golang.org/api) (google cloud options)

## Docs

### Prerequisites π

 - [Go 1.18](https://go.dev/) 
 - Ngrok (I mention ngrok but you can use other tools that supports tcp tunnels)
 - Firebase Project (Firebase Storage)

### Set up π»

    git clone https://github.com/PiterWeb/MC-Sync-Upload.git
    cd ./MC-Sync-Upload
    go mod tidy

Configure the project with the .env , serviceAccountKey.json and  accounts.toml files.

### Build π¨

    go build .

An executable for your OS will be created

### Play πͺ

Share the executables (mc-sync-upload / download) and also ngrok with your friends

When you open the world in LAN go to ngrok and type:

    ngrok tcp <port> --region <region to host>

And send the ngrok domain name to your friends as an IP for the server
(Example: 74329sadoijjiosad.ngrok.io )
