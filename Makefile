all:
	go get github.com/advancedlogic/GoOse
	go get github.com/mattn/go-sqlite3
	go install crawler

	docker build -f Dockerfile.crawler -t munch-crawler .
