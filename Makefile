DSN=postgres://devuser:password@localhost/go_movies?sslmode=disable
BINARY_NAME=backend


## build: builds the application
build:
	@echo "building..."
	env CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/${BINARY_NAME} ./cmd/api
	@echo "built!"


## run: builds then runs the application
run: build
	@echo "starting..."
	@env DSN=${DSN} ./bin/${BINARY_NAME}&
	@echo "started!"

## start: alias to run
start: run

## stop: stops the application
stop:
	@echo "stopping the application..."
	@-pkill -SIGTERM -f "./bin/${BINARY_NAME}"

## restart: restarts the application
restart: stop start

