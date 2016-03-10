all: fetch compile_crawler compile_emailer compile_director compile_webapp

compile_crawler: src/crawler/crawler.go
	go install crawler

compile_director: src/director/director.go
	go install director

compile_emailer: src/emailer/emailer.go
	go install emailer

compile_webapp: src/webapp/webapp.go
	go install webapp

fetch:
	go get github.com/advancedlogic/GoOse
	go get github.com/mattn/go-sqlite3
	go get github.com/michaelklishin/rabbit-hole
	go get github.com/JesusIslam/tldr
	go get github.com/dchest/htmlmin 
	go get github.com/mailgun/mailgun-go
	go get github.com/go-yaml/yaml

build:
	docker build -f Dockerfile.crawler -t munch-crawler .
	docker build -f Dockerfile.director -t munch-director .

run:
	docker-compose up

db:
	sqlite3 broccoli.db < src/munch/database/schema.sql
