all: parcello
	bower install
	go get -u
	go generate ./...
	go install github.com/andrei-m/bank/bankcmd

parcello:
	go get -u github.com/phogolabs/parcello/cmd/parcello
	go install github.com/phogolabs/parcello/cmd/parcello

