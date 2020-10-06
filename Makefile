BIN			= dot42
BINDIR		= bin

GOCMD		= go
GOBUILD		= $(GOCMD) build
GOCLEAN		= $(GOCMD) clean
GOINST		= $(GOCMD) install

all:		build

build:
			$(GOBUILD) -o $(BINDIR)/$(BIN) -v

clean:
			$(GOCLEAN)
