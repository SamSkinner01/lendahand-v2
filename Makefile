include .env

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
.PHONY: air
air:
	~/go/bin/air -c .air.toml

.PHONY: dev
dev:
	sqlc generate
	~/go/bin/templ generate
	~/go/bin/air -c .air.toml

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql ${LENDAHAND_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${LENDAHAND_DB_DSN} up
