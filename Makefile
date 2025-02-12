GO   := go
NAME := fibonacci

all: build

build:
	CGO_ENABLED=0 GOOS=linux $(GO) build -v -o $(NAME)

run:
	CGO_ENABLED=0 GOOS=linux $(GO) run -v $(NAME)

test:
	$(GO) test -v

clean:
	rm -rf $(NAME)

fclean: clean

.PHONY: all build test clean fclean
