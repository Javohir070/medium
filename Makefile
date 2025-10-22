-include .env

.SILENT:

DB_URL := "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DB)?parseTime=true"
database := $(DB_URL)
tidy:
     	@go mod tidy
		@go mod vendor

run:
		@go run cmd/main.go

migration:
        @migrate create -ext sql -dir ./migration -seq $(name)

migrateup:
		@migrate -path ./migration -database $(database) -verbose up
migratedown:
		@migrate -path ./migration -database $(database) -verbose down