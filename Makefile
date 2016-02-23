all:
	go get github.com/advancedlogic/GoOse
	go get github.com/mattn/go-sqlite3
	go install crawler
	go install director

build:
	docker build -f Dockerfile.crawler -t munch-crawler .
	docker build -f Dockerfile.director -t munch-director .

run:
	docker-compose up
