run r:
	@echo [Running Mini URL Builder API]
	@MINI_URLs_HOST=http://localhost:8081 MONGODB_URI="mongodb://root:password@localhost:27017/marketingDB?authSource=admin" SYNC_SERVICE_HOST=http://localhost:8079 JWT_TOKEN_SEED=supersecret APP_PORT=8080 APP_ENV=dev APP_VERSION=0.1 go run src/main.go

.PHONY: run r