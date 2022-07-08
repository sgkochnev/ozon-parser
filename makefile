run:
	go run ./cmd/main.go

docker-up:
	docker-compose up

docker-down:
	docker-compose down

del-volumes:
	sudo rm -R volumes
