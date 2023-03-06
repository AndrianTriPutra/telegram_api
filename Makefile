dev: 
	ENV=dev docker compose -f docker/docker-compose.yml down 
	ENV=dev docker compose -f docker/docker-compose.yml -p telegram_api up --force-recreate -d --build 
prod: 
	ENV=prod docker compose -f docker/docker-compose.yml down 
	ENV=prod docker compose -f docker/docker-compose.yml -p telegram_api up --force-recreate -d --build 