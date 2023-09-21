## Gateway API v1.0.0
##  
PORT=3000

.PHONY: build clean

help:          ## Show this help
	@sed -n '/^[a-z#]/p' Makefile | sed 's/:.*#/:\t/g' | sed 's/## //g'
	#@sed -n '/^[a-z#]/p' Makefile | sed 's/:.*#/:\t/g' | sed 's/## //g'

build: clean   ## Build the docker image
	docker build -t api-gateway .

run: build     ## Run the container (and build)
	docker run --rm --name api-gateway -p ${PORT}:3000 api-gateway:latest 

clean:         ## Remove the image
	#docker container rm -f api-gateway
	docker image rm -f api-gateway

