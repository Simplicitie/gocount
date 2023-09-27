# Building

To build the application as a single all inclusive executable use, golang's build command arguments as below
```bash
CGO_ENABLED=0 go build -v -ldflags="-w -s" -o gocount
```
You can also add other arguments for example to change the target operating system and architecture.  IE to build for windows 63 bit you would run
```bash
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags="-w -s" -o gocount.exe
```

# Configuration

The application is designed to use environment variables that define a connection to a Redis instance

* "REDIS_USER" - The username to connect as
* "REDIS_PASS" - The password to authenticate with
* "REDIS_HOST" - (default "localhost") The hostname or IP address where Redis is found
* "REDIS_PORT" - (default "6379") The port Redis is published on

These variables can be supplied via command line when running the application via `go run` or when it is compiled
```bash
# non-compiled
REDIS_HOST=123.123.456.3 REDIS_USER=bob REDIS_PASSWORD=secret_password REDIS_PORT=6379 go run main.go
# compiled
REDIS_HOST=123.123.456.3 REDIS_USER=bob REDIS_PASSWORD=secret_password REDIS_PORT=6379 ./gocount
```

# Usage

The application publishes a REST interface on port `8080` with the following endpoints

* GET Methods
  * `/`  - Displays a greeting and usage information
  * `/count` - Displays the current counter's value
  * `/incr` - Increments the counter's value by 1 and displays the new value
  * `/reset` - Sets the count back to 0.  **NOTE:** This does not fully erase the key in Redis, it just forces the value to be returned to 0.

Each of these endpoints can be reached via a web browser (IE: http://localhost:8080/count) or with the `curl` command:
```
curl localhost:8080/count
```
