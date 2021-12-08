PATCH_DIR	= patch

DB_IMG_NAME = go-pg

MIGRATION_PATH	= db/migration


patch:
	cd patch && sh patch.sh patch.json

init-db:
	docker run --name $(DB_IMG_NAME) -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres

restart-db:
	docker start $(DB_IMG_NAME)

create-db:
	docker exec -it $(DB_IMG_NAME) createdb --username=root --owner=root ohlc

drop-db:
	docker exec -it $(DB_IMG_NAME) dropdb ohlc

clean-db:
	docker kill $(DB_IMG_NAME)
	docker rm $(DB_IMG_NAME)

migrateup:
	migrate -path $(MIGRATION_PATH) -database "postgresql://root:secret@localhost:5432/ohlc?sslmode=disable" -verbose up

migratedown:
	migrate -path $(MIGRATION_PATH) -database "postgresql://root:secret@localhost:5432/ohlc?sslmode=disable" -verbose down

sqlc:
	sqlc generate


.PHONY:	patch init-db restart-db create-db drop-db clean-db migrateup migratedown sqlc
