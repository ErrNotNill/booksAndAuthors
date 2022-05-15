build:
	docker-compose build books
#& docker volume create my-vol & docker volume ls & docker volume inspect my-vol

run:
	docker-compose up books

test:
	go test -v ./...

swag:
	swag init -g cmd/main.go

#migrate:
	#migrate -path ./schema -database 'jdbc:mysql://(host=localhost,port=3306,key1=value1)/mysql' up

