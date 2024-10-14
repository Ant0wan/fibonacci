GO   := go
NAME := fibonacci

all: build

build:
	CGO_ENABLED=0 GOOS=linux $(GO) build -v -o $(NAME)

run:
	CGO_ENABLED=0 GOOS=linux $(GO) run -v $(NAME)

clean:
	rm -rf $(NAME)

fclean: clean

.PHONY: all build  clean fclean
