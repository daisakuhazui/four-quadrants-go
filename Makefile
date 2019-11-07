serve:
	go run main.go

test:
	go test -cover

cover:
	go test -coverprofile cover.out
