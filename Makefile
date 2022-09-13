.PHONY: docker-build-image app db db-migrate run vendor
init: vendor docker-build-image  
setup: db db-migrate app 
run: docker-run-image

docker-build-image:
	@docker build -f Dockerfile -t s3-golang-assignment-app:latest .
app:
	@docker-compose up -d app
db:
	@docker-compose up -d db
db-migrate:
	@docker-compose up -d db-migrate
docker-run-image:
	@docker-compose run -p 5000:5000 --rm app sh -c "go run cmd/serverd/main.go"
vendor:
	@go mod tidy
	@go mod vendor