
.PHONY: down clean

run: clean 
	docker compose up

clean: 
	docker compose down
	docker image rm -f api-gateway
	docker container rm -f api-gateway
