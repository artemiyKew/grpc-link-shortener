.PHONY: build

build: 
	go build -v ./cmd/gateway-service

.PHONY: run

run: build
	./gateway-service

compose:
	docker compose up
.DEFAULT_GOAL = build