serve:
	go run main.go

test:
	go test ./backend -cover

cover:
	go test -coverprofile cover.out
