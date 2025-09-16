DIR=${CURDIR}

up:
	docker compose up -d minio
	$(MAKE) service-up SERVICE=file-storage

service-up:
	docker compose up -d ${SERVICE}-postgres
	timeout 3
	$(MAKE) migrations-up SERVICE=${SERVICE}
	docker compose up -d ${SERVICE}

service-rebuild:
	docker compose build ${SERVICE}

.PHONY:migrations-up
migrations-up:
	docker run --rm -v $(DIR)/pkg/${SERVICE}/repository/postgres/migrations:/migrations \
		--network shaream_default migrate/migrate -path=/migrations/ \
		-database postgres://dev:123456@${SERVICE}-postgres:5432/${SERVICE}?sslmode=disable up