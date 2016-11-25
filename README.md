# Epirank

## Docker running 

```
docker run --rm --name epiranking 
	-v "/Users/junger_m/Desktop/hehe":/tmp 
	-e EPIRANK_LOGIN=INTRA_EMAIL
	-e EPIRANK_PASSWORD=INTRA_PASSWORD 
	epirank -t
```