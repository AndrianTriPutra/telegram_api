# telegram_api 

## endpoint
- /docker        

## getting started
- please switch to branch basic for started
- change token and chatID
- ENV=dev go run .
- curl --location 'http://localhost:8008/docker' --form 'message="[START] Service Producer Audio Ads"'

## deploy
- make dev  : for development
- make prod : for production
- logs      : docker logs -f telegram_api