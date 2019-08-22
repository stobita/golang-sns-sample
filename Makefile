DEV_COMPOSE=deployments/development/docker-compose.yml

dev-up:
	docker-compose -f $(DEV_COMPOSE) up -d
dev-down:
	docker-compose -f $(DEV_COMPOSE) down
dev-logs:
	docker-compose -f $(DEV_COMPOSE) logs -f
dev-ps:
	docker-compose -f $(DEV_COMPOSE) ps
dev-bash:
	docker-compose -f $(DEV_COMPOSE) exec api ash

migrate-create:
	docker-compose -f $(DEV_COMPOSE) exec api go run cmd/migrate/migrate create $(NAME) sql
migrate:
	docker-compose -f $(DEV_COMPOSE) exec api go run cmd/migrate/migrate.go

