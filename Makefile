.PHONY: up-auth
up-auth:
	go run cmd/auth/main.go

.PHONY: up-client
up-client:
	go run cmd/client/main.go
