PORT=3000

.PHONY: build clean

build:
	docker build -t api-gateway .

run: clean build 
	docker run --name api-gateway -p ${PORT}:3000 -d api-gateway:latest 



clean: 
	docker container rm -f api-gateway
	docker image rm -f api-gateway
