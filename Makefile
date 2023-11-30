.PHONY: build

build: 
	go build -v ./cmd/gateway-service

.PHONY: run

run: build
	./gateway-service


.DEFAULT_GOAL = build