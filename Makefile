
IMAGE_NAME=shop_backend
CONTAINER_NAME=shop_backend_container
HOST_PORT=10000
CONTAINER_PORT=10000

.PHONY: backend build run clean

build:
	docker build -t $(IMAGE_NAME) .

run:
	-docker rm -f $(CONTAINER_NAME) || true
	docker run -d --name $(CONTAINER_NAME) -p $(HOST_PORT):$(CONTAINER_PORT) $(IMAGE_NAME)

backend: build run

clean:
	-docker rm -f $(CONTAINER_NAME) || true
	-docker rmi $(IMAGE_NAME) || true