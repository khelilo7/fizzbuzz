# Fizzbuzz-api

Fizzbuzz-api is an API that get the Fizzbuzz string.

## Start Application

Launching postgres with docker-compose

```bash
docker-compose up -d
```
to check your application is up and runnig you can use
```bash 
docker logs [container]
```
to shutdown the postgres 
```bash
docker-compose down
```

To start the application
```bash
go run cmd/main.go
```

## Test
First you have to run the docker postgres image, everything is already configured inside the docker-compose
```bash
docker-compose up -d
```
Then you just have to wait for the database to be created you can check that with
```bash 
docker logs [container]
```
Once the database is created  
```bash
cd ..
go test -v ./... --coverprofile cover.out
```

to shutdown the postgres container 
```bash
docker-compose down
```

## API

The REST API to the annonce app is described below.

## Get fizzbuzz string

### Request

`POST /fizzbuzz/`

    curl -i -H 'Accept: application/json' -d 'int1=2&int2=3&limit=6&str1=fizz&str2=buzz' http://localhost:8081/fizzbuzz/

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    1,fizz,buzz,fizz,5,fizzbuzz

## Create api stats

### Request

`Get /api_stats/`

    curl -i -H 'Accept: application/json' http://localhost:8081/api_stats

### Response

    HTTP/1.1 201 Created
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 201 Created
    Connection: close
    Content-Type: application/json
    Location: /annonce/1
    Content-Length: 36

    {api_stats...}

## Get request stats

### Request

`GET /statistics`

    curl -i -H 'Accept: application/json' http://localhost:8080/statistics

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 36

    {stats...}

## Tables
### Request
- id  BIGSERIAL PRIMARY KEY NOT NULL
- creation_date TIMESTAMP WITH TIME ZONE  DEFAULT NOW() NOT NULL
- last_update TIMESTAMP WITH TIME ZONE  DEFAULT NOW() NOT NULL
- int1   integer   DEFAULT 0 NOT NULL
- int2   integer   DEFAULT 0 NOT NULL
- limite   integer   DEFAULT 0
- str1   VARCHAR(50)   DEFAULT ''
- str2   VARCHAR(50)   DEFAULT ''

# Author
Khalil JRIDI.