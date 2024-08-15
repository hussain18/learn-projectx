build:
	go build -o ./bin/learn-projectx

run: build	
	./bin/learn-projectx

test:
	go test -v ./...