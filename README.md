# Epirank 🏆

Epirank is a program that retrieves all students and their GPAs from EPITECH
and makes them available by web & JSON.

The first time you will start the server, it will update the datase and then will update it every day around 3am.

## Project

The webservice runs using :
- [Gin](https://github.com/gin-gonic/gin) framework and listen on the port `8080`.
- [Go-sqlite3](https://github.com/mattn/go-sqlite3) as a SQLite driver.
- [Logrus](https://github.com/Sirupsen/logrus) for logging.

## Route

There is only one route :

- `GET /` : It can take some parameters in the query.

|Parameter|Description|Optional|
|---|---|---|
|city|The city we want (STG, ...). If parameter is not present we ask for every city.|✅|
|promotion|The promotion we want (tek1, tek2, ...)|✅|
|format|We can ask JSON. If not present, returns the HTML ranking page|✅|

## Run from sources

If you are using Fish shell
```
set -gx EPIRANK_LOGIN ""
set -gx EPIRANK_PASSWORD ""
```

Then :
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

### Run from the sources

- The database which is storing all students data is in `/tmp`
- EPIRANK_LOGIN is your EPITECH's email
- EPIRANK_PASSWORD is your EPITECH's intranet password
- GIN_MODE [debug/release]

```
docker run --name epirank
	-v "~/epirank_data":/tmp
	-e EPIRANK_LOGIN=INTRA_EMAIL
	-e EPIRANK_PASSWORD=INTRA_PASSWORD
	-e GIN_MODE=release
	epirank -t
```
