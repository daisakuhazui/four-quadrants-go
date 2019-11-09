serve:
	go run main.go

test:
	go test -v ./backend -cover

cover:
	go test -v ./backend -coverprofile cover.out &&\
	go tool cover -html=cover.out
