# Project lotery_viking

This project is an API for a lottery system for Viking. This will make the application more scalable and easier to maintain.

## Getting Started

Based on the Makefile, you can run the following commands to build and run the application.

You need to create a .env file from the .env.example file and fill in the necessary information.

Using the Framework Gin, you need to set the environment variable GIN_MODE=release to run the application in production mode.

```bash
# .env
GIN_MODE=release
```
In development mode, the application will run with the default configuration.

```bash
# .env
GIN_MODE=debug
```

> **Note:** The application uses a Mysql database, so you need to have a Mysql instance running.


### Prerequisites

- Go 1.2X
- Mysql
- Docker


## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```
