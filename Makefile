run r:
	@echo [Running Mini URL Builder API]
	@MINI_URLs_HOST=http://localhost:8081 MONGODB_URI="mongodb://root:password@localhost:27017/marketingDB?authSource=admin" SYNC_SERVICE_HOST=http://localhost:8079 JWT_TOKEN_SEED=supersecret APP_PORT=8080 APP_ENV=dev APP_VERSION=0.1 go run src/main.go

test t:
	@echo [Running Mini URL Builder API Tests]
	@go test ./src/...

run-compose rc:
	@echo [Running Mini URL Solution in Docker Compose]
	@docker compose up -d --build

migrate-dbs mg:
	@echo [Adding data to both MongoDB and Redis]
	@docker exec -it mongoMiniURL bash -c 'mongosh marketingDB -u root -p password --authenticationDatabase admin --eval "db.miniurls.insertOne({ original_url: '\''https://www.google.com'\'', new_url: '\''abc123'\'' })"'
	@docker exec -it redisMiniURL redis-cli set xyz789 "https://www.apple.com"

.PHONY: run r run-compose rc migrate-dbs mg