createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrateup:
	migrate -path=sql/migrations -database "mysql://admin:admin@admin@tcp(localhost:3306)/courses" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://admin:admin@admin@tcp(localhost:3306)/courses" -verbose down

. PHONY: createmigration migrateup migratedown