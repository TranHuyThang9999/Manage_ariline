BINARY_NAME = main.out

build:
	go build -o $(BINARY_NAME) main.go

run: build
	./$(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)
.PHONY: up down

up:
	docker-compose -f docker-compose.yaml up -d

down:
	docker-compose -f docker-compose.yaml down
