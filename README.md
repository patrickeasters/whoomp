# whoomp

Just a really dumb app that plays "Whoomp, there it is!" a lot of times and stores state in Redis

## Build

```
CGO_ENABLED=0 GOOS=linux go build
docker build -t patrickeasters/whoomp .
```

## run

```
docker run -d -p 3000:3000 patrickeasters/whoomp
```
