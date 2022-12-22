mongod:
	docker-compose -f ./docker-compose.yml --env-file .env up -d

prod:
	GOOS=linux GOARCH=amd64 go build main.go && scp -r main templates yogasearch:/home/ubuntu/crawler_job