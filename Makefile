all: clean binary

clean:
	go clean -v

binary:
	go build -v


