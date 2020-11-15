.PHONY: server
server:
	TZ=US/Eastern go run server/main.go -port 8010 &
	TZ=Asia/Tokyo go run server/main.go -port 8020 &
	TZ=Europe/London go run server/main.go -port 8030 &

.PHONY: client
client:
	go run client/main.go
