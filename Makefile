open:
	cd frontend && ng serve --open --ssl

serve_frontend:
	cd frontend && ng serve --ssl

serve_backend:
	go run main.go

test:
	go test -v ./backend -cover

cover:
	go test -v ./backend -coverprofile cover.out &&\
	go tool cover -html=cover.out
