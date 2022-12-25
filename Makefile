EXEC = main

BIN_PATH = ./bin
BIN_FILES := $(shell ls $(BIN_PATH) | wc -l)
EMPTY_BIN := "Bin folder is empty. Build first."
CHECK_BIN := if [ $(BIN_FILES) = 0 ]; then echo $(EMPTY_BIN)

GOOS := $(shell uname | tr [:upper:] [:lower:])
GOARCH := $(shell dpkg --print-architecture | tr [:upper:] [:lower:])

# Air binary will be $(go env GOPATH)/bin/air
install-air:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

dev:
	air

cover:
	@/usr/bin/python -c "import webbrowser; webbrowser.open('cover.html');" &

build:
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o $(BIN_PATH)/$(EXEC) $(EXEC).go

clean:
	@$(CHECK_BIN); \
	else go clean && rm $(BIN_PATH)/*; \
	fi

run:
	@$(CHECK_BIN); \
	else ${BIN_PATH}/$(EXEC); \
	fi

containers-build:
	docker-compose build

containers-run:
	docker-compose up --remove-orphans