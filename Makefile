all: libs run

run: 
	go run main.go

build: fclean libs
	go build main.go

libs:
	go get github.com/shomali11/slacker

clean:
	rm -Rf ./main
fclean: clean
re: fclean all