run:
	-go run cmd/app/main.go
dev:
	nodemon --signal SIGHUP --exec "make run" -e "go"