migrate = migrate -path database/migrations 
migrate += -database "postgres://$(user):$(password)@localhost:5432/$(db)?sslmode=disable"

compose = docker-compose

migrate-up:
	$(migrate) -verbose up

migrate-down:
	$(migrate) -verbose down

build: 
	go build -o myriadcode cmd/web/main.go

run: build
	./myriadcode

docker:
	$(compose) up

docker-build:
	$(compose) build

docker-dev:
	$(compose) -f docker-compose.dev.yml up --build