all: clean binary

clean:
	go clean -v

binary:
	go build -v -mod=vendor

vendor:
	go mod vendor

