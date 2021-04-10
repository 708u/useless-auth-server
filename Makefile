.PHONY: up-auth
up-auth:
	go run cmd/auth/main.go

.PHONY: up-client
up-client:
	go run cmd/client/main.go

.PHONY: update-debug-conf
update-degug-conf:
	\cp -f configs/auth/config.yml.example cmd/auth/configs/auth/config.yml
	\cp -f configs/client/config.yml.example cmd/client/configs/client/config.yml
