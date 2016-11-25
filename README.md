# Epirank

Epirank retrieves all students and their GPAs from EPITECH in a sqlite database.

## Docker

### Build the Docker image from the sources

```
docker build -t epirank .
```

### Run from the sources

```
docker run --rm --name epiranking
	-v "~/epirank_data":/tmp
	-e EPIRANK_LOGIN=INTRA_EMAIL
	-e EPIRANK_PASSWORD=INTRA_PASSWORD
	epirank -t
```
