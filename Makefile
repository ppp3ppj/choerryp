.PHONY: tailwind-watch
tailwind-watch:
	bun tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	bun tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: build
build:
	make tailwind-build && make templ-generate && go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

.PHONY: vet
vet:
	go vet ./...

migrate-create:
	echo "Creating migration choerrypdb"
	migrate create -ext sql -dir internal/databases/migrations -seq choerrypdb

migrate-up:
	echo "Migrating-up database"
	migrate -source file://internal/databases/migrations -database 'postgres://choerrypagent:123456@localhost:4444/choerrypdb?sslmode=disable' -verbose up

migrate-down:
	echo "Migrating-down database"
	migrate -source file://internal/databases/migrations -database 'postgres://choerrypagent:123456@localhost:4444/choerrypdb?sslmode=disable' -verbose down
