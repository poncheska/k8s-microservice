.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 go build -o ./bin/micro-services/fast/fast-bin ./cmd/micro-services/fast/main.go
	GOOS=linux CGO_ENABLED=0 go build -o ./bin/micro-services/slow/slow-bin ./cmd/micro-services/slow/main.go
	GOOS=linux CGO_ENABLED=0 go build -o ./bin/monolith/monolith-bin ./cmd/monolith/main.go
	docker build -t fast-ms ./bin/micro-services/fast
	docker build -t slow-ms ./bin/micro-services/slow
	docker build -t monolith ./bin/monolith/

.PHONY: run-ms
run-ms:
	docker run -p 8080:8080 slow-ms:latest

.PHONY: run-mnt
run-mnt:
	docker run -p 8080:8080 monolith:latest