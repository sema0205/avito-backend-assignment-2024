include .env
export

compose-up:
	docker-compose up --build -d && docker-compose logs -f
.PHONY: compose-up

compose-down:
	docker-compose down --remove-orphans
.PHONY: compose-down

docker-rm-volume:
	docker volume rm pg-data
.PHONY: docker-rm-volume

migrate-up:
	migrate -path db/migrations -database '$(POSTGRES_URL)?sslmode=disable' up
.PHONY: migrate-up

migrate-down:
	echo "y" | migrate -path db/migrations -database '$(POSTGRES_URL)?sslmode=disable' down
.PHONY: migrate-down

test:
	go test -v ./...

mockgen:
	mockgen -source=internal/service/service.go -destination=internal/mocks/servicemocks/service.go -package=servicemocks
	mockgen -source=internal/repository/repository.go -destination=internal/mocks/repomocks/repository.go -package=repositorymocks
	mockgen -source=pkg/auth/provider.go -destination=internal/mocks/authmocks/provider.go -package=authmocks
	mockgen -source=pkg/cache/provider.go -destination=internal/mocks/cachemocks/provider.go -package=cachemocks
.PHONY: mockgen

swag:
	swag init -g internal/app/app.go --parseInternal --parseDependency
