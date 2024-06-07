GO_BINARY_NAME=reverse_proxy
watch:
	nodemon --exec go run main.go server --signal SIGTERM
start:
	go run main.go server
build:
	go build -a -v -o $(GO_BINARY_NAME) main.go