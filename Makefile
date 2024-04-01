migrate-create:
	echo "Creating migration choerrypdb"
	migrate create -ext sql -dir pkg/databases/migrations -seq choerrypdb

migrate-up:
	echo "Migrating-up database"
	migrate -source file://pkg/databases/migrations -database 'postgres://choerrypagent:123456@localhost:4444/choerrypdb?sslmode=disable' -verbose up

migrate-down:
	echo "Migrating-down database"
	migrate -source file://pkg/databases/migrations -database 'postgres://choerrypagent:123456@localhost:4444/choerrypdb?sslmode=disable' -verbose down
