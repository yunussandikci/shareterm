# [![Shareterm](https://abload.de/img/shareterm_bannerp1jvz.png)](https://shareterm.tech)

*Current version: 0.0.4*

**ShareTerm** is minimalistc, open source online terminal log sharing tool where the server has zero knowledge of pasted data.
It uses **client-side encryption** with AES-256 Bits. It fastens one simple process while you are developing, and **saves your time**. So that, you don't waste your time with trying to copy, paste, beautify. You **just pipe the `shareterm`** and it returns a simple shareble link.

# Shaterm-CLI

## Usage
You just need to pipe shareterm to your command and it returns sharing url. By default, it uses **ShareTerm** servers.\
Example:
```console
foo@bar:~$ echo "Hey this is shareterm" | shareterm
https://shareterm.tech/read/<name>?key=<key>
```

## Installation
Shareterm available via HomeBrew. You can install with tap and install commands.
```console
brew tap yunussandikci/shareterm
brew install shareterm-cli
```

## Configuration
You can configure shareterm-cli to connect your own server via setting `SHARETERM_HOST` environment variable.\
By default it uses `https://shareterm.tech` which is demo instance that doesn't promise availability of your data for long term.\
Example: 
```console
export SHARETERM_HOST=https://myownserver.com
```

# Shaterm-Server
Shareterm server is responsible from storing encrypted paste data and serving decrypted paste datas via it's web interface.\
It runs on `8080` port and  saved data to `data` folder in it.
## Installation
### Docker
You can directly run server with Docker.\
Example docker-compose.yaml
```yaml
version: '3'
services:
  shareterm:
    image: yunussandikci/shareterm:latest
    restart: always
    volumes:
      - shaterm-volume:/app/data
    networks:
      - web
volumes:
  shareterm-volume:
networks:
  web:
```
### Direct (Development Purposes)
On server folder, `go mod tidy && go run main.go`
