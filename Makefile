COMPILER=go
BUILD=build
INSTALL=install
SRC=main.go page.go server.go rest.go
EXEC=direc-server

all:
	jekyll build -s ./html/ -d ./html/_site --incremental
	$(COMPILER) $(BUILD) -o $(EXEC) $(SRC)

# The install might not work as it is only part of the software.
install:
	$(COMPILER) $(INSTALL) -o $(EXEC) $(SRC)

jekyll:
	jekyll build -s ./html/ -d ./html/_site --incremental --watch


clear:
	rm $(EXEC)
	rm $GOPATH/bin/$(EXEC)
