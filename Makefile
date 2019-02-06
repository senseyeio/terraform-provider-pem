default: build

build:
	go build 

test:
	go test ./... -v -parallel 20

testacc: 
	TF_ACC=1 go test ./... -v -parallel 20
