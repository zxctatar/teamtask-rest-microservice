local:
	set -a; \
	. ./.env; \
	set +a; \
	go run cmd/userservice/main.go --config=./config/local.yaml

migrate_all_up:
	set -a; \
	. ./.env; \
	set +a; \
	migrate -path ./migrations -database "postgres://$$MIG_NAME:$$MIG_PASS@$$HOST:$$PORT/$$DBNAME?sslmode=$$SSLMODE" up