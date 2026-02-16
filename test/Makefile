APP = test

.PHONY: test
test: 
	go test ./... -coverprofile=coverage.out --cover

.PHONY: build
build: 
	env GOOS=linux GOARCH=amd64 go build -o test main.go

.PHONY: install
install:
	go mod tidy

.PHONY: run
run:
	go run main.go

.PHONY: build-docker
build-docker:
	docker build -t $(APP):latest -f docker/Dockerfile .

.PHONY: serve
serve:
	./$(APP)

.PHONY: config
config:
	cp .env.example .env
