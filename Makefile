include .env

postgres-run:
	docker run --name my-postgres -e POSTGRES_USER=$(PG_USER) -e POSTGRES_PASSWORD=$(PG_PASSWORD) -e POSTGRES_DB=$(PG_NAME) -p $(PG_PORT):$(PG_PORT) -d postgres:14-alpine3.17