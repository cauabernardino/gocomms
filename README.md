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

The API is configured to use MySQL database listening in the default port 3306. You can install it directly accessing [here](https://dev.mysql.com/downloads/mysql/), or use it from a Docker image if you prefer, looking [here](https://hub.docker.com/_/mysql).

The query for creating the database can be found in [database](/api/database).

- API

Clone the repo
```bash
git clone https://github.com/cauabernardino/gocomms.git
```

- Enter the API folder
```bash
cd gocomms/api
```

- Download dependencies
```bash
go mod download
```

- Change the name of `.env-sample` to `.env` and complete with your own environment variables

- The `GenerateSecretKey` function in the `tools` module helps generate a proper key for encrypting the JWT. You can run it putting the following code in the `main.go` file in the first run, saving the generated key to your `.env`, and then removing it.
```go
func init() {
	tools.GenerateSecretKey()
}

```

- Start the API

```bash
# Without compiling
go run ./main.go

# Compiling
go build -o api
./api
```

The API standard port is `5000`.

### Web

- Enter the API folder
```bash
cd gocomms/web
```

- Download dependencies
```bash
go mod download
```
- Change the name of `.env-sample` to `.env` and complete with your own environment variables

- The `GenerateKeys` function in the `tools` module helps generate proper hash and block keys for encrypting the cookies. You can run it putting the following code in the `main.go` file in the first run, saving the generated keys to your `.env`, and then removing it.
```go
func init() {
	tools.GenerateKeys()
}

```

- Start the web application
```bash
# Without compiling
go run web

# Compiling
go build -o web
./web
```

The API standard port is `8080`.
