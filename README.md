# keid
This is a simple microservice in go using redis as primary datastore
and chi as http router.

## Goals

- [x] Create a simple microservice in go
- [ ] Use redis as primary datastore
- [ ] Use chi as http router

## How to run

### Run redis server

```bash
docker run --name redis -p 6379:6379 -d redis
```

### Run the microservice

```go
go run main.go
```

## References

- [Redis](https://redis.io/)
- [Go](https://golang.org/)
- [Chi](https://github.com/go-chi/chi)
