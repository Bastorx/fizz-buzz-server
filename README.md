# Fizz Buzz Test

## Get started

- Docker way

```
docker-compose build
docker-compose up
```

_Note by default the port 4000 is used_

- Without Docker

```
go get -d -v ./...
go install -v ./...
go build -v ./...
./fizz-buzz-server
```

_Note : Be sure that fizz-buzz-server file has execution right, if necessary : `chmod +x ./fizz-buzz-server`_

## Run test

```
go test .
```