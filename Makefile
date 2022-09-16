.PHONY: docker-build-image app db db-migrate run vendor sqlboiler-gen
init: docker-build-image 
setup: db db-migrate app 
run: 
	@docker-compose run -p 5000:5000 --rm app sh -c "go run cmd/serverd/main.go"

# TODO: make test command

docker-build-image:
	@docker build -f Dockerfile -t s3-golang-dev .

app:
	@docker-compose up -d app

db:
	@docker-compose up -d db

db-migrate:
	@docker-compose up -d db-migrate

sqlboiler-gen:
	@sqlboiler psql

vendor:
	@go mod tidy
	@go mod vendor
