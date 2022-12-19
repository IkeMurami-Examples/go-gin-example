# Gin Example

[Gin](https://github.com/gin-gonic/gin) is a HTTP web framework written in Go (Golang).

# Getting Started

## Init go module

Init go module:

```
go mod init github.com/IkeMurami-Examples/go-gin-example
```

Install basic dependencies:

```
go get -u github.com/spf13/cobra@latest
go get -u github.com/spf13/viper@latest
go get -u github.com/gin-gonic/gin
```

Install Cobra CLI (if it isn't installed):

```
go install github.com/spf13/cobra-cli@latest
```

Initialize Cobra Interface and add a Viper support:

```
cobra-cli init --author "Ike Murami murami.ike@gmail.com" --license GPLv3 --viper
cobra-cli add start --author "Ike Murami murami.ike@gmail.com" 
```

So, we can start our program:

```
go run main.go start
```

## Deploy setting

```
mkdir build && cd build
mkdir config
touch config/gin-example.yaml
touch Dockerfile docker-compose.yml
```