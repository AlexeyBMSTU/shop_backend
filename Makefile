IMAGE_NAME=shop_backend
CONTAINER_NAME=shop_backend_container
HOST_PORT=10000
CONTAINER_PORT=10000

.PHONY: backend build run clean

build:
	-docker build -t $(IMAGE_NAME) .

run:
	-docker-compose up -d

run-dev:
	-docker-compose up -d --build

test:
	docker-compose run test

backend-dev: build run-dev

backend: build run

clean:
	-docker-compose down
	-docker rmi $(IMAGE_NAME) || true

logs:
	-docker-compose logs -f