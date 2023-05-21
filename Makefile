.PHONY: compose
compose:
	docker-compose up -d
.PHONY: compose-down
compose-down:
	docker-compose down --remove-orphans

.PHONY: build
build:
	docker-compose down --remove-orphans
	docker-compose build

.PHONY: test
test:
	go test -cover ./...