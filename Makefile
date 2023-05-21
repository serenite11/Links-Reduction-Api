.PHONY: compose
compose:
	docker-compose up -d
.PHONY: compose-down
compose-down:
	migrate -path ./migrations -database postgres://user:qwerty@localhost:5433/links-reduction-db?sslmode=disable down
	docker-compose down --remove-orphans

.PHONY: build
build:
	docker-compose down --remove-orphans
	docker-compose build

.PHONY: test
test:
	go test -cover ./...