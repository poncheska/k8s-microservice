.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 go build -o ./bin/micro-services/fast/fast-bin ./cmd/micro-services/fast/main.go
	GOOS=linux CGO_ENABLED=0 go build -o ./bin/micro-services/slow/slow-bin ./cmd/micro-services/slow/main.go
	GOOS=linux CGO_ENABLED=0 go build -o ./bin/monolith/monolith-bin ./cmd/monolith/main.go
	docker build -t poncheska/fast-ms ./bin/micro-services/fast
	docker build -t poncheska/slow-ms ./bin/micro-services/slow
	docker build -t poncheska/monolith ./bin/monolith/

.PHONY: run-ms
run-ms:
	docker run -p 8080:8080 poncheska/slow-ms

.PHONY: run-mnt
run-mnt:
	docker-compose -f ./bin/monolith/docker-compose.yml up