# Mini URL Builder API

## Overview

This Golang API will generate Mini URLs based on a given range of IDs
and will store them in a mongoDB database.

## Local workspace

First: clone the repository:

```bash
git clone https://github.com/DiegoSepuSoto/mini-url-builder-api
cd mini-url-builder-api
```

Then, download the dependencies

```bash
go mod download
```

Now you can run the application using the Makefile

```bash
make run
```

You can even run the whole solution using **docker compose**:

```bash
make run-compose
```

which create the full stack of the solution:
- Redis cache
- MongoDB
- [Distributed Sync](https://github.com/DiegoSepuSoto/distributed-sync-mock)
- [Mini URL Service](https://github.com/DiegoSepuSoto/mini-url-service)
- Mini URL Builder API (this service)

There's also an option for you to insert basic data on both the data layers with:

```bash
make migrate-dbs
```

The available endpoint is the following:

```bash
curl --location --request POST 'localhost:8080/mini-url' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.8x2hIBGylPBtKnAoEP8wJqqXbXaQyOK0z8bjpasZGfo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "original_url": "https://www.apple.com"
}'
```

which will create a mini URL to be used.

Also, you can access:

- Prometheus metrics at: **localhost:8080/metrics**
- Swagger documentation at: **localhost:8080/swagger/index.html**

### Tech Stack

- Golang library - Echo framework for http server
- Golang library - Logrus for application logs
- Golang library - Viper for application environment variables
- Golang library - Testify for unit testing
- Prometheus metrics
- Docker & Docker Compose
