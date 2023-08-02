ifeq ($(OS),Windows_NT)
  EXECUTABLE_EXTENSION := .exe
else
  EXECUTABLE_EXTENSION :=
endif

GO_FILES = $(shell find . -type f -name '*.go')

all: balloon

balloon: $(GO_FILES)
	cd cmd/balloon && go build && cd ../..
	rm -f balloon
	ln -s cmd/balloon/balloon$(EXECUTABLE_EXTENSION) balloon


clean:
	cd cmd/balloon && go clean
	rm -f balloon
