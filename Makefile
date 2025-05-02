DOCKER_USERNAME ?= misterdelle
APPLICATION_NAME ?= miner-and-commander
GIT_HASH ?= $(shell git log --format="%h" -n 1)

build:
	env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o build/${APPLICATION_NAME}.exe

build-linux:
	env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/${APPLICATION_NAME}-linux

build-arm:
	env GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "-s -w" -o build/${APPLICATION_NAME}-arm

build-docker-amd:
	docker build --tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:${GIT_HASH} .

build-docker-arm:
	docker buildx build --platform linux/arm64/v8 --output type=docker -t ${DOCKER_USERNAME}/${APPLICATION_NAME}:${GIT_HASH} .
	docker tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:${GIT_HASH} ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest

push:
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:${GIT_HASH}
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest

build-docker-arm-push:
	make build-docker-arm push
