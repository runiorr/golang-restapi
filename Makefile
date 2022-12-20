EXEC = main

BIN_PATH = ./bin
BIN_FILES := $(shell ls $(BIN_PATH) | wc -l)

GOOS := $(shell uname | tr [:upper:] [:lower:])
GOARCH := $(shell dpkg --print-architecture | tr [:upper:] [:lower:])

# Binary will be $(go env GOPATH)/bin/air
install-air:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

dev:
	air

cover:
	@/usr/bin/python -c "import webbrowser; webbrowser.open('cover.html');" &

build:
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o $(BIN_PATH)/$(EXEC) $(EXEC).go


clean:
	@if [ $(BIN_FILES) = 0 ]; then echo "Bin folder is empty"; \
	else go clean && rm $(BIN_PATH)/* ;\
	fi

run:
	${BIN_PATH}/$(EXEC)