# Mini URL Builder API

## Overview

This Golang API will generate Mini URLs based on a given range of IDs
and will store them in a mongoDB database

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
