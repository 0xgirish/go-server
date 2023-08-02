migrate:
	@cd migrations && goose mysql "root:1234@/todos" up

migrate-down:
	@cd migrations && goose mysql "root:1234@/todos" down

build:
	@go build
