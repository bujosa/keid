# keid
This is a simple api in go using redis as primary datastore
and chi as http router.

## Goals

- [x] Create a simple api in go
- [ ] Use redis as primary datastore
- [x] Use chi as http router

## How to run

### Run redis server

```bash
docker run --name redis -p 6379:6379 -d redis
```

### Run the api

```go
go run main.go
```

## References

- [Redis](https://redis.io/)
- [Go](https://golang.org/)
- [Chi](https://github.com/go-chi/chi)
