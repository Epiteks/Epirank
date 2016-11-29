# Epirank 🏆

Epirank is a program that retrieves all students and their GPAs from EPITECH
and makes them available by web & JSON.

The first time you will start the server, it will update the datase and then will update it every day around 3am.

## Project

The webservice runs using :
- [Gin](https://github.com/gin-gonic/gin) framework and listen on the port `8080`.
- [Go-sqlite3](https://github.com/mattn/go-sqlite3) as a SQLite driver.
- [Logrus](https://github.com/Sirupsen/logrus) for logging.

## Endpoint

There is only one route :

- `GET /` : It can take some parameters in the query.

|Parameter|Description|Optional|
|---|---|---|
|city|The city we want (STG, ...). If parameter is not present we ask for every city.|✅|
|promotion|The promotion we want (tek1, tek2, ...)|✅|
|format|We can ask JSON. If not present, returns the HTML ranking page|✅|

## Authentication file

```json
{
	"login": "email",
	"password": "password"
}
```

If you want to delete the configuration file after the launch, set the environment variable :

If you are using Fish shell:
```
set -gx EPIRANK_DELETE_CREDENTIALS true
```


## Run from sources

```
go build
./Epirank
```

## Run with Docker 🐳

### Build the Docker image from the sources

```
docker build -t epirank .
```

### Pull image from Docker Hub

TODO

### Docker run

- The database which is storing all students data is in `/tmp`. The authentication file too.
- EPIRANK_DELETE_CREDENTIALS : if set at true, deletes the authentication file
- GIN_MODE [debug/release]

```
docker run -d --name epirank
	-v "~/epirank_data":/tmp
	-e EPIRANK_DELETE_CREDENTIALS=true
	-e GIN_MODE=release
	epirank
```
