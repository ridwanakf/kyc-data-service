# go build command
build:
	@echo " >> building binaries"
	@go build -v -o kyc-data-service cmd/*.go

# go run command
run: build
	./kyc-data-service

