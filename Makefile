
.PHONY: down clean

run: 
	docker compose up -d

reload:
	docker compose build

clean: 
	docker compose down
	docker image rm -f api-gateway
	docker container rm -f api-gateway
