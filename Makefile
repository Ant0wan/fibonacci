GO   := go
NAME := fibonacci

all: build

build:
	CGO_ENABLED=1 GOOS=linux CGO_LDFLAGS="-L/usr/local/lib -lgmp" CGO_CFLAGS="-I/usr/local/include" $(GO) build -v -o $(NAME)

run:
	CGO_ENABLED=1 GOOS=linux CGO_LDFLAGS="-L/usr/local/lib -lgmp" CGO_CFLAGS="-I/usr/local/include" $(GO) run -v $(NAME)

clean:
	rm -rf $(NAME)

fclean: clean

.PHONY: all build  clean fclean
