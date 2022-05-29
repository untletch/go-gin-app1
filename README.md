## About The Project
Example App

## Clone the Repo

    git clone https://github.com/untletch/go-gin-app1.git

## Run the app

    cd go-gin-app1
    go build -o app
    ./app

## Run the tests

    go test -v

# REST API

The REST API to the app is described below.

## Get list of Articles

### Request

`GET /article/`

    curl "localhost:8080" \
    --header "Accept: application/json" \
    --include

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Sun, 29 May 2022 12:58:05 GMT
    Content-Length: 113

    [{"id":1,"title":"Article 1","content":"Article 1 body"},{"id":2,"title":"Article 2","content":"Article 2 body"}]

## Get a specific Article

### Request

`GET /article/id`

    curl "localhost:8080/article/view/1" \
    --header "Accept: application/json" \
    --include

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Sun, 29 May 2022 13:04:13 GMT
    Content-Length: 55

    {"id":1,"title":"Article 1","content":"Article 1 body"}

## Get a non-existent Article

### Request

`GET /article/view/id`

    curl "localhost:8080/article/view/1000" \
    --header "Accept: application/json" \
    --include

### Response

    HTTP/1.1 404 Not Found
    Content-Type: application/json; charset=utf-8
    Date: Sun, 29 May 2022 12:56:05 GMT
    Content-Length: 54

    {"Error Code":404,"Error Message":"Article not found"}
