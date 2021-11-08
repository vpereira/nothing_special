all:
	docker run -ti -v "$(PWD):/web" golang:1.13 bash -c "cd /web && go clean && go build"
clean:
	go clean
test:
	go test -v ./...
