## Gateway API v1.0.0
##  
PORT=3000

.PHONY: build clean fetch-token

help:          ## Show this help
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

build: clean   ## Build the docker image
	docker build -t api-gateway .

run: build     ## Run the container (and build)
	docker run --rm --name api-gateway -p ${PORT}:3000 api-gateway:latest 

fetch-token:
	curl 'https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyCmlfaciH4N_Ydih2RzNXEWr2G_V1En1sw' -H 'Content-Type: application/json' --data-binary '{"email":"example@example.com","password":"123456","returnSecureToken":true}'

clean:         ## Remove the image
	#docker container rm -f api-gateway
	docker image rm -f api-gateway

