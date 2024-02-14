# Random HTTP Cats

## Description

This is a small Gin server named `random-http-cats` that retrieves a random HTTP status from `https://http.cat/`.

## Installation

To get a development environment running, clone the repository and navigate into the directory:

```bash
git clone <repository-link>
cd <repository-directory>
```

Then, install the dependencies:

```bash
go mod download
```

## Usage

To start the server, run:

```bash
go run main.go
```

Or you can build the project using the Makefile:

```bash
make build
```

And then run the generate binary:

```bash
./random-http-cats
```

### Docker

The application is also containerized using Docker. You can build and run the Docker image with:

```bash
docker build -t random-http-cats .
docker run -p 8080:8080 random-http-cats
```
