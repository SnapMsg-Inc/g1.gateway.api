## Gateway API v1.0.0
##  
PORT=3001

.PHONY: build clean fetch-id-token format

help:          ## Show this help
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

build: clean   ## Build the docker image
	docker build -t api-gateway .

run: build     ## Run the container (and build)
	docker run --rm --name api-gateway -p ${PORT}:3001 api-gateway:latest 

run-local:
	swag init
	go run main.go

format:        ## Format the source code
	gofmt -s -w . 
	#go fmt github.com/SnapMsg-Inc/g1.gateway.api

fetch-id-token:   ## Get a testing id token for the user EMAIL=<email> PASS=<password> (no quotes)
	curl 'https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyCmlfaciH4N_Ydih2RzNXEWr2G_V1En1sw' -H 'Content-Type: application/json' --data-binary '{"email":"${EMAIL}","password":"${PASS}","returnSecureToken":true}'

clean:         ## Remove the image
	#docker container rm -f api-gateway
	docker image rm -f api-gateway

