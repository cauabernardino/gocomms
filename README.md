# GoComms 

ON DEVELOPMENT.

An application that resembles a textual social network, made in Go. 

The idea is that the application has few dependencies, to create as much as possible of web application functionalities from scratch, to learn more about the Go language and web development.


*Based on the final project of the course [Aprenda Golang](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/).*


## Dependencies

The list of external packages used are:
- [gorilla/mux](https://github.com/gorilla/mux) for handling routing
- [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql) for connecting to MySQL database
- [checkmail](https://github.com/badoux/checkmail) for validating e-mail format
- [jwt-go](https://github.com/dgrijalva/jwt-go) for creating and validating JSON Web Tokens
- [GoDotEnv](https://github.com/joho/godotenv) for loading environment variables
- [Go Cryptography](https://golang.org/x/crypto) for encrypting passwords


## How to run ⚙️

- Database

The API is configured to use MySQL database listening in the default port 3306. You can installed it directly accessing [here](https://dev.mysql.com/downloads/mysql/), or use it from a Docker image if you prefer, looking [here](https://hub.docker.com/_/mysql).

- API

Clone the repo
```bash
git clone https://github.com/cauabernardino/gocomms.git
```

Enter the API folder
```bash
cd gocomms/api
```

Download dependencies
```bash
go mod download
```

Start the API

Clone the repo
```bash
# Without compiling
go run ./main.go

# Compiling
go build -o api
./api
```