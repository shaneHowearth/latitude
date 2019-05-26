.PHONY: all
all:
	go build -o latitude cmd/main.go
	docker build -t "latitude:shane" .
	docker run -t -d --name=shane latitude:shane
