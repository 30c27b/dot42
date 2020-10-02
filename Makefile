BIN			= dot42

GOCMD		= go
GOBUILD		= $(GOCMD) build
GOCLEAN		= $(GOCMD) clean
GOINST		= $(GOCMD) install

all:		build

build:
			$(GOBUILD) -o $(BIN) -v

clean:
			$(GOCLEAN)
