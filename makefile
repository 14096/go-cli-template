GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
MAIN_PACKAGE=main.go


clean:
	@$(GOCLEAN)
	@rm -rf ./bin


tidy:
	clear
	@go mod tidy


build:
	clear
	@$(GOBUILD) -o bin/myApp $(MAIN_PACKAGE)


run: build
	clear
	@./bin/myApp


test: build
	clear
	@$(GOTEST) ./tests


myCommand: build
	clear
	@./bin/myApp my-command -u "john doe" -m "jdoe@mail.com"

